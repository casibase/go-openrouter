package openrouter

import (
	"context"
	"testing"
)

func TestClient_CreateChatCompletion(t *testing.T) {
	client, _ := NewClient("", "", "")

	req := &ChatCompletionRequest{
		Model: "claude-2",
		Messages: []ChatCompletionMessage{
			{
				Role:    ChatMessageRoleSystem,
				Content: "You are a helpful assistant.",
			},
			{
				Role:    ChatMessageRoleUser,
				Content: "what is today",
			},
		},
		Stream:      false,
		Temperature: nil,
		TopP:        nil,
	}

	t.Log(client.CreateChatCompletion(context.Background(), req))
	//
	//r, err := client.CreateChatCompletionStream(context.Background(), req)
	//if err != nil {
	//	t.Error(err)
	//}
	//t.Log(r)
	//for {
	//	fmt.Println(1)
	//	r, err := r.Recv()
	//	if err != nil {
	//		fmt.Println(err.Error())
	//		if errors.Is(err, io.EOF) {
	//			fmt.Println(1)
	//			break
	//		}
	//		t.Error(err)
	//	}
	//	t.Logf("%#v", r.Choices)
	//}
}
