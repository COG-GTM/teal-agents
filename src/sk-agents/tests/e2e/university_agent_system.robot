*** Settings ***
Documentation     Full system integration tests for University Agent System
...               Tests the complete flow: Streamlit UI -> FastAPI Agent -> External APIs
...               TODO: Implement in Phase 2
Resource          keywords/agent_lifecycle.robot
Resource          keywords/api_testing.robot
Resource          keywords/external_mocking.robot

*** Variables ***

*** Test Cases ***
