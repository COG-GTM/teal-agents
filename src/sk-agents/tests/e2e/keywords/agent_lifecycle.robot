*** Settings ***
Documentation     Keywords for managing agent lifecycle (start/stop/health checks)
Library           ../libraries/ProcessManager.py

*** Keywords ***
Start University Agent
    [Documentation]    Start the FastAPI agent service

Stop University Agent
    [Documentation]    Stop the FastAPI agent service

Check Agent Health
    [Documentation]    Verify agent is responding
