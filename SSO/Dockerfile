FROM golang:1.22.0

WORKDIR /app

RUN ls -la

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

CMD ["go", "run", "/app/cmd/main.go"]