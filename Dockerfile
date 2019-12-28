FROM golang:1.13.5-alpine3.10 as builder

WORKDIR /go/src/app

COPY . .

RUN apk update && apk add --no-cache make

RUN make build

FROM alpine:3.10 as runner

WORKDIR /usr/bin

RUN apk update && apk add --no-cache ca-certificates

COPY --from=builder /go/src/app/bin/api .

CMD [ "./api" ]

EXPOSE 3000
