package openai

type OpenaiClient struct {
	apikey string
}

// Returns vector embeddings for given text
func (c *OpenaiClient) Embed(text string) ([]float64, error) {
	key := c.apikey

	return createEmbedding(text, key)
}

func (c *OpenaiClient) Complete(text string) (string, error) {
	key := c.apikey

	return createCompletion(text, key)
}

func NewOpenaiClient(key string) OpenaiClient {
	return OpenaiClient{
		apikey: key,
	}
}
