package teal

import (
	"context"
	"time"

	"github.com/COG-GTM/teal-agents/pkg/types"
)

type TaskStatus string

const (
	TaskStatusRunning   TaskStatus = "Running"
	TaskStatusPaused    TaskStatus = "Paused"
	TaskStatusCompleted TaskStatus = "Completed"
	TaskStatusFailed    TaskStatus = "Failed"
	TaskStatusCanceled  TaskStatus = "Canceled"
)

type AgentTaskItem struct {
	TaskID           string                   `json:"task_id"`
	Role             string                   `json:"role"`
	Item             types.MultiModalItem     `json:"item"`
	RequestID        string                   `json:"request_id"`
	Updated          time.Time                `json:"updated"`
	PendingToolCalls []map[string]interface{} `json:"pending_tool_calls,omitempty"`
	ChatHistory      []types.HistoryMessage   `json:"chat_history,omitempty"`
}

type AgentTask struct {
	TaskID      string          `json:"task_id"`
	SessionID   string          `json:"session_id"`
	UserID      string          `json:"user_id"`
	Items       []AgentTaskItem `json:"items"`
	CreatedAt   time.Time       `json:"created_at"`
	LastUpdated time.Time       `json:"last_updated"`
	Status      TaskStatus      `json:"status"`
}

type HITLResponse struct {
	SessionID    string                   `json:"session_id"`
	TaskID       string                   `json:"task_id"`
	RequestID    string                   `json:"request_id"`
	ToolCalls    []map[string]interface{} `json:"tool_calls"`
	ApprovalURL  string                   `json:"approval_url"`
	RejectionURL string                   `json:"rejection_url"`
}

type ResumeRequest struct {
	Action string `json:"action"`
}

type TealAgentConfig struct {
	Agent AgentConfig `yaml:"agent"`
}

type AgentConfig struct {
	Name          string   `yaml:"name"`
	Model         string   `yaml:"model"`
	SystemPrompt  string   `yaml:"system_prompt"`
	Temperature   *float64 `yaml:"temperature,omitempty"`
	Plugins       []string `yaml:"plugins,omitempty"`
	RemotePlugins []string `yaml:"remote_plugins,omitempty"`
}

type TealAgent interface {
	Invoke(ctx context.Context, authToken string, message types.UserMessage) (interface{}, error)

	InvokeStream(ctx context.Context, authToken string, message types.UserMessage) (<-chan interface{}, error)

	Resume(ctx context.Context, authToken string, requestID string, action ResumeRequest, stream bool) (interface{}, error)
}

type StateManager interface {
	Create(ctx context.Context, task *AgentTask) error

	Load(ctx context.Context, taskID string) (*AgentTask, error)

	LoadByRequestID(ctx context.Context, requestID string) (*AgentTask, error)

	Update(ctx context.Context, task *AgentTask) error
}

type Builder interface {
	Build(config TealAgentConfig, stateManager StateManager) (TealAgent, error)
}
