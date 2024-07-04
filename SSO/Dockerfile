FROM golang:1.22.4

WORKDIR /SSO

RUN ls -la

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY ../SSO .

CMD ["go", "run", "/SSO/cmd/main.go"]