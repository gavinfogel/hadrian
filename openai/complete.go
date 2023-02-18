package openai

import (
	"encoding/json"
	"fmt"
)

const completion_url = base_url + "/v1/completions"

// "id": "cmpl-uqkvlQyYK7bGYrRHQ0eXlWi7",
//   "object": "text_completion",
//   "created": 1589478378,
//   "model": "text-davinci-003",
//   "choices": [
//     {
//       "text": "\n\nThis is indeed a test",
//       "index": 0,
//       "logprobs": null,
//       "finish_reason": "length"
//     }
//   ],
//   "usage": {
//     "prompt_tokens": 5,
//     "completion_tokens": 7,
//     "total_tokens": 12
//   }

type _CompletionResponse struct {
	Id      string `json:"id"`
	Object  string `json:"object"`
	Created string `json:"created"`
	Model   string `json:"model"`
	Choices []struct {
		Text         string `json:"text"`
		Index        int    `json:"index"`
		Logprobs     string `json:"logprobs"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
}

// Turns string into proper request body format
func makeCompletionRequestBody(prompt string) []byte {
	reqBody, err := json.Marshal(map[string]interface{}{
		"model":      "text-davinci-003", // Most expensive best model
		"prompt":     prompt,
		"max_tokens": 2048,
	})

	if err != nil {
		fmt.Println(err)
	}

	return reqBody
}

func parseCompletionResponseBody(res []byte) _CompletionResponse {
	var jsonRes _CompletionResponse

	json.Unmarshal(res, &jsonRes)

	return jsonRes
}

func createCompletion(text string, key string) (string, error) {
	b := makeCompletionRequestBody(text)

	_, resBody, err := openAIhttp.request("POST", completion_url, b, key) // maybe todo: catch error here

	response := parseCompletionResponseBody(resBody)

	if len(response.Choices) == 0 {
		return "", err
	}

	return response.Choices[0].Text, err
}
