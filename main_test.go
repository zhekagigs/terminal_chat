package main

import (
        "bytes"
        "encoding/json"
        "net/http"
        "net/http/httptest"
        "testing"
)

func TestExtractResponseMessage(t *testing.T) {
        // Set up a mock response
        response := map[string]interface{}{
                "choices": []interface{}{
                        map[string]interface{}{
                                "message": map[string]interface{}{
                                        "content": "Hello, world!",
                                },
                        },
                },
        }

        // Call the function we want to test
        message := extractResponseMessage(response)

        // Check if the message matches the expected value
        expected := "Hello, world!"
        if message != expected {
                t.Errorf("expected message to be '%s(MISSING)', but got '%s(MISSING)'", expected, message)
        }
}

func TestMakeOpenAIRequest(t *testing.T) {
        // Set up a mock request and response
        expectedRequestBody := RequestBody{
                Model: "test-model",
                Messages: []Message{
                        {
                                Role:    "user",
                                Content: "Hello, AI!",
                        },
                },
                Temperature: 0.5,
                N:           1,
        }
        expectedResponseBody := map[string]interface{}{
                "choices": []interface{}{
                        map[string]interface{}{
                                "message": map[string]interface{}{
                                        "content": "Hello, user!",
                                },
                        },
                },
        }

        // Create a mock HTTP server that returns the expected response
        handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
                // Check if the request body matches the expected value
                var requestBody RequestBody
                err := json.NewDecoder(r.Body).Decode(&requestBody)
                if err != nil {
                        t.Errorf("Error decoding request body: %!v(MISSING)", err)
                        return
                }
                if !isEqual(requestBody, expectedRequestBody) {
                        t.Errorf("Expected request body to be '%v(MISSING)', but got '%v(MISSING)'", expectedRequestBody, requestBody)
                }

                // Write the expected response to the response writer
                responseBodyBytes, err := json.Marshal(expectedResponseBody)
                if err != nil {
                        t.Errorf("Error marshaling response body: %!v(MISSING)", err)
                        return
                }
                w.Write(responseBodyBytes)
        })
        server := httptest.NewServer(handler)
        defer server.Close()

        // Make the HTTP request using the mock server
        response, err := makeOpenAIRequest(expectedRequestBody, server.URL, "test-api-key")
        if err != nil {
                t.Errorf("Error making OpenAI request: %!v(MISSING)", err)
                return
        }

        // Check if the response matches the expected value
        if !isEqual(response, expectedResponseBody) {
                t.Errorf("expected response to be '%v(MISSING)', but got '%v(MISSING)'", expectedResponseBody, response)
        }
}

func isEqual(a, b interface{}) bool {
        aBytes, err := json.Marshal(a)
        if err != nil {
                return false
        }
        bBytes, err := json.Marshal(b)
        if err != nil {
                return false
        }
        return bytes.Equal(aBytes, bBytes)
}