package infrastructure

import (
	"log"
	"net/http"
	"sort"
	"time"

	"github.com/gin-gonic/gin"
)

type throttle struct {
	perMinute int
	tries     []time.Time
}

func (t *throttle) diff() float64 {
	return t.last().Sub(t.first()).Abs().Seconds()
}

func (t *throttle) first() time.Time {
	return t.tries[0]
}

func (t *throttle) last() time.Time {
	return t.tries[len(t.tries)-1]
}

func (t *throttle) isTooManyRequest() bool {
	return t.diff() <= 60 && len(t.tries) > t.perMinute
}

func (t *throttle) shouldBeReset() bool {
	return t.diff() > 60
}

func RateLimiter(perMinute int) gin.HandlerFunc {
	requetsPerIp := make(map[string]map[string][]time.Time)
	return func(c *gin.Context) {
		throttle := throttle{
			tries:     getTries(c, &requetsPerIp),
			perMinute: perMinute,
		}

		logTries(c, &throttle)

		if throttle.shouldBeReset() {
			log.Printf("RateLimiter: limit completed")
			requetsPerIp[c.ClientIP()][c.FullPath()] = []time.Time{throttle.last()}
		}

		if throttle.isTooManyRequest() {
			c.AbortWithStatusJSON(
				http.StatusTooManyRequests,
				gin.H{"error": "too many request"},
			)
			return
		}

		c.Next()
	}
}

func getTries(c *gin.Context, requetsPerIp *map[string]map[string][]time.Time) []time.Time {
	_, hasIp := (*requetsPerIp)[c.ClientIP()]

	if !hasIp {
		(*requetsPerIp)[c.ClientIP()] = make(map[string][]time.Time)
	}

	_, ok := (*requetsPerIp)[c.ClientIP()][c.FullPath()]

	if !ok {
		(*requetsPerIp)[c.ClientIP()][c.FullPath()] = []time.Time{time.Now()}
	} else {
		(*requetsPerIp)[c.ClientIP()][c.FullPath()] = append((*requetsPerIp)[c.ClientIP()][c.FullPath()], time.Now())
	}

	tries := (*requetsPerIp)[c.ClientIP()][c.FullPath()]

	sort.Slice(tries, func(i, j int) bool {
		return tries[i].Before(tries[j])
	})

	return tries
}

func logTries(c *gin.Context, throttle *throttle) {
	log.Printf(
		"RateLimiter: ip=%s path=%s diff=%f perMinute=%d tries=%d",
		c.ClientIP(),
		c.FullPath(),
		throttle.diff(),
		throttle.perMinute,
		len(throttle.tries),
	)
}
