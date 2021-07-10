FROM golang

ADD . .

ENV GOPATH=""
RUN go mod vendor

CMD ["go", "run", "main.go"]