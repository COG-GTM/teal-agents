package kernel

import (
	"context"
)

type ModelType string

const (
	ModelTypeOpenAI    ModelType = "openai"
	ModelTypeAnthropic ModelType = "anthropic"
	ModelTypeGoogle    ModelType = "google"
)

type ChatMessage struct {
	Role    string
	Content string
}

type CompletionRequest struct {
	Messages    []ChatMessage
	Model       string
	Temperature float64
	MaxTokens   int
	Stream      bool
}

type CompletionResponse struct {
	Content          string
	TokensUsed       int
	FinishReason     string
	PromptTokens     int
	CompletionTokens int
}

type ChatCompletionClient interface {
	Complete(ctx context.Context, req *CompletionRequest) (*CompletionResponse, error)

	CompleteStream(ctx context.Context, req *CompletionRequest) (<-chan string, error)

	GetModelType() ModelType
}

type Builder interface {
	BuildForModel(modelName string) (ChatCompletionClient, error)
}
