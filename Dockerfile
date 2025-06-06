FROM golang:1.24.2

RUN apt-get update && apt-get install -y \
    build-essential \
    libvips-dev \
    pkg-config \
    git \
 && rm -rf /var/lib/apt/lists/*

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

CMD ["go", "run", "main.go"]
