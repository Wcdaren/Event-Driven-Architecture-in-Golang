FROM golang:1.21-alpine AS builder

WORKDIR /server

# Retrieve application dependencies.
# This allows the container build to reuse cached dependencies.
# Expecting to copy go.mod and if present go.sum.
COPY go.* ./
RUN go mod download

# Copy local code to the container image.
COPY . ./

# Build the binary.
RUN go build -v -o server ./cmd/server

FROM alpine:latest AS runtime

COPY --from=builder /server/docker/wait-for .
RUN chmod +x /wait-for

# Copy the binary to the production image from the builder stage.
COPY --from=builder /server/server /server/server

CMD ["/server/server"]
