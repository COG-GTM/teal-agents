*** Settings ***
Documentation     Keywords for mock setup/teardown of external APIs
Library           ../libraries/MockServer.py

*** Keywords ***
Start Mock Universities API
    [Documentation]    Start mock server for universities.hipolabs.com

Stop Mock Universities API
    [Documentation]    Stop the mock universities API server

Configure Mock Response
    [Documentation]    Set up a specific mock response
    [Arguments]    ${endpoint}    ${response_data}
