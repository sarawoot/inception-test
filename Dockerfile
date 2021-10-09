FROM golang:1.16-alpine AS build

ENV GOOS=linux
WORKDIR /app
COPY . .
RUN apk add build-base
RUN go mod download
RUN go mod verify

RUN GOOS=linux go build -ldflags="-w -s"  -o api cmd/api/main.go

FROM alpine:3.14

WORKDIR /app
RUN apk add build-base
COPY --from=build /app/api .
COPY --from=build /app/payment.db .
COPY --from=build /usr/local/go/lib/time/zoneinfo.zip /
ENV TZ=Asia/Bangkok
ENV ZONEINFO=/zoneinfo.zip

ENTRYPOINT [ "./api" ]