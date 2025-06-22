package airepo

import (
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty/v2"
)

type GeminiClient struct {
	apiKey string
	client *resty.Client
}

func NewGeminiClient(apiKey string) *GeminiClient {
	return &GeminiClient{
		apiKey: apiKey,
		client: resty.New(),
	}
}

// Implements usecase.GeminiAI
func (g *GeminiClient) GenerateText(prompt string) (string, error) {
	reqBody := map[string]interface{}{
		"contents": []map[string]interface{}{
			{
				"parts": []map[string]string{
					{"text": prompt},
				},
			},
		},
	}

	resp, err := g.client.R().
		SetHeader("Content-Type", "application/json").
		SetQueryParam("key", g.apiKey).
		SetBody(reqBody).
		Post("https://generativelanguage.googleapis.com/v1beta/models/gemini-pro:generateContent")

	if err != nil {
		return "", err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		return "", err
	}

	// Parse response
	candidates, ok := result["candidates"].([]interface{})
	if !ok || len(candidates) == 0 {
		return "", fmt.Errorf("no candidates returned")
	}

	content, ok := candidates[0].(map[string]interface{})["content"].(map[string]interface{})
	if !ok {
		return "", fmt.Errorf("content missing in response")
	}

	parts, ok := content["parts"].([]interface{})
	if !ok || len(parts) == 0 {
		return "", fmt.Errorf("parts missing in response")
	}

	text, ok := parts[0].(map[string]interface{})["text"].(string)
	if !ok {
		return "", fmt.Errorf("text missing in response")
	}

	return text, nil
}
