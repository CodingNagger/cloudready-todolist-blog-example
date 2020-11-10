FROM golang:1.15.3-alpine3.12

EXPOSE 8080/tcp

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./cmd/server
RUN go install -v ./cmd/server

CMD [ "sh", "-c", "server", "-endpoint", "$AWS_ENDPOINT" ]