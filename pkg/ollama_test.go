package pkg

import "testing"

func TestOllamaClassifier_call(t *testing.T) {
	type fields struct {
		BaseUrl string
		Model   string
	}
	type args struct {
		prompt string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Test with a simple prompt",
			fields: fields{
				BaseUrl: setupMockServer().URL + "/api/generate",
				Model:   "llama2",
			},
			args:    args{prompt: "Is the sky blue?"},
			want:    "TRUE",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ai := &OllamaClassifier{
				BaseUrl: tt.fields.BaseUrl,
				Model:   tt.fields.Model,
			}
			got, err := ai.call(tt.args.prompt)
			if (err != nil) != tt.wantErr {
				t.Errorf("call() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("call() got = %v, want %v", got, tt.want)
			}
		})
	}
}
