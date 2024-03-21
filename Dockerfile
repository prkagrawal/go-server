# # Use an official Python runtime as a parent image.
# FROM golang:1.22

# # Set the working directory to /app
# WORKDIR /usr/src/app

# # pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
# COPY go.mod go.sum ./
# RUN go mod download && go mod verify

# COPY . .
# RUN go build -v -o /usr/local/bin/app ./...

# # Make port 80 available to the world outside this container
# EXPOSE 8080

# # # Define environment variable
# # ENV NAME World

# # Run app when the container launches
# CMD ["app"]

# # Start from golang base image
# FROM golang:1.22

# # Install git.
# # Git is required for fetching the dependencies.
# # RUN apk update && apk add --no-cache git && apk add --no-cach bash && apk add build-base

# # Setup folders
# RUN mkdir /app
# WORKDIR /app

# # Copy the source from the current directory to the working Directory inside the container
# COPY . .
# # COPY .env .

# # Download all the dependencies
# RUN go get -d -v ./...

# # Install the package
# RUN go install -v ./...

# # Build the Go app
# RUN go build -o /build

# # Expose port 8080 to the outside world
# EXPOSE 8080

# # Run the executable
# CMD [ "/build" ]


FROM golang:1.22

WORKDIR /app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o main .

CMD ["/app/main"]