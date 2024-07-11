FROM golang:1.22.4

WORKDIR /server

RUN ls -la

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY ../grpc/server .
COPY ../cmd/grpc/server .

CMD ["go", "run", "main.go"]