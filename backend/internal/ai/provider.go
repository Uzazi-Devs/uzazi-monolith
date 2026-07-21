package ai

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func NewGemmaProvider() GemmaProvider {
	return GemmaProvider{
		Model:  "gemma-4-26b-a4b-it", // pick the size that fits your latency/cost needs
		APIKey: os.Getenv("GEMINI_API_KEY"),
		client: &http.Client{Timeout: 30 * time.Second},
	}
}

const geminiBaseURL = "https://generativelanguage.googleapis.com/v1beta/models"

func (g GemmaProvider) call(ctx context.Context, req genContentRequest) (string, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return "", fmt.Errorf("marshal request: %w", err)
	}

	url := fmt.Sprintf("%s/%s:generateContent", geminiBaseURL, g.Model)
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(body))
	if err != nil {
		return "", fmt.Errorf("build request: %w", err)
	}
	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("x-goog-api-key", g.APIKey)

	resp, err := g.client.Do(httpReq)
	if err != nil {
		return "", fmt.Errorf("call gemini api: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("read response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("gemini api error (%d): %s", resp.StatusCode, string(respBody))
	}

	var parsed genContentResponse
	if err := json.Unmarshal(respBody, &parsed); err != nil {
		return "", fmt.Errorf("parse response: %w", err)
	}
	if len(parsed.Candidates) == 0 || len(parsed.Candidates[0].Content.Parts) == 0 {
		return "", fmt.Errorf("empty response from gemini api")
	}
	return parsed.Candidates[0].Content.Parts[0].Text, nil
}
