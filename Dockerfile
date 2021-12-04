FROM golang:1.17-alpine

WORKDIR /backend_mini

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o mainfile

EXPOSE 9000

CMD ["./mainfile"]
