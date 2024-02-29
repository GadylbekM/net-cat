# Use a Golang base image
FROM golang:1.17-alpine

WORKDIR /app

COPY . /app

# Build the Golang application inside the container
RUN go build -o chat-app

# Expose the port your chat application is listening on
EXPOSE 8000

# Run your chat application
CMD ["./chat-app"]
