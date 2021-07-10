FROM golang

ADD . .

CMD ["go", "run", "main.go"]