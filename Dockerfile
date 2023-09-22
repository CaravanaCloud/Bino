FROM golang:1.20-buster
WORKDIR /app
COPY . /app
RUN go mod download && go build -o bino
CMD ["./bino"]