package throw

import (
	"errors"
)

func MissingKey() error {
	return errors.New("GEMINI_API_KEY environment variable not set\n\n" +
		"You can get this from https://aistudio.google.com/app/apikey and set like this:\n\n" +
		"# Linux\n	export GEMINI_API_KEY=\"yourapikey\"\n" +
		"# Windows \n	set GEMINI_API_KEY=\"yourapikey\"")
}
