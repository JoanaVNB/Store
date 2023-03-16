FROM golang:latest

WORKDIR /app
ADD . .
COPY . /app
EXPOSE 5000

RUN go build main.go
ENTRYPOINT ["./main"]

