FROM golang:1.22.4

WORKDIR /accounts

RUN ls -la

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY ../http/accounts .
COPY ../cmd/http/server .

CMD ["go", "run", "main.go"]