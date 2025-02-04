package xvtranslategooglefree_test

import (
	"testing"
	"v1/translate4me/libs/xvtranslategooglefree" // Correct import path
)

func TestTranslate(t *testing.T) {
	tests := []struct {
		name           string
		source         string
		sourceLang     string
		targetLang     string
		expectedResult string
		expectError    bool
	}{
		{
			name:           "Valid Translation",
			source:         "Hello",
			sourceLang:     "en",
			targetLang:     "es",
			expectedResult: "Hola", // Spanish translation of "Hello"
			expectError:    false,
		},
		{
			name:           "Empty Source Text",
			source:         "",
			sourceLang:     "en",
			targetLang:     "es",
			expectedResult: "",
			expectError:    true,
		},
		{
			name:           "Invalid Source Language",
			source:         "Hello",
			sourceLang:     "xx", // Invalid source language
			targetLang:     "es",
			expectedResult: "",
			expectError:    true,
		},
		{
			name:           "Invalid Target Language",
			source:         "Hello",
			sourceLang:     "en",
			targetLang:     "xx", // Invalid target language
			expectedResult: "",
			expectError:    true,
		},
	}

	// Loop through the test cases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := xvtranslategooglefree.Translate(tt.source, tt.sourceLang, tt.targetLang)
			if (err != nil) != tt.expectError {
				t.Errorf("Translate() error = %v, wantErr %v", err, tt.expectError)
				return
			}
			if got != tt.expectedResult {
				t.Errorf("Translate() = %v, want %v", got, tt.expectedResult)
			}
		})
	}
}
