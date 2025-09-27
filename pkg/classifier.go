package pkg

import (
	"fmt"
	"strings"
)

func NewOllamaClassifier(model string) *OllamaClassifier {
	return &OllamaClassifier{
		BaseUrl: "http://localhost:11434/api/generate",
		Model:   model,
	}
}

func (ai *OllamaClassifier) IsFact(k Knowledge) bool {
	prompt := fmt.Sprintf(`You are a fact-checker. Analyze this statement and respond with Only "True" or "False":
	Statement: "%s"
	True if :
	- The statement is objectively verifiable and can be proven with evidence.
	- Measurable and observable data supports the statement.
    - Scientific consensus or historical records confirm the statement.
    - Mathematical or logical truths validate the statement.

	False if :
	- The statement is subjective and based on personal opinions, interpretations, or beliefs.
	- Personal belief or feeling
    - Unclear or ambiguous 
    - Can not be proven

	Response:`, k.Content)

	resp, _ := ai.call(prompt)
	return strings.Contains(strings.ToUpper(resp), "TRUE")
}

func (ai *OllamaClassifier) IsOpinion(k Knowledge) bool {
	prompt := fmt.Sprintf(`You are analyzing subjective vs objective language. Respond with ONLY "TRUE" or "FALSE":

Statement: "%s"

TRUE if:
- Personal preference or belief
- Uses opinion words (should, better, prefer, think, feel)
- Subjective judgment or taste
- Reasonable people could disagree

FALSE if:
- Objective fact or measurement
- Scientific or mathematical truth
- Observable reality

Response:`, k.Content)

	resp, _ := ai.call(prompt)
	return strings.Contains(strings.ToUpper(resp), "TRUE")
}

func ValidUpdate(x, y Knowledge, ai *OllamaClassifier) bool {
	// ∀x,y: ValidUpdate(x,y) ↔ (Fact(x) ∧ Opinion(y))
	return ai.IsFact(x) && ai.IsOpinion(y)
}
