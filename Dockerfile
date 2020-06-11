FROM golang:1.13

RUN mkdir /app

ADD . /app

WORKDIR /app

RUN go get -d -v ./...

RUN go install -v ./...

EXPOSE 9090

CMD ["/app/main"]