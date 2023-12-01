package openrouter

import (
	"bufio"
	"context"
	utils "github.com/casibase/go-openrouter/internal"
)

type ChatCompletionStream struct {
	streamReader
}

// CreateChatCompletionStream — API call to create a chat completion w/ streaming
// support. It sets whether to stream back partial progress. If set, tokens will be
// sent as data-only server-sent events as they become available, with the
// stream terminated by a data: [DONE] message.
func (c *Client) CreateChatCompletionStream(
	ctx context.Context,
	request *ChatCompletionRequest,
) (stream *ChatCompletionStream, err error) {
	urlSuffix := "/chat/completions"
	request.Model = wrapperModels[request.Model]
	if !checkSupportsModel(request.Model) {
		err = ErrCompletionUnsupportedModel
		return
	}
	request.Stream = true
	req, err := c.newStreamRequest(ctx, "POST", urlSuffix, request)
	if err != nil {
		return
	}

	resp, err := c.config.HTTPClient.Do(req) //nolint:bodyclose // body is closed in stream.Close()
	if err != nil {
		return
	}
	if isFailureStatusCode(resp) {
		return nil, c.handleErrorResp(resp)
	}

	stream = &ChatCompletionStream{
		streamReader: streamReader{
			emptyMessagesLimit: c.config.EmptyMessagesLimit,
			reader:             bufio.NewReader(resp.Body),
			response:           resp,
			errAccumulator:     utils.NewErrorAccumulator(),
			unmarshaler:        &utils.JSONUnmarshaler{},
		},
	}
	return
}
