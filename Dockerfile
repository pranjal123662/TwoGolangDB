# Use the official golang image as a base image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /go/src/app

# Copy the local package files to the container's workspace
COPY . .

# Build the TwoDB app
RUN go build -o TwoDB .

# Clone your git repository
RUN git clone https://github.com/pranjal123662/TwoGolangDB.git 
# Expose port 7456 to the outside world
EXPOSE 7456

# Command to run the executable
CMD ["./TwoDB"]
