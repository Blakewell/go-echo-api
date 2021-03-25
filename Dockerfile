FROM golang:1.12-alpine AS GO_BUILD 

RUN apk add --no-cache git

COPY api /code
WORKDIR /code
RUN go build -o /go/bin/server/api

FROM alpine:latest
WORKDIR /app
COPY --from=GO_BUILD /go/bin/server/ ./
EXPOSE 1323 
CMD ./api