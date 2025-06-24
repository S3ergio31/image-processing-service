FROM golang:1.24.2

RUN apt-get update && apt-get install -y \
    build-essential \
    libvips-dev \
    libsqlite3-dev \
    pkg-config \
    git \
 && rm -rf /var/lib/apt/lists/*

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN echo "alias tests='go test ./tests/... -count=1'" >> ~/.bashrc

CMD ["go", "run", "main.go"]