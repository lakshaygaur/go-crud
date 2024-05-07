# CONTAINER FOR BUILDING BINARY
FROM golang:1.21 AS build

# INSTALL DEPENDENCIES
# RUN go install github.com/gobuffalo/packr/v2/packr2@v2.8.3
COPY . .
RUN go mod ridy

# BUILD BINARY
# COPY . /src
RUN make build
RUN pwd

# CONTAINER FOR RUNNING BINARY
FROM alpine:3.18
COPY --from=build /src/dist/zkevm-node /app/zkevm-node
RUN apk update && apk add postgresql15-client
EXPOSE 8123
CMD ["/bin/sh", "-c", "/app/zkevm-node run"]