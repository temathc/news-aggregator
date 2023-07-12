FROM golang:latest

WORKDIR /app/news-aggreagtor

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o /news-aggreagtor-service ./cmd/service/main.go
RUN go build -o /news-aggreagtor-ticker ./cmd/ticker/main.go
RUN go install github.com/pressly/goose/v3/cmd/goose@latest

CMD ["/news-aggreagtor-service"]