package cli_frmwk

import (
	"fmt"
	"github.com/c-bata/go-prompt"
	"os"
	"strings"
)

func NewHandler(prompt string, commands ...Command) Handler {
	return Handler{prompt: prompt, commands: commands}
}

func (this Handler) completer(d prompt.Document) []prompt.Suggest {
	return prompt.FilterHasPrefix(this.completion, d.GetWordBeforeCursor(), true)
}

func (this Handler) GetInput() string {
	return prompt.Input(this.prompt, this.completer)
}

func (this Handler) Handle(input string) {
	args := strings.Split(input, " ")
	notfound := true

	for _, c := range this.commands {
		if args[0] == c.Name {
			c.Exec(args, c)
			this.completion = append(this.completion, prompt.Suggest{
				Text: strings.Join(args, " "),
			})
			notfound = false
		}
	}

	if notfound {
		fmt.Println(Red + "Command not found: '" + args[0] + "'" + White)
	}
}

func (this *Handler) UseOSArgs() Handler {
	this.use_os_args = true
	return *this
}

func (this *Handler) SetPrompt(prompt string) {
	this.prompt = prompt
}

func (this *Handler) AddCommand(command Command) {
	// check to verify required args come before any non-required
	in_required := true
	for _, arg := range command.Args {
		if !arg.Required && in_required {
			in_required = false
		} else if arg.Required && !in_required {
			panic("Command \"" + command.Name + "\" has required argument after non-required arguments!\n\tArgument: " + arg.Name.Full)
			os.Exit(1)
		}
	}
	this.commands = append(this.commands, command)
	this.completion = append(this.completion, prompt.Suggest{
		Text:        command.Name,
		Description: command.Description,
	})
}
