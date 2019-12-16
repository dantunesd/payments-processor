FROM golang:1.13.5-alpine3.10 as builder

WORKDIR /go/src/app

COPY . .

RUN go build -o api

FROM alpine:3.10 as runner

WORKDIR /usr/bin

RUN apk update && apk add --no-cache ca-certificates

COPY --from=builder /go/src/app/api .

CMD [ "./api" ]

EXPOSE 3000
