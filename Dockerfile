#1st build
FROM golang:1.10
WORKDIR /go/src/github.com/emailtovamos/m-bff

COPY vendor ./vendor
COPY cli ./cli
COPY bff ./bff

RUN GOOS=linux go install ./cli

#2nd Stage

FROM alpine:latest
WORKDIR /app/

COPY --from=0 /go/bin/cli ./binary

CMD ./binary