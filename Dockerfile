FROM golang:1.16-buster
WORKDIR /app
COPY . /app
RUN go mod download && go build -o bino
CMD ["./bino"]