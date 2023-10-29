FROM golang:1.21

WORKDIR /app
COPY . .
RUN go mod tidy

RUN go build -o main /app/cmd
CMD ["./main"]
