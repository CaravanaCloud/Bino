FROM golang:1.16-buster
WORKDIR /app
COPY main.go go.mod go.sum /app
RUN go mod download && go build -o bino
CMD ["./bino"]