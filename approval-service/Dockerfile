FROM golang:1.19.3-alpine AS builder

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN go build -o approval-service /app/main.go


FROM alpine

RUN mkdir /app

COPY --from=builder /app/approval-service /app

CMD [ "/app/approval-service" ]