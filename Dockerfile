FROM golang:1.17-alpine as builder

WORKDIR /app

COPY . ./

RUN go mod download

RUN go build -o /go-app

FROM alpine:3.13.6

COPY --from=builder /go-app /go-app

EXPOSE 8080

CMD ["/go-app"]