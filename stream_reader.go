package openrouter

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"

	utils "github.com/casibase/go-openrouter/internal"
)

var (
	ErrTooManyEmptyStreamMessages = errors.New("stream has sent too many empty messages")
)

type streamReader struct {
	emptyMessagesLimit uint
	isFinished         bool

	reader         *bufio.Reader
	response       *http.Response
	errAccumulator utils.ErrorAccumulator
	unmarshaler    utils.Unmarshaler
}

func (stream *streamReader) Recv() (response *ChatCompletionResponse, err error) {
	if stream.isFinished {
		err = io.EOF
		return
	}

	response, err = stream.processLines()
	return
}

func (stream *streamReader) processLines() (*ChatCompletionResponse, error) {
	var emptyMessagesCount uint

	for {
		rawLine, readErr := stream.reader.ReadBytes('\n')
		if readErr != nil {
			respErr := stream.unmarshalError()
			if respErr != nil {
				return nil, fmt.Errorf("error, %w", respErr.Error)
			}
			return nil, readErr
		}

		var headerData = []byte("data:")
		noSpaceLine := bytes.TrimSpace(rawLine)
		if !bytes.HasPrefix(noSpaceLine, headerData) {
			writeErr := stream.errAccumulator.Write(noSpaceLine)
			if writeErr != nil {
				return nil, writeErr
			}
			emptyMessagesCount++
			if emptyMessagesCount > stream.emptyMessagesLimit {
				return nil, ErrTooManyEmptyStreamMessages
			}

			continue
		}

		noPrefixLine := bytes.TrimSpace(bytes.TrimPrefix(noSpaceLine, headerData))

		if string(noPrefixLine) == "[DONE]" {
			stream.isFinished = true
			return nil, io.EOF
		}

		var response ChatCompletionResponse
		unmarshalErr := stream.unmarshaler.Unmarshal(noPrefixLine, &response)
		if unmarshalErr != nil {
			return nil, unmarshalErr
		}

		return &response, nil
	}
}

func (stream *streamReader) unmarshalError() (errResp *ErrorResponse) {
	errBytes := stream.errAccumulator.Bytes()
	if len(errBytes) == 0 {
		return
	}

	err := stream.unmarshaler.Unmarshal(errBytes, &errResp)
	if err != nil {
		errResp = nil
	}

	return
}

func (stream *streamReader) Close() {
	stream.response.Body.Close()
}
