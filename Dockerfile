FROM golang

ADD . .

ENV GOPATH=""
RUN go mod vendor && go install

CMD ["analytics"]