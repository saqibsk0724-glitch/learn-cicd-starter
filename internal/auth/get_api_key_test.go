package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name       string
		authHeader string
		want       string
		wantErr    bool
	}{
		{
			name:       "valid API key",
			authHeader: "ApiKey abc123",
			want:       "abc123",
			wantErr:    false,
		},
		{
			name:       "missing authorization header",
			authHeader: "",
			want:       "",
			wantErr:    true,
		},
		{
			name:       "invalid authorization scheme",
			authHeader: "Bearer abc123",
			want:       "",
			wantErr:    true,
		},
		{
			name:       "malformed authorization header",
			authHeader: "ApiKey",
			want:       "",
			wantErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			headers := http.Header{}

			if tt.authHeader != "" {
				headers.Set("Authorization", tt.authHeader)
			}

			got, err := GetAPIKey(headers)

			if (err != nil) != tt.wantErr {
				t.Fatalf("GetAPIKey() error = %v, wantErr %v", err, tt.wantErr)
			}

			if got != tt.want {
				t.Errorf("GetAPIKey() = %q, want %q", got, tt.want)
			}
		})
	}
}
