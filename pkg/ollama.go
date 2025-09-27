package pkg

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func (ai *OllamaClassifier) call(prompt string) (string, error) {

	var oresp OllamaResponse

	reqBody := OllamaRequest{
		Model:  ai.Model,
		Prompt: prompt,
		Stream: false,
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		fmt.Printf("Error marshalling request body: %v\n", err)
		return "FALSE", err
	}

	resp, err := http.Post(ai.BaseUrl, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Printf("Error making POST request: %v\n", err)
		return "FALSE", err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %v\n", err)
		return "FALSE", err
	}

	err = json.Unmarshal(body, &oresp)
	if err != nil {
		fmt.Printf("Error unmarshalling response body: %v\n", err)
		return "FALSE", err
	}

	return strings.TrimSpace(oresp.Response), nil

}
