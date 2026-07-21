package ai

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
