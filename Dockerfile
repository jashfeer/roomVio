# Build stage
FROM golang:1.19-alpine3.16 AS builder
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go build -o main main.go

# Run stage 
FROM alpine:3.16
COPY --from=builder /app .

CMD [ "./main" ]