FROM golang:1.19.0

RUN mkdir /app
WORKDIR /app

COPY . .
RUN go mod download
RUN go mod verify
RUN go build -o main .

CMD ["./main"]