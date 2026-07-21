package ai

import (
	"context"
	"encoding/base64"
)

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
