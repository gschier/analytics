FROM golang:alpine

ADD . .

ENV GOPATH=""
RUN go mod vendor && go install

CMD ["analytics"]