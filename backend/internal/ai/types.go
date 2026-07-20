package ai

import (
	"context"
	"net/http"
)

// Analysis is the structured output of a text-analysis call. It is a
// screening signal for a health worker to review, never a diagnosis — do
// not add a field here that implies a clinical conclusion (e.g. "hasPPD").
type Analysis struct {
	Summary string   `json:"summary"`
	Labels  []string `json:"labels"`
}

// InferenceProvider abstracts the model backend. Swapping Gemma for Kimi
// later is a config change (AI_PROVIDER=...), not a rewrite — every
// concrete provider (GemmaProvider, KimiProvider, ...) implements this and
// nothing outside the ai package should depend on a concrete type.
type InferenceProvider interface {
	// AnalyzeText returns labels drawn from the fixed vocabulary in labels.go.
	// Implementations must constrain the model to that vocabulary via the
	// prompt in prompts.go — callers rely on labels.go's isConcerningLabel
	// matching against it.
	AnalyzeText(ctx context.Context, text string) (Analysis, error)

	// Transcribe returns plain text transcription of the given audio bytes.
	// Implementations decide their own accepted audio encoding; callers in
	// the media package are responsible for sending a compatible format.
	Transcribe(ctx context.Context, audio []byte) (string, error)
}

// GemmaProvider calls the Gemini API's hosted Gemma models.
type GemmaProvider struct {
	Model  string
	APIKey string
	client *http.Client
}

// --- request/response shapes for generateContent ---

type genContentRequest struct {
	Contents          []content `json:"contents"`
	SystemInstruction *content  `json:"systemInstruction,omitempty"`
}

type content struct {
	Parts []part `json:"parts"`
}

type part struct {
	Text       string      `json:"text,omitempty"`
	InlineData *inlineData `json:"inlineData,omitempty"`
}

type inlineData struct {
	MimeType string `json:"mimeType"`
	Data     string `json:"data"` // base64
}

type genContentResponse struct {
	Candidates []struct {
		Content struct {
			Parts []struct {
				Text string `json:"text"`
			} `json:"parts"`
		} `json:"content"`
	} `json:"candidates"`
}
