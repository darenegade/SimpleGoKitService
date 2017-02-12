FROM golang:1.7.5

EXPOSE 8081

WORKDIR /go/src/github.com/darenegade/SimpleGoKitService
COPY . .

RUN go get -d -v
RUN go install -v

CMD ["SimpleGoKitService"]