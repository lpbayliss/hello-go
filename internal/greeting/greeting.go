package greeting

import (
	"errors"
	"fmt"
	"strings"
	"unicode"
)

var (
	ErrNameEmpty   = errors.New("name cannot be empty")
	ErrNameTooLong = errors.New("name cannot exceed 100 characters")
	ErrNameInvalid = errors.New("name contains invalid characters")
)

// ValidateName checks if a name is valid for greeting
// Returns error if name is empty, too long, or contains invalid chars
func ValidateName(name string) error {
	if strings.TrimSpace(name) == "" {
		return ErrNameEmpty
	}

	if len(name) > 100 {
		return ErrNameTooLong
	}

	for _, r := range name {
		if !unicode.IsLetter(r) && !unicode.IsSpace(r) && r != '-' && r != '\'' {
			return ErrNameInvalid
		}
	}

	return nil
}

// GenerateGreeting creates a personalized greeting message
// Returns default greeting if name is empty
func GenerateGreeting(name string) string {
	name = strings.TrimSpace(name)
	if name == "" {
		return "Hello, World!"
	}
	return fmt.Sprintf("Hello, %s!", name)
}

// FormatGreeting applies formatting to a greeting string
func FormatGreeting(greeting string, uppercase bool) string {
	if uppercase {
		return strings.ToUpper(greeting)
	}
	return greeting
}
