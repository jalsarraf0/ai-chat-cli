package openai

import (
	"context"
	"os"

	gpt "github.com/sashabaranov/go-openai"
)

// Stream returns a channel streaming tokens from the OpenAI API using the
// default key from the environment. It closes the channel when the stream
// finishes or on error.
func Stream(prompt string) <-chan string {
	ch := make(chan string)
	go func() {
		defer close(ch)
		key := os.Getenv("OPENAI_API_KEY")
		if key == "" {
			return
		}
		client := gpt.NewClient(key)
		req := gpt.ChatCompletionRequest{
			Model:  gpt.GPT3Dot5Turbo,
			Stream: true,
			Messages: []gpt.ChatCompletionMessage{
				{Role: gpt.ChatMessageRoleUser, Content: prompt},
			},
		}
		stream, err := client.CreateChatCompletionStream(context.Background(), req)
		if err != nil {
			return
		}
		defer stream.Close()
		for {
			resp, err := stream.Recv()
			if err != nil {
				return
			}
			if len(resp.Choices) == 0 {
				continue
			}
			ch <- resp.Choices[0].Delta.Content
		}
	}()
	return ch
}
