FROM golang:1.23.0 AS build

COPY go.mod go.sum /app/
WORKDIR /app
RUN go mod download

# Copy the source code
COPY . /app

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o email-service

FROM scratch
COPY --from=build /app/email-service /

# Run
CMD ["/email-service"] 
