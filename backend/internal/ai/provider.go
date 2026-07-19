package ai

import "context"

type Analysis struct {
	Summary string   `json:"summary"`
	Labels  []string `json:"labels"`
}

// InferenceProvider abstracts the model backend. Swapping Gemma for Kimi later
// is a config change (AI_PROVIDER=...), not a rewrite.
type InferenceProvider interface {
	AnalyzeText(ctx context.Context, text string) (Analysis, error)
	Transcribe(ctx context.Context, audio []byte) (string, error)
}

// GemmaProvider is a stub. ponytail: returns canned output — wire to a real
// Gemma endpoint when the inference service exists.
type GemmaProvider struct{ Model string }

func (g GemmaProvider) AnalyzeText(_ context.Context, _ string) (Analysis, error) {
	return Analysis{Summary: "stub analysis", Labels: []string{"stub"}}, nil
}

func (g GemmaProvider) Transcribe(_ context.Context, _ []byte) (string, error) {
	return "", nil
}

// NewProvider selects the provider by name (from AI_PROVIDER).
func NewProvider(name string) InferenceProvider {
	switch name {
	case "gemma":
		return GemmaProvider{Model: "gemma-2"}
	// case "kimi":
	// 	return KimiProvider{...} // add when needed — no interface change
	default:
		return GemmaProvider{Model: "gemma-2"}
	}
}
