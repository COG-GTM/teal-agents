package types

import (
	"context"
)

type TokenUsage struct {
	CompletionTokens int `json:"completion_tokens"`
	PromptTokens     int `json:"prompt_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

type ExtraData map[string]interface{}

type InvokeResponse struct {
	SessionID    string      `json:"session_id,omitempty"`
	Source       string      `json:"source,omitempty"`
	RequestID    string      `json:"request_id,omitempty"`
	TokenUsage   TokenUsage  `json:"token_usage"`
	ExtraData    ExtraData   `json:"extra_data,omitempty"`
	OutputRaw    string      `json:"output_raw,omitempty"`
	OutputParsed interface{} `json:"output_pydantic,omitempty"`
}

type PartialResponse struct {
	SessionID     string `json:"session_id,omitempty"`
	Source        string `json:"source,omitempty"`
	RequestID     string `json:"request_id,omitempty"`
	OutputPartial string `json:"output_partial"`
}

type IntermediateTaskResponse struct {
	TaskNo   int            `json:"task_no"`
	TaskName string         `json:"task_name"`
	Response InvokeResponse `json:"response"`
}

type HistoryMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ContentType string

const (
	ContentTypeText  ContentType = "text"
	ContentTypeImage ContentType = "image"
)

type MultiModalItem struct {
	ContentType ContentType `json:"content_type"`
	Content     string      `json:"content"`
}

type UserMessage struct {
	SessionID   string            `json:"session_id,omitempty"`
	TaskID      string            `json:"task_id,omitempty"`
	Items       []MultiModalItem  `json:"items"`
	UserContext map[string]string `json:"user_context,omitempty"`
}

type BaseHandler interface {
	Invoke(ctx context.Context, inputs map[string]interface{}) (*InvokeResponse, error)
	InvokeStream(ctx context.Context, inputs map[string]interface{}) (<-chan StreamResponse, error)
}

type StreamResponse struct {
	Partial   *PartialResponse
	Final     *InvokeResponse
	Error     error
	IsPartial bool
	IsFinal   bool
}

type Config struct {
	APIVersion  string                 `yaml:"apiVersion"`
	Name        string                 `yaml:"name,omitempty"`
	ServiceName string                 `yaml:"service_name,omitempty"`
	Version     string                 `yaml:"version"`
	Description string                 `yaml:"description,omitempty"`
	InputType   string                 `yaml:"input_type,omitempty"`
	OutputType  string                 `yaml:"output_type,omitempty"`
	Spec        map[string]interface{} `yaml:"spec,omitempty"`
}
