FROM golang

ADD . .

RUN go mod vendor

CMD ["go", "run", "main.go"]