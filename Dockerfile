FROM golang:1.20

WORKDIR /app

RUN go install github.com/cosmtrek/air@latest

COPY go.mod go.sum ./
RUN go mod download
 RUN export CGO_CFLAGS="-g -O2 -Wno-return-local-addr"
CMD ["air", "-c", ".air.toml"]