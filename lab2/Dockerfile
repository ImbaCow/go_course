#build stage
FROM golang:alpine AS builder
RUN apk add --no-cache git gcc bash musl-dev make
# WORKDIR /go/bin/app
# COPY . .
RUN go env -w GO111MODULE=auto
# RUN go get github.com/golangci/golangci-lint/cmd/golangci-lint
# RUN go mod tidy

# #final stage
# FROM alpine:latest
# RUN apk --no-cache add ca-certificates
# COPY --from=builder /go/bin/app /app
# ENTRYPOINT ./app
# LABEL Name=lab1 Version=0.0.1
# EXPOSE 8000
