FROM golang:1.14

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["app"]

FROM golang:alpine as builder
RUN apk --no-cache add build-base git mercurial gcc make
ADD . /src
RUN cd /src && make build

FROM alpine
WORKDIR /app
COPY --from=builder /src/web-stash-api /app/
ENTRYPOINT ./web-stash-api