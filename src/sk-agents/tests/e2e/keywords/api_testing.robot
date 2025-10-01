*** Settings ***
Documentation     Keywords for HTTP request/response testing
Library           RequestsLibrary

*** Keywords ***
Send Agent Request
    [Documentation]    Send a POST request to the agent API
    [Arguments]    ${endpoint}    ${payload}

Verify Agent Response
    [Documentation]    Verify the agent response structure and content
    [Arguments]    ${response}
