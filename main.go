package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
	"github.com/briandowns/spinner"
	"github.com/fatih/color"
)

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type RequestBody struct {
	Messages    []Message `json:"messages"`
	Model      string `json:"model"`
	Temperature float64 `json:"temperature"`
	N          int    `json:"n"`
}

func main() {
	// Parse command line arguments
	filePaths := flag.String("files", "", "Comma-separated list of file paths to use as input")
	flag.Parse()

	if *filePaths == "" {
		fmt.Println("Please provide a list of file paths using the --files flag")
		return
	}

	// Read files into memory
	inputs, err := readFiles(strings.Split(*filePaths, ","))
	if err != nil {
		fmt.Printf("Error reading input files: %v\n", err)
		return
	}

	openai_api_key := os.Getenv("OPENAI_API_KEY")
	url := "https://api.openai.com/v1/chat/completions"

	// Prompt user for first message and insert it into RequestBody
	// firstMessage := getUserInput()
	requestBody := RequestBody{

		Model: "gpt-3.5-turbo",
		Messages: []Message{
			{
				Role:    "user",
				Content: strings.Join(inputs, ""),
			},
		},
		Temperature: 0.7,
		N:           1,
	};

	// Make initial HTTP request with requestBody
	response, err := makeOpenAIRequest(requestBody, url, openai_api_key)
	if err != nil {
		log.Fatal("Error making OpenAI request: ", err)
	}

	// Loop to continuously prompt user for input and make requests to OpenAI API
	for {

		// Extract response message and print to console
		message := extractResponseMessage(response)
		color.New(color.FgYellow).Printf("AI: " + message + "\n")

		// Prompt user for next message and insert it into RequestBody
		userInput := getUserInput()
		requestBody.Messages = append(requestBody.Messages, Message{
			Role:    "user",
			Content: userInput,
		})

		// Make HTTP request with updated RequestBody
		response, err = makeOpenAIRequest(requestBody, url, openai_api_key)
		if err != nil {
			log.Fatal("Error making OpenAI request: ", err)
		}
	}
}

func getUserInput() string {
	reader := bufio.NewReader(os.Stdin)
	color.New(color.FgGreen).Printf("You: ")
	userInput, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("Error reading user input: ", err)
	}
	// Remove newline character from the end of the input
	userInput = strings.TrimSuffix(userInput, "\n")
	return userInput
}

func readFiles(filePaths []string) ([]string, error) {
	var inputs []string

	for _, filePath := range filePaths {
		file, err := os.Open(filePath)
		if err != nil {
			return nil, fmt.Errorf("error opening file %s: %v", filePath, err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			inputs = append(inputs, scanner.Text())
		}

		if err := scanner.Err(); err != nil {
			return nil, fmt.Errorf("error reading file %s: %v", filePath, err)
		}
	}

	return inputs, nil
}
func makeOpenAIRequest(requestBody RequestBody, url string, openai_api_key string) (map[string]interface{}, error) {
	spinner := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
	spinner.Start()
	defer spinner.Stop()

	// Construct HTTP request with RequestBody
	requestBodyBytes, err := json.Marshal(requestBody)
	if err != nil {
		return nil, fmt.Errorf("error marshaling request body: %v", err)
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBodyBytes))
	if err != nil {
		return nil, fmt.Errorf("error creating HTTP request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", openai_api_key))

	// Make HTTP request and parse response body
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making HTTP request: %v", err)
	}
	defer resp.Body.Close()

	var data map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, fmt.Errorf("error decoding response body: %v", err)
	}

	return data, nil
}

func extractResponseMessage(response map[string]interface{}) string {
	choices := response["choices"].([]interface{})
	message := choices[0].(map[string]interface{})["message"].(map[string]interface{})["content"].(string)
	return message
}
