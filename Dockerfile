FROM golang:1.7.5

EXPOSE 8081

WORKDIR /go/src/app
COPY . .

RUN go get -d -v
RUN go install -v

CMD ["app"]