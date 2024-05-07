# CONTAINER FOR BUILDING BINARY
FROM golang:1.22.2 AS build

# INSTALL DEPENDENCIES
# RUN go install github.com/gobuffalo/packr/v2/packr2@v2.8.3
COPY go.mod go.sum /src/
RUN cd /src && go mod download
# COPY . .
# RUN go mod download


# BUILD BINARY
COPY . /src
RUN cd /src && make build-alpine

# CONTAINER FOR RUNNING BINARY
FROM alpine:3.18
COPY --from=build /src/build/cmd /app/gocrud
RUN apk update 
# && apk add postgresql15-client
EXPOSE 8000
CMD ["/bin/sh", "-c", "/app/gocrud start"]