# LLM Usage

`ai-chat ask` streams chat completions to stdout. Set your OpenAI key in `ai-chat.yaml` (`openai_api_key`) or `AI_CHAT_API_KEY`.

| Flag | Config key | Purpose |
|------|------------|---------|
| `--model` | `model` | model ID |
| `--temperature` | `temperature` | sampling temperature |
| `--max_tokens` | `max_tokens` | token limit |

`AICHAT_BASE_URL` overrides the OpenAI API host. `AICHAT_TIMEOUT` sets the request timeout.

## Unit vs Live

Run offline tests with:

```bash
go test -tags unit ./...
```

Live tests hit the real API only when `AI_CHAT_API_KEY` is set. Use `make live-openai-test` to run them.
