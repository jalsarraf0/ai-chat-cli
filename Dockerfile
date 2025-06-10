FROM scratch
COPY ai-chat-cli /ai-chat-cli
ENTRYPOINT ["/ai-chat-cli"]
