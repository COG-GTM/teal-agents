# University Agent System E2E Tests

Robot Framework end-to-end tests for the University Agent System.

## Discovered System Components

### 1. UniversityPlugin
**Location**: `src/orchestrators/assistant-orchestrator/example/university/custom_plugins.py`

The core plugin providing university search functionality:
- `search_universities(query: str)` - Search universities by name
- `get_universities_by_country(country: str)` - Find universities in a specific country

**External API**: universities.hipolabs.com
- No authentication required
- Returns JSON array of university objects
- Free tier, no rate limiting mentioned

### 2. Streamlit UI
**Location**: `src/orchestrators/assistant-orchestrator/example/university/streamlit_ui.py`

Interactive chat interface:
- Default URL: http://localhost:8001 (configurable)
- Connects to `/UniversityAgent/0.1` endpoint
- Chat-based interaction with conversation history
- Formats university data for user-friendly display

**Key Functions**:
- `call_university_agent()` - POST requests to agent API
- `format_university_data()` - Display formatting
- `check_agent_status()` - Health check

### 3. Agent Configuration
**Location**: `src/orchestrators/assistant-orchestrator/example/university/config.yaml`

Sequential agent configuration:
- **API Version**: skagents/v1
- **Service Name**: UniversityAgent
- **Version**: 0.1
- **Model**: gemini-2.0-flash-lite
- **Plugin**: UniversityPlugin

### 4. FastAPI Service
**Location**: `src/sk-agents/src/sk_agents/app.py`, `appv1.py`

Dynamic endpoint creation based on configuration:
- **Endpoint Pattern**: `/{name}/{version}` (e.g., `/UniversityAgent/0.1`)
- **Docs**: `/{name}/{version}/docs`
- Supports both REST and WebSocket routes
- Uses custom chat completion factory

### 5. Gemini Integration
**Location**: `src/sk-agents/src/sk_agents/chat_completion/custom/gemini_chat_completion_factory.py`

Custom completion factory for Google Gemini:
- **Supported Models**: gemini-1.5-flash, gemini-1.5-pro, gemini-1.0-pro, gemini-2.0-flash-lite
- **Environment Variable**: GEMINI_API_KEY
- Integrates with Semantic Kernel via GoogleAIChatCompletion

## Architecture

```
User → Streamlit UI (Port 8502)
         ↓ POST /UniversityAgent/0.1
      FastAPI Agent Service (Port 8001)
         ↓ LLM calls
      Gemini API (gemini-2.0-flash-lite)
         ↓ Plugin invocation
      UniversityPlugin
         ↓ HTTP GET
      universities.hipolabs.com API
```

## Test Modes

### Local Mode
Tests start and manage the FastAPI and Streamlit processes:
- Starts agent service on port 8001
- Starts Streamlit UI on port 8502
- Can mock external APIs
- Full process lifecycle control

### UAT Mode
Tests run against already-running instances:
- Assumes services are already deployed
- No process management
- Uses real external APIs (typically)
- For testing deployed environments

## Configuration

Test configuration is in `resources/environment_configs.yaml`:
- Service URLs and ports
- Startup commands (local mode)
- Environment variables
- Timeouts and settings
- Mock server configuration

## Test Data

Test scenarios and expected responses are in:
- `resources/test_data.yaml` - Test cases and expected results
- `resources/mock_responses.json` - Canned API responses for mocking

## Phase 1 Status

✅ Directory structure created
✅ Placeholder test files with basic Robot Framework structure
✅ Keyword resource files for different testing aspects
✅ Configuration files for local and UAT modes
✅ Python library files with class definitions and docstrings
✅ Dependencies added to pyproject.toml
✅ System components documented

## Next Steps (Future Phases)

- **Phase 2**: Implement agent lifecycle keywords and basic API tests
- **Phase 3**: Implement Streamlit UI automation
- **Phase 4**: Implement external API mocking
- **Phase 5**: Implement full system integration tests
- **Phase 6**: CI/CD integration and reporting

## Running Tests (Future)

Once implemented, tests will be run with:
```bash
# Local mode
robot --variable MODE:local tests/e2e/university_agent_system.robot

# UAT mode
robot --variable MODE:uat tests/e2e/university_agent_system.robot
```

## Dependencies

Robot Framework libraries:
- robotframework - Core framework
- robotframework-requests - HTTP API testing
- robotframework-process - Process management
- robotframework-seleniumlibrary - Browser automation

Install with:
```bash
cd ~/repos/teal-agents/src/sk-agents
uv sync
```
