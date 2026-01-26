package opencode

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

// OpenCode api request types

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type OpenCodeRequest struct {
	Model         string    `json:"model"`
	MaxTokens     int       `json:"max_tokens"`
	SystemMessage string    `json:"system_message"`
	Message       []Message `json:"message"`
}

// OpenCode response - returns raw JSON string for direct parsing
type OpenCodeResponse struct {
	JSONResponse string `json:"-"` // Raw JSON response from the model
}

// OpenCode client code - Direct HTTP implementation for Zen API

type Client struct {
	apiKey     string
	baseURL    string
	httpClient *http.Client
}

func NewClient(apiKey string) *Client {
	return &Client{
		apiKey:  apiKey,
		baseURL: "https://opencode.ai/zen/v1",
		httpClient: &http.Client{
			Timeout: 5 * time.Minute,
		},
	}
}

func (c *Client) CreateMessage(ctx context.Context, req OpenCodeRequest) (*OpenCodeResponse, error) {
	// Convert to OpenAI-compatible format for Grok Code Fast 1
	openAIReq := c.convertToOpenAIFormat(req)

	// Make the API call
	jsonData, err := json.Marshal(openAIReq)
	if err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequestWithContext(ctx, "POST", c.baseURL+"/chat/completions", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("Authorization", "Bearer "+c.apiKey)

	resp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed with status %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Parse OpenAI response and extract content
	var openAIResp struct {
		Choices []struct {
			Message struct {
				Content string `json:"content"`
			} `json:"message"`
		} `json:"choices"`
	}

	if err := json.Unmarshal(body, &openAIResp); err != nil {
		return nil, err
	}

	if len(openAIResp.Choices) == 0 {
		return nil, errors.New("no choices in response")
	}

	return &OpenCodeResponse{
		JSONResponse: openAIResp.Choices[0].Message.Content,
	}, nil
}

// convertToOpenAIFormat converts our request format to OpenAI-compatible format
func (c *Client) convertToOpenAIFormat(req OpenCodeRequest) map[string]any {
	messages := []map[string]string{}

	// Add system message if present
	if req.SystemMessage != "" {
		messages = append(messages, map[string]string{
			"role":    "system",
			"content": req.SystemMessage,
		})
	}

	// Add user messages
	for _, msg := range req.Message {
		messages = append(messages, map[string]string{
			"role":    msg.Role,
			"content": msg.Content,
		})
	}

	return map[string]any{
		"model":       "big-pickle",
		"messages":    messages,
		"max_tokens":  req.MaxTokens,
		"temperature": 0.1, // Low temperature for consistent structured output
	}
}
