FROM golang:1.17-alpine

WORKDIR /app

COPY . ./

RUN go mod download

RUN go build -o /Go-app

EXPOSE 3000

CMD [ "/Go-app" ]