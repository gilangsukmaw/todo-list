FROM golang:alpine

RUN docker system prune

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go mod tidy

RUN go build -o binary

ENTRYPOINT ["/app/binary"]

EXPOSE 8080

CMD ["http"]
