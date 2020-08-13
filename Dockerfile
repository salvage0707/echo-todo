FROM golang:1.15.0 as builder

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

WORKDIR /go/src/github.com/app
ADD . .

RUN go build main.go

FROM alpine
COPY --from=builder /go/src/github.com/app /app

RUN ls /app

CMD /app/main $PORT