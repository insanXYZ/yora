package throw

import "errors"

func ClientGeminiKey() error {
	return errors.New("GEMINI_API_KEY has invalid or broke key")
}
