*** Settings ***
Documentation     Keywords for Streamlit UI automation
Library           SeleniumLibrary
Library           ../libraries/StreamlitHelper.py

*** Keywords ***
Open Streamlit UI
    [Documentation]    Open the Streamlit application in browser

Send Chat Message
    [Documentation]    Send a message in the Streamlit chat interface
    [Arguments]    ${message}

Verify University Results Displayed
    [Documentation]    Check that university search results are shown
