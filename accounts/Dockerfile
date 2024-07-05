FROM golang:1.22.4

WORKDIR /accounts

RUN ls -la

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY ../accounts .

CMD ["go", "run", "/accounts/cmd/main.go"]