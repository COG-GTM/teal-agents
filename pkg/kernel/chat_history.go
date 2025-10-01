package kernel

import (
	"time"
)

type AuthorRole string

const (
	AuthorRoleUser      AuthorRole = "user"
	AuthorRoleAssistant AuthorRole = "assistant"
	AuthorRoleSystem    AuthorRole = "system"
	AuthorRoleTool      AuthorRole = "tool"
)

type ContentItem interface {
	GetType() string
	GetContent() interface{}
}

type TextContent struct {
	Text string `json:"text"`
}

func (t TextContent) GetType() string         { return "text" }
func (t TextContent) GetContent() interface{} { return t.Text }

type ImageContent struct {
	DataURI string `json:"data_uri"`
	URL     string `json:"url,omitempty"`
}

func (i ImageContent) GetType() string         { return "image" }
func (i ImageContent) GetContent() interface{} { return i.DataURI }

type ChatMessageContent struct {
	Role      AuthorRole    `json:"role"`
	Items     []ContentItem `json:"items"`
	Name      string        `json:"name,omitempty"`
	Timestamp time.Time     `json:"timestamp,omitempty"`
}

type ChatHistory struct {
	Messages []ChatMessageContent
}

func NewChatHistory() *ChatHistory {
	return &ChatHistory{
		Messages: make([]ChatMessageContent, 0),
	}
}

func (h *ChatHistory) AddMessage(message ChatMessageContent) {
	h.Messages = append(h.Messages, message)
}

func (h *ChatHistory) AddUserMessage(text string) {
	h.AddMessage(ChatMessageContent{
		Role:      AuthorRoleUser,
		Items:     []ContentItem{TextContent{Text: text}},
		Timestamp: time.Now(),
	})
}

func (h *ChatHistory) AddAssistantMessage(text string) {
	h.AddMessage(ChatMessageContent{
		Role:      AuthorRoleAssistant,
		Items:     []ContentItem{TextContent{Text: text}},
		Timestamp: time.Now(),
	})
}

func (h *ChatHistory) GetMessages() []ChatMessageContent {
	return h.Messages
}
