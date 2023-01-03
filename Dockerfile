FROM golang:alpine as build

RUN mkdir /app

WORKDIR /app

# Copy all files
COPY . .

# Compile the source code into a binary executable file
RUN go build -o main ./cmd/api/.

FROM alpine:latest

WORKDIR /app

# Copy the binary file and migrations folder from the build stage
COPY --from=build /app/main /app/main

# Run the binary file
ENTRYPOINT [ "/app/main" ]

EXPOSE 80
