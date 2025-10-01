package state

import (
	"context"
	"time"
)

type TaskStatus string

const (
	TaskStatusRunning   TaskStatus = "Running"
	TaskStatusPaused    TaskStatus = "Paused"
	TaskStatusCompleted TaskStatus = "Completed"
	TaskStatusFailed    TaskStatus = "Failed"
)

type TaskState struct {
	TaskID    string                   `json:"task_id"`
	SessionID string                   `json:"session_id"`
	UserID    string                   `json:"user_id"`
	Messages  []map[string]interface{} `json:"messages"`
	Status    TaskStatus               `json:"status"`
	CreatedAt time.Time                `json:"created_at"`
	UpdatedAt time.Time                `json:"updated_at"`
	Metadata  map[string]interface{}   `json:"metadata"`
}

type RequestState struct {
	RequestID string                 `json:"request_id"`
	TaskID    string                 `json:"task_id"`
	Status    TaskStatus             `json:"status"`
	CreatedAt time.Time              `json:"created_at"`
	UpdatedAt time.Time              `json:"updated_at"`
	Metadata  map[string]interface{} `json:"metadata"`
}

type Manager interface {
	CreateTask(ctx context.Context, sessionID, userID string) (string, string, error)

	GetTask(ctx context.Context, taskID string) (*TaskState, error)

	UpdateTask(ctx context.Context, task *TaskState) error

	CreateRequest(ctx context.Context, taskID string) (string, error)

	GetRequest(ctx context.Context, requestID string) (*RequestState, error)

	UpdateRequest(ctx context.Context, request *RequestState) error
}
