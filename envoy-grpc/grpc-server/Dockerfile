FROM golang:1.15-alpine as builder

RUN apk add --update make

COPY cmd /build/cmd/
COPY gen /build/gen/
COPY internal /build/internal/
COPY go.mod go.sum Makefile /build/

RUN ls -la /build/

WORKDIR /build

#ENV GO111MODULE=on
RUN make deps

# Build for linux
ENV GOOS=linux
ENV GOARCH=amd64
RUN go build -o server ./cmd/hello

# Build the main container (Linux Runtime)
FROM alpine:latest
WORKDIR /root/

# Copy the linux binary
COPY --from=builder /build/server /bin/

ENTRYPOINT /bin/server
