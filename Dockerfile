FROM golang:1.21

WORKDIR /go/github.com/trstruth/sheep

COPY . .

COPY ./conf/sheep.toml /etc/sheep/sheep.toml

RUN go build ./...

CMD ["./sheep"]