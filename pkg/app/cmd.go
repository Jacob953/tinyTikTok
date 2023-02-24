package app

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// NewCommand creates a new sub command instance based on the given command name and other options.
func NewCommand(usage string, desc string, opts ...CommandOption) *Command {
	cmd := &Command{
		usage: usage,
		desc:  desc,
	}

	for _, o := range opts {
		o(cmd)
	}

	return cmd
}

// Command is a sub command structure of a cli application.
type Command struct {
	usage    string
	desc     string
	options  CliOptions
	commands []*Command
	runFunc  RunCommandFunc
}

// AddCommand adds sub command to the current command.
func (cmd *Command) AddCommand(c *Command) {
	cmd.commands = append(cmd.commands, c)
}

// AddCommands adds multiple sub commands to the current command.
func (cmd *Command) AddCommands(cmds ...*Command) {
	cmd.commands = append(cmd.commands, cmds...)
}

func (cmd Command) cobraCommand() *cobra.Command {
	cobraCMD := &cobra.Command{
		Use:   cmd.usage,
		Short: cmd.desc,
	}
	cobraCMD.SetOut(os.Stdout)
	cobraCMD.Flags().SortFlags = false

	if len(cmd.commands) > 0 {
		for _, command := range cmd.commands {
			cobraCMD.AddCommand(command.cobraCommand())
		}
	}

	if cmd.runFunc != nil {
		cobraCMD.Run = cmd.runCommand
	}

	if cmd.options != nil {
		for _, f := range cmd.options.Flags().FlagSets {
			cobraCMD.Flags().AddFlagSet(f)
		}
	}

	addHelpCommandFlag(cmd.usage, cobraCMD.Flags())

	return cobraCMD
}

func (cmd Command) runCommand(_ *cobra.Command, args []string) {
	if cmd.runFunc != nil {
		if err := cmd.runFunc(args); err != nil {
			fmt.Printf("%v %v\n", color.RedString("Error:"), err)
			os.Exit(1)
		}
	}
}

// CommandOption defines optional parameters for initializing the command structure.
type CommandOption func(*Command)

// WithCommandOptions to open the application's function to read from the command line.
func WithCommandOptions(opt CliOptions) CommandOption {
	return func(cmd *Command) {
		cmd.options = opt
	}
}

// RunCommandFunc defines the application's command startup callback function.
type RunCommandFunc func(args []string) error

// WithCommandRunFunc is used to set the application's command startup callback function option.
func WithCommandRunFunc(run RunCommandFunc) CommandOption {
	return func(cmd *Command) {
		cmd.runFunc = run
	}
}
