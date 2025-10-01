package models

import "fmt"

type ModelNotFoundError struct {
	ModelName string
}

func (e *ModelNotFoundError) Error() string {
	return fmt.Sprintf("model not found: %s", e.ModelName)
}

type NotImplementedError struct {
	Feature string
}

func (e *NotImplementedError) Error() string {
	return fmt.Sprintf("not implemented: %s - will be completed in future session", e.Feature)
}

type ProviderError struct {
	Provider string
	Message  string
	Cause    error
}

func (e *ProviderError) Error() string {
	if e.Cause != nil {
		return fmt.Sprintf("%s provider error: %s: %v", e.Provider, e.Message, e.Cause)
	}
	return fmt.Sprintf("%s provider error: %s", e.Provider, e.Message)
}

func (e *ProviderError) Unwrap() error {
	return e.Cause
}
