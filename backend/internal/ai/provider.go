package ai

import (
	"bytes"
	"context"
	"encoding/base64"
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

func (g GemmaProvider) AnalyzeText(ctx context.Context, text string) (Analysis, error) {
	req := genContentRequest{
		SystemInstruction: &content{Parts: []part{{
			Text: `Analyze the following journal entry for signs of postpartum depression risk.
Respond ONLY with JSON in this exact shape: {"summary": "...", "labels": ["..."]}`,
		}}},
		Contents: []content{{Parts: []part{{Text: text}}}},
	}

	raw, err := g.call(ctx, req)
	if err != nil {
		return Analysis{}, err
	}

	var result Analysis
	if err := json.Unmarshal([]byte(raw), &result); err != nil {
		return Analysis{}, fmt.Errorf("model did not return valid JSON: %w", err)
	}
	return result, nil
}

func (g GemmaProvider) Transcribe(ctx context.Context, audio []byte) (string, error) {
	req := genContentRequest{
		Contents: []content{{Parts: []part{
			{Text: "Transcribe this audio exactly, no commentary."},
			{InlineData: &inlineData{
				MimeType: "audio/mp3", // match whatever format the Android/web client actually records
				Data:     base64.StdEncoding.EncodeToString(audio),
			}},
		}}},
	}
	return g.call(ctx, req)
}

// NewProvider selects the provider by name (from AI_PROVIDER).
func NewProvider(name string) InferenceProvider {
	switch name {
	case "gemma":
		return NewGemmaProvider()
	// case "kimi":
	// 	return KimiProvider{...} // add when needed — no interface change
	default:
		return NewGemmaProvider()
	}
}
