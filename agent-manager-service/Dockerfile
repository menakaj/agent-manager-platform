FROM golang:1.25-alpine AS build-env

RUN mkdir /app
WORKDIR /app
COPY go.mod . 
COPY go.sum .

# Get dependancies - will also be cached if we won't change mod/sum
RUN go mod download

# COPY the source code as the last step
COPY . .


# Build the binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /go/bin/app -buildvcs=false

FROM alpine:3

# Create a non-root user/group (UID:GID = 1000:1000)
RUN addgroup -g 1000 app && adduser -D -u 1000 -G app app

# Create directory for the OpenAPI schema
RUN mkdir -p /app/clients/openchoreosvc && chown -R 1000:1000 /app

# Copy binary with ownership set
COPY --from=build-env --chown=1000:1000 /go/bin/app /go/bin/agent-manager-service

# Copy the OpenAPI schema file
COPY --from=build-env --chown=1000:1000 /app/clients/openchoreosvc/default-openapi-schema.yaml /app/clients/openchoreosvc/default-openapi-schema.yaml

# Drop root privileges
USER 1000:1000

# Set working directory to /app so relative paths work
WORKDIR /app

ENV GODEBUG=x509negativeserial=1
ENTRYPOINT ["/go/bin/agent-manager-service"]
