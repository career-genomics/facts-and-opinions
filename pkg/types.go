package pkg

type Knowledge struct {
	Content string
}

type OllamaClassifier struct {
	BaseUrl string
	Model   string
}

type OllamaRequest struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
	Stream bool   `json:"stream"`
}

type OllamaResponse struct {
	Response string `json:"response"`
	Done     bool   `json:"done"`
}
