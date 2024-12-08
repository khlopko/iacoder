package core

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Chat struct {
	coder *Coder
	history *History
}

func NewChat(coder *Coder, errorChan chan<- error) *Chat {
	history := NewHistory()
	go func() {
		err := history.Prepare()
		errorChan <- err
	}()
	return &Chat{coder, history}
}

type ParsedMessage struct {
	Path    string `json:"path"`
	Content string `json:"content"`
}

type TaskResult struct {
	Message *string
	Err error
}

func (c *Chat) ExecuteNewTask(task string) ([]TaskResult, error) {
	messages, err := c.coder.SendMessage(task)
	if err != nil {
		// TODO: Setup logging system
		//fmt.Printf("Couldn't get anything from AI with error: %+v\n", err)
		return nil, err
	}

	if len(messages) == 0 {
		fmt.Println("AI didn't provide any response :(")
		return nil, errors.New("Empty response")
	}

	result := []TaskResult{}

	for _, message := range messages {

		var parsedMessages []ParsedMessage
		err = json.Unmarshal([]byte(message), &parsedMessages)
		if err != nil {
			decodeErr := errors.New(fmt.Sprintf("Decoding failed: %+v", err))
			result = append(result, TaskResult{&message, decodeErr})
		}

		for _, parsedMessage := range parsedMessages {
			err = os.WriteFile(parsedMessage.Path, []byte(parsedMessage.Content), 777)
			if err != nil {
				writeErr := errors.New(fmt.Sprintf("Write at %s failed: %+v", parsedMessage.Path, err))
				result = append(result, TaskResult{&message, writeErr})
			} else {
				successMsg := fmt.Sprintf("Saved to %s", parsedMessage.Path)
				result = append(result, TaskResult{&successMsg, nil})
			}
		}

	}

	return result, nil
}
