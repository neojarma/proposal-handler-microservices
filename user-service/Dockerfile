FROM golang:1.19.3-alpine AS builder

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN go build -o user-service /app/main.go


FROM alpine

RUN mkdir /app

COPY --from=builder /app/user-service /app

CMD [ "/app/user-service" ]