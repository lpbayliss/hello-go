package greeting

import (
	"strings"
	"testing"
)

func TestValidateName(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr error
	}{
		{"valid simple name", "John", nil},
		{"valid name with space", "John Doe", nil},
		{"valid name with hyphen", "Mary-Jane", nil},
		{"valid name with apostrophe", "O'Brien", nil},
		{"empty string", "", ErrNameEmpty},
		{"only spaces", "   ", ErrNameEmpty},
		{"too long", strings.Repeat("a", 101), ErrNameTooLong},
		{"exactly 100 chars", strings.Repeat("a", 100), nil},
		{"contains numbers", "John123", ErrNameInvalid},
		{"contains special chars", "John@Doe", ErrNameInvalid},
		{"contains underscore", "John_Doe", ErrNameInvalid},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateName(tt.input)
			if err != tt.wantErr {
				t.Errorf("ValidateName(%q) = %v, want %v", tt.input, err, tt.wantErr)
			}
		})
	}
}

func TestGenerateGreeting(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{"empty name", "", "Hello, World!"},
		{"simple name", "John", "Hello, John!"},
		{"name with spaces", "  John  ", "Hello, John!"},
		{"full name", "John Doe", "Hello, John Doe!"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GenerateGreeting(tt.input)
			if got != tt.want {
				t.Errorf("GenerateGreeting(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestFormatGreeting(t *testing.T) {
	tests := []struct {
		name      string
		greeting  string
		uppercase bool
		want      string
	}{
		{"no formatting", "Hello, World!", false, "Hello, World!"},
		{"uppercase", "Hello, World!", true, "HELLO, WORLD!"},
		{"empty string no format", "", false, ""},
		{"empty string uppercase", "", true, ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FormatGreeting(tt.greeting, tt.uppercase)
			if got != tt.want {
				t.Errorf("FormatGreeting(%q, %v) = %q, want %q", tt.greeting, tt.uppercase, got, tt.want)
			}
		})
	}
}
