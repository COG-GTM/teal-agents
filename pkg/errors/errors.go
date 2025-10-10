package errors

import "fmt"

type AgentError struct {
	Message string
	Cause   error
}

func (e *AgentError) Error() string {
	if e.Cause != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Cause)
	}
	return e.Message
}

func (e *AgentError) Unwrap() error {
	return e.Cause
}

type AuthenticationError struct {
	AgentError
}

func NewAuthenticationError(message string, cause error) *AuthenticationError {
	return &AuthenticationError{
		AgentError: AgentError{Message: message, Cause: cause},
	}
}

type AgentInvokeError struct {
	AgentError
}

func NewAgentInvokeError(message string, cause error) *AgentInvokeError {
	return &AgentInvokeError{
		AgentError: AgentError{Message: message, Cause: cause},
	}
}

type PersistenceError struct {
	AgentError
}

func NewPersistenceError(message string, cause error) *PersistenceError {
	return &PersistenceError{
		AgentError: AgentError{Message: message, Cause: cause},
	}
}

type ConfigurationError struct {
	AgentError
}

func NewConfigurationError(message string, cause error) *ConfigurationError {
	return &ConfigurationError{
		AgentError: AgentError{Message: message, Cause: cause},
	}
}

type HITLInterventionRequired struct {
	AgentError
	FunctionCalls []interface{}
}

func NewHITLInterventionRequired(message string, functionCalls []interface{}) *HITLInterventionRequired {
	return &HITLInterventionRequired{
		AgentError:    AgentError{Message: message},
		FunctionCalls: functionCalls,
	}
}
