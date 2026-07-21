package ai

import (
	"context"
	"encoding/json"
	"fmt"
)

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
