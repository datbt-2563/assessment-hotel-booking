# Builder
FROM golang:1.21.5-alpine3.19 as builder

RUN apk update && apk upgrade && \
    apk --update add git make

WORKDIR /app

COPY . .

RUN make hotel-booking

# Distribution
FROM alpine:latest

RUN apk update && apk upgrade && \
    apk --update --no-cache add tzdata && \
    mkdir /app

WORKDIR /app

COPY --from=builder /app/hotel-booking /app

CMD /app/hotel-booking
