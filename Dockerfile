FROM golang:1.13.5-alpine3.10

WORKDIR /payments-processor

COPY . .

RUN go build

EXPOSE 3000

CMD [ "./payments-processor" ]
