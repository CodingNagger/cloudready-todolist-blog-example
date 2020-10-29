FROM golang:1.15

ENV AWS_REGION=
ENV AWS_ENDPOINT=

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./cmd/server
RUN go install -v ./cmd/server

CMD server -region $AWS_REGION -endpoint $AWS_ENDPOINT