#build stage
FROM golang:alpine AS builder
WORKDIR /iitk-coin
RUN apk add --no-cache gcc musl-dev linux-headers
COPY . .
RUN go build 

#final stage
FROM alpine:latest
WORKDIR /iitk-coin
COPY --from=builder /iitk-coin/iitk-coin ./
LABEL Name=iitk-coin Version=0.0.1

EXPOSE 8080
ENTRYPOINT ["./iitk-coin"]
