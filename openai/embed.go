package openai

import (
	"encoding/json"
	"fmt"
)

// Actual endpoint

const embedding_url = base_url + "/v1/embeddings"

// Types for embedding api response

type _OpenAIEmbeddingData struct {
	Embedding []float64 `json:"embedding"`
	Index     int       `json:"index"`
	Object    string    `json:"object"`
}

type _OpenAIEmbeddingUseage struct {
	PromptTokens int `json:"prompt_tokens"`
	TotalTokens  int `json:"total_tokens"`
}

type _OpenAIEmbeddingResponse struct {
	Data   []_OpenAIEmbeddingData `json:"data"`
	Object string                 `json:"object"`
	Model  string                 `json:"model"`
	Usage  _OpenAIEmbeddingUseage `json:"usage"`
}

// Turns string into proper request body format
func makeEmbeddingRequestBody(text string) []byte {
	reqBody, err := json.Marshal(map[string]interface{}{
		"input": text,
		"model": "text-embedding-ada-002", // Best + cheapest model
	})

	if err != nil {
		fmt.Println(err)
	}

	return reqBody
}

// Parses response body into struct
func parseEmbeddingResponseBody(res []byte) _OpenAIEmbeddingResponse {
	var jsonRes _OpenAIEmbeddingResponse

	json.Unmarshal(res, &jsonRes)

	return jsonRes
}

// Makes embedding from text using openai api
func createEmbedding(text string, key string) ([]float64, error) {
	// make body
	b := makeEmbeddingRequestBody(text)

	// Call openai api
	_, resBody, err := openAIhttp.request("POST", embedding_url, b, key) // maybe todo: catch error here

	response := parseEmbeddingResponseBody(resBody)

	return response.Data[0].Embedding, err // Extract out just the embedding, but
	// could get any other info from the response types defined above
}
