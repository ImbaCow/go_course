#build stage
FROM golang:alpine AS builder
RUN apk add --no-cache git
# WORKDIR /go/src/app
RUN go env -w GO111MODULE=auto
# RUN go get -d -v ./...
# RUN go install -v ./...

# #final stage
# FROM alpine:latest
# RUN apk --no-cache add ca-certificates
# COPY --from=builder /go/bin/app /app
# ENTRYPOINT ./app
# LABEL Name=lab1 Version=0.0.1
# EXPOSE 8000
