FROM golang:1.8.3

ADD . /go/src/github.com/ykhrustalev/highloadcup

WORKDIR /go/src/github.com/ykhrustalev/highloadcup

RUN make build && make install

EXPOSE 80

ENV PORT=80

CMD /go/bin/app
