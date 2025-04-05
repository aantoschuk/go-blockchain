# Build stage
FROM golang:1.24.2-bookworm AS build

RUN groupadd -g 1234 bccontainergroup && \
    useradd -m -u 1234 -g bccontainergroup bccontaineruser

WORKDIR /home/bccontaineruser

# As there is no actually dependencies copy everything
# Otherwise copy go.mod go.sum ./
COPY . .

# Same as the previous one - nothing to install
# Download dependencies
# RUN go mod download

# Copy the rest of the application
# COPY . .

# Build the Go application
RUN go build -o /home/bccontaineruser/app .

# Final stage - copy the binary into a clean image
FROM golang:1.24.2-bookworm AS final

RUN groupadd -g 1234 bccontainergroup && \
    useradd -m -u 1234 -g bccontainergroup bccontaineruser

WORKDIR /home/bccontaineruser

# Copy the compiled binary from the build stage
COPY --from=build /home/bccontaineruser/app .

# Ensure the app runs as the non-root user
USER bccontaineruser

CMD ["./app"]

