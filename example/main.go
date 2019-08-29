package main

import (
	"github.com/c-bata/go-prompt"
	cobraprompt "github.com/tomyutao/cobra-prompt"
	"github.com/tomyutao/cobra-prompt/example/cmd"
)

func main() {
	shell := &cobraprompt.CobraPrompt{
		RootCmd:                cmd.RootCmd,
		DynamicSuggestionsFunc: handleDynamicSuggestions,
		ResetFlagsFlag:         true,
		ShowAlias:              true,
		ShowHiddenFlag:         true,
		ShowHiddenCommand:      true,
		GoPromptOptions: []prompt.Option{
			prompt.OptionTitle("cobra-prompt-example"),
			prompt.OptionPrefix(">(^'^)> "),
			prompt.OptionMaxSuggestion(10),
		},
	}
	shell.Run()
}

func handleDynamicSuggestions(annotation string, _ prompt.Document) []prompt.Suggest {
	switch annotation {
	case "GetFood":
		return GetFood()
	default:
		return []prompt.Suggest{}
	}
}

func GetFood() []prompt.Suggest {
	return []prompt.Suggest{
		{Text: "apple", Description: "Green apple"},
		{Text: "tomato", Description: "Red tomato"},
	}
}
