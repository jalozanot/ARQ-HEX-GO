# build stage
FROM golang:latest as builder

# Add Maintainer Info
LABEL maintainer="Franklin carrero <mauriciocarrero15@gmail.com>"

ENV GO111MODULE=on

WORKDIR /go/src/github.com/jalozanot/demoCeiba/


# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

COPY . .


# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o demo_api .

# final stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /go/src/github.com/jalozanot/demoCeiba/demo_api .

# Expose port 8081 to the world:
EXPOSE 8084

CMD ["./demo_api"]
