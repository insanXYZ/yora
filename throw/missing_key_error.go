package throw

import (
	"errors"
)

func MissingKey() error {
	return errors.New("GEMINI_API_KEY environment variable not set\n\nYou can set this with:\n# Linux\n	export GEMINI_API_KEY = yourapikey\n# Windows (btw, i dont know what operating system is this)\n	$Env:GEMINI_API_KEY = yourapikey")
}
