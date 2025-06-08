FROM golang:1.24 AS build
WORKDIR /app

# first copy go module files to leverage Docker layer caching
COPY go.mod go.sum ./
RUN go mod download

# then copy the source
COPY . .
RUN go build -o ai-chat-cli ./cmd/ai-chat-cli

FROM scratch
COPY --from=build /app/ai-chat-cli /ai-chat-cli
ENTRYPOINT ["/ai-chat-cli"]
