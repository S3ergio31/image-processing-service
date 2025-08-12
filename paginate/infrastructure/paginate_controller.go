package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/S3ergio31/image-processing-service/paginate/application"
	"github.com/S3ergio31/image-processing-service/paginate/domain"
	"github.com/gin-gonic/gin"
)

func PaginateController(c *gin.Context) {
	paginator := application.Paginator{
		ImageRepository: NewSqlitePaginateImageRepository(),
	}

	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))

	imagePagination := domain.NewImagePagination(
		c.GetString("username"),
		page,
		limit,
	)

	images := paginator.Paginate(imagePagination)

	c.JSON(http.StatusOK, ToImageResponse(images))
}
