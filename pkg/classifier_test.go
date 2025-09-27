package pkg

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func setupMockServer() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/api/generate" {
			fmt.Fprintln(w, `{"response": "TRUE", "done": true}`)
		}
	}))
}

func TestNewOllamaClassifier(t *testing.T) {
	type args struct {
		model string
	}
	tests := []struct {
		name string
		args args
		want *OllamaClassifier
	}{
		{
			name: "Test with model 'llama2'",
			args: args{model: "llama2"},
			want: &OllamaClassifier{
				BaseUrl: "http://localhost:11434/api/generate",
				Model:   "llama2",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewOllamaClassifier(tt.args.model); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewOllamaClassifier() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOllamaClassifier_IsFact(t *testing.T) {
	type fields struct {
		BaseUrl string
		Model   string
	}
	type args struct {
		k Knowledge
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "Test with a factual statement",
			fields: fields{
				BaseUrl: setupMockServer().URL + "/api/generate",
				Model:   "llama2",
			},
			args: args{k: Knowledge{Content: "The Earth orbits the Sun."}},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ai := &OllamaClassifier{
				BaseUrl: tt.fields.BaseUrl,
				Model:   tt.fields.Model,
			}
			if got := ai.IsFact(tt.args.k); got != tt.want {
				t.Errorf("IsFact() = %v, want %v", got, tt.want)
			}
		})
	}
	defer setupMockServer().Close()
}

func TestOllamaClassifier_IsOpinion(t *testing.T) {
	type fields struct {
		BaseUrl string
		Model   string
	}
	type args struct {
		k Knowledge
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "Test with an opinion statement",
			fields: fields{
				BaseUrl: setupMockServer().URL + "/api/generate",
				Model:   "llama2",
			},
			args: args{k: Knowledge{Content: "Chocolate ice cream is the best flavor."}},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ai := &OllamaClassifier{
				BaseUrl: tt.fields.BaseUrl,
				Model:   tt.fields.Model,
			}
			if got := ai.IsOpinion(tt.args.k); got != tt.want {
				t.Errorf("IsOpinion() = %v, want %v", got, tt.want)
			}
		})
	}
	defer setupMockServer().Close()
}

func TestValidUpdate(t *testing.T) {
	type args struct {
		x  Knowledge
		y  Knowledge
		ai *OllamaClassifier
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test with valid fact and opinion",
			args: args{
				x:  Knowledge{Content: "The Earth orbits the Sun."},
				y:  Knowledge{Content: "Chocolate ice cream is the best flavor."},
				ai: &OllamaClassifier{BaseUrl: setupMockServer().URL + "/api/generate", Model: "llama2"},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidUpdate(tt.args.x, tt.args.y, tt.args.ai); got != tt.want {
				t.Errorf("ValidUpdate() = %v, want %v", got, tt.want)
			}
		})
	}
}
