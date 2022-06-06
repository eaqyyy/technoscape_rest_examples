FROM golang:1.14.9-alpine
RUN mkdir app
ADD . /app
WORKDIR /app
RUN go mod tidy
RUN go build -o main .
CMD ["/app/main"]