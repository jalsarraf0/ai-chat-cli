FROM golang:1.24 AS build
WORKDIR /src
COPY . .
RUN go build -o ai-chat-cli ./cmd/ai-chat-cli

FROM scratch
COPY --from=build /src/ai-chat-cli /ai-chat-cli
ENTRYPOINT ["/ai-chat-cli"]
