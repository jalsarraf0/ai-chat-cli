# LLM Usage

`ai-chat ask` streams chat completions to stdout. Set your OpenAI key in `config.yaml` (`openai_api_key`) or `AI_CHAT_API_KEY`.

| Flag | Config key | Purpose |
|------|------------|---------|
| `--model` | `model` | model ID |
| `--temperature` | `temperature` | sampling temperature |
| `--max_tokens` | `max_tokens` | token limit |

`AICHAT_BASE_URL` overrides the OpenAI API host. `AICHAT_TIMEOUT` sets the request timeout.
