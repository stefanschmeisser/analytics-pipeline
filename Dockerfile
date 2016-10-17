# Base this docker container off the official golang docker image.
# Docker containers inherit everything from their base.
FROM golang:1.6

# Copy the example-app directory (where the Dockerfile lives) into the container.
COPY . /go/src/go-websockert-server

WORKDIR /go/src/go-websockert-server

# Download and install any required third party dependencies into the container.
RUN go-wrapper download
RUN go-wrapper install

# Set the PORT environment variable inside the container
ENV PORT 3030

# Expose port 3000 to the host so we can access the gin proxy
EXPOSE 3030

CMD ["go-wrapper", "run"]
