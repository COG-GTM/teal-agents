# Teal Agents Go Migration

## Overview
This document outlines the approach for migrating the Teal Agents framework from Python to Go (Golang). The migration is being done incrementally across multiple sessions, with each session focusing on a specific component.

## Migration Strategy

### Session Breakdown
1. **Session 1 (Current): Project Foundation** ✅
   - Go module setup with `go.mod`
   - Core directory structure
   - Basic interfaces for three agent types
   - Initial CI/CD workflows
   - This documentation

2. **Session 2: Core Types and Configuration** ✅
   - Complete implementation of types package
   - YAML configuration loading
   - Environment variable handling
   - Telemetry integration (OpenTelemetry)

3. **Session 3: Core Semantic Kernel Port - Basic Interfaces** ✅
   - Core kernel interfaces (ChatHistory, ChatMessageContent, ContentItems)
   - Function calling types (FunctionCallContent, FunctionResultContent)
   - Agent builder pattern implementation
   - Model provider abstraction (OpenAI, Anthropic, Google)
   - Plugin system interfaces
   - Error handling types
   - All foundational interfaces for agent implementations

4. **Session 4: Sequential Agent Implementation**
   - Complete Sequential agent with task orchestration
   - Task execution with chat history
   - Token usage tracking
   - Streaming support

5. **Session 5: Chat Agent Implementation**
   - Complete Chat agent with conversation handling
   - User context augmentation
   - Memory and state management

6. **Session 6: TealAgent (HITL) Implementation**
   - Human-in-the-loop tool approval
   - Authentication and authorization
   - Task state persistence
   - Resume capability

7. **Session 7: State Management**
   - In-memory state manager
   - Redis state manager
   - Task and request lifecycle

8. **Session 8: API Layer**
   - REST endpoints
   - Server-Sent Events (SSE) streaming
   - WebSocket support
   - A2A protocol integration

9. **Session 9: Testing and Examples**
   - Unit tests for all components
   - Integration tests
   - Example configurations and usage

10. **Session 10: Documentation and Deployment**
    - Complete API documentation
    - Deployment guides
    - Migration guide for users

## Architecture

### Directory Structure
```
teal-agents/
├── go.mod                          # Go module definition
├── go.sum                          # Dependency checksums
├── pkg/                            # Public libraries
│   ├── agents/                     # Agent implementations
│   │   ├── sequential/             # Sequential agent
│   │   ├── chat/                   # Chat agent
│   │   └── teal/                   # TealAgent (HITL)
│   ├── kernel/                     # LLM integration layer
│   ├── config/                     # Configuration management
│   ├── api/                        # Web API layer
│   ├── state/                      # State management
│   ├── telemetry/                  # Observability
│   └── types/                      # Core types and interfaces
├── cmd/                            # Command-line applications
│   └── teal-agents/                # Main server binary
├── internal/                       # Private application code
│   └── testutil/                   # Testing utilities
├── examples/                       # Example configurations
└── .github/workflows/              # CI/CD pipelines
    ├── go-check.yaml               # Linting and testing
    └── go-build.yaml               # Docker build and push
```

### Agent Types

#### 1. Sequential Agent
Executes predefined tasks in order, with each task's output feeding into the next.

**Python Reference**: `src/sk-agents/src/sk_agents/skagents/v1/sequential/sequential_skagents.py`

**Go Interface**: `pkg/agents/sequential/sequential.go`

**Key Features**:
- Task orchestration with ordering
- Intermediate result streaming
- Token usage aggregation across tasks
- Support for multimodal inputs
- Output transformation

#### 2. Chat Agent
Handles conversational interactions with memory and context management.

**Python Reference**: `src/sk-agents/src/sk_agents/skagents/v1/chat/chat_agents.py`

**Go Interface**: `pkg/agents/chat/chat.go`

**Key Features**:
- Single agent conversations
- Chat history maintenance
- User context augmentation
- Streaming responses
- Token tracking

#### 3. TealAgent (HITL)
Provides human-in-the-loop capabilities with authentication and state management.

**Python Reference**: `src/sk-agents/src/sk_agents/tealagents/v1alpha1/agent/handler.py`

**Go Interface**: `pkg/agents/teal/teal.go`

**Key Features**:
- Human approval for high-risk tool calls
- User authentication and authorization
- Task state persistence (Redis or in-memory)
- Resume capability after approval/rejection
- Multimodal message support

## Technology Choices

### Core Dependencies

#### LLM Framework: langchaingo
- **Repository**: https://github.com/tmc/langchaingo
- **Reason**: Mature Go LLM framework (7.7k+ stars) with support for multiple providers
- **Replaces**: Python's semantic-kernel
- **Features**:
  - OpenAI, Anthropic, Google integration
  - Agent and chain abstractions
  - Tool/function calling
  - Streaming support

#### Web Framework: TBD (Session 8)
Options being considered:
- **Gin**: Fast, minimal overhead, popular (74k+ stars)
- **Fiber**: Express-like API, very fast (31k+ stars)
- **Echo**: High performance, minimal (28k+ stars)
- **Standard net/http**: No dependencies, full control

Decision will be made in Session 8 based on:
- Performance requirements
- Middleware needs
- SSE/WebSocket support
- Team preferences

#### Redis Client: go-redis
- **Repository**: https://github.com/redis/go-redis
- **Reason**: Official Redis Go client, widely used
- **Replaces**: Python's redis package
- **Features**:
  - Redis Streams support
  - Connection pooling
  - Cluster support

#### Configuration: gopkg.in/yaml.v3
- **Reason**: Standard Go YAML library
- **Replaces**: Python's pyyaml + pydantic-yaml

#### Testing: Go standard library + testify
- **Standard testing package**: Built-in test framework
- **testify**: Assertions and mocking (github.com/stretchr/testify)
- **Replaces**: Python's pytest

### Go Idioms and Patterns

#### Error Handling
Go uses explicit error returns instead of exceptions:
```go
result, err := agent.Invoke(ctx, inputs)
if err != nil {
    return nil, fmt.Errorf("agent invocation failed: %w", err)
}
```

#### Context for Cancellation
All long-running operations accept `context.Context`:
```go
func (a *Agent) Invoke(ctx context.Context, inputs map[string]interface{}) (*types.InvokeResponse, error)
```

#### Channels for Streaming
Streaming uses channels instead of async iterables:
```go
responseChan, err := agent.InvokeStream(ctx, inputs)
for response := range responseChan {
    if response.IsPartial {
        // Handle partial response
    }
}
```

#### Interfaces in Consumer Packages
Go convention: define interfaces where they're used, not where they're implemented.

## Configuration

### Environment Variables
The Go implementation will maintain compatibility with existing Python environment variables:
- `TA_SERVICE_CONFIG`: Path to service configuration YAML
- `TA_STATE_MANAGEMENT`: State backend (in-memory/redis)
- `TA_REDIS_*`: Redis connection parameters
- `TA_GITHUB`: Enable dynamic code loading

### YAML Configuration Format
The Go implementation will support the same YAML configuration format as Python for backward compatibility.

## API Compatibility

### REST Endpoints
The Go API will maintain the same endpoint structure:
- `POST /{agent}/{version}` - Synchronous invocation
- `POST /{agent}/{version}/sse` - Server-Sent Events streaming
- `POST /{agent}/{version}/stream` - WebSocket streaming

### Response Format
JSON response format will match Python implementation for API compatibility.

## Testing Strategy

### Unit Tests
- Test coverage target: >80%
- Use table-driven tests (Go idiom)
- Mock external dependencies (LLM providers, Redis)

### Integration Tests
- Test full agent workflows
- Test state persistence
- Test HITL approval flows

### Performance Tests
- Benchmark critical paths
- Compare performance with Python implementation
- Identify optimization opportunities

## Deployment

### Docker
A new Dockerfile (`Dockerfile.go`) will be created for Go builds:
- Multi-stage build for minimal image size
- Static binary compilation
- No Python runtime overhead

### Kubernetes
Existing Kubernetes configurations will be updated to support Go deployments alongside Python.

## Migration Timeline

Estimated timeline (assuming 1 session per week):
- **Weeks 1-2**: Foundation and core types (Sessions 1-2)
- **Weeks 3-4**: LLM integration (Session 3)
- **Weeks 5-7**: Agent implementations (Sessions 4-6)
- **Weeks 8-9**: State and API (Sessions 7-8)
- **Week 10**: Testing and examples (Session 9)
- **Week 11**: Documentation and deployment (Session 10)

Total: ~11 weeks for complete migration

## Benefits of Go Migration

1. **Performance**: 
   - Compiled binary (no interpreter overhead)
   - Better concurrency primitives (goroutines vs threads)
   - Lower memory footprint

2. **Deployment**:
   - Single static binary
   - No dependency management issues
   - Smaller container images

3. **Type Safety**:
   - Compile-time type checking
   - No runtime type errors
   - Better IDE support

4. **Concurrency**:
   - Native goroutines for parallel execution
   - Channels for communication
   - Better resource utilization

5. **Maintenance**:
   - Simpler dependency management
   - Faster CI/CD builds
   - Easier debugging

## Challenges and Mitigation

### Challenge 1: LLM Library Maturity
- **Issue**: langchaingo less mature than Python semantic-kernel
- **Mitigation**: Abstract LLM interaction behind interfaces, easy to swap implementations

### Challenge 2: Team Go Experience
- **Issue**: Team may be more familiar with Python
- **Mitigation**: Comprehensive documentation, code reviews, gradual rollout

### Challenge 3: Testing Migration
- **Issue**: Need to verify equivalent behavior
- **Mitigation**: Parallel testing with Python implementation during migration

## References

### Python Codebase
- Main agent framework: `src/sk-agents`
- Core types: `src/sk-agents/src/sk_agents/ska_types.py`
- Sequential agents: `src/sk-agents/src/sk_agents/skagents/v1/sequential/sequential_skagents.py`
- Chat agents: `src/sk-agents/src/sk_agents/skagents/v1/chat/chat_agents.py`
- TealAgents: `src/sk-agents/src/sk_agents/tealagents/v1alpha1/agent/handler.py`

### Go Resources
- langchaingo: https://github.com/tmc/langchaingo
- Go Redis: https://github.com/redis/go-redis
- Go best practices: https://go.dev/doc/effective_go

## Contributing

For this migration project:
1. Each session focuses on a specific component
2. All changes go through PR review
3. Tests must pass before merging
4. Documentation must be updated with code changes

## Questions and Support

For questions about the migration:
- Review this document first
- Check the Python implementation references
- Consult the Go interfaces in `pkg/` directories
- Ask in team channels

---

Last Updated: Session 3 - Core Semantic Kernel Port - Basic Interfaces
Next Session: Session 4 - Sequential Agent Implementation
