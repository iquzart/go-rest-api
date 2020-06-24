# Start from the latest golang base image
FROM golang:alpine AS build-env

WORKDIR /src

COPY . .

# Download dependencies 
RUN go mod download

# Build the Go app
RUN go build -o app


# final stage
FROM alpine

# Add Maintainer Info
LABEL maintainer="Muhammed Iqbal <iquzart@hotmail.com>"

ENV GIN_MODE=release
ENV PORT=3000

WORKDIR /app

# Copy app fron build-env
COPY --from=build-env /src/app /app/

# Expose port 8080 to the outside world
EXPOSE 3000

# Run the executable
CMD ["./app"]
