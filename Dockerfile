FROM golang:1.17-alpine as builder

WORKDIR /app

COPY . ./

RUN go mod download

RUN go build -o /Go-app


FROM alpine:3.13.6

RUN apk add --no-cache ca-certificates

COPY --from=builder /Go-app /Go-app

EXPOSE 3000

CMD ["/Go-app"]