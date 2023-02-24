package app

import (
	"fmt"
	"io"
	"os"

	"github.com/fatih/color"
	"github.com/marmotedu/component-base/pkg/version"
	"github.com/marmotedu/component-base/pkg/version/verflag"
	"github.com/marmotedu/errors"
	"github.com/marmotedu/iam/pkg/log"
	"github.com/moby/term"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var progressMessage = color.GreenString("==>")

// NewApp creates a new application instance based on the given application name,
// binary name, and other options.
func NewApp(name, basename string, opts ...Option) *App {
	app := &App{
		name:     name,
		basename: basename,
	}

	for _, opt := range opts {
		opt(app)
	}

	app.buildCommand()

	return app
}

// App is the main structure of a cli application.
type App struct {
	basename    string
	name        string
	description string
	options     CliOptions
	runFunc     RunFunc
	silence     bool
	noVersion   bool
	noConfig    bool
	commands    []*Command
	args        cobra.PositionalArgs
	cmd         *cobra.Command
}

// Run is used to launch the application.
func (app *App) Run() {
	err := app.cmd.Execute()
	if err != nil {
		fmt.Printf("%v %v\n", color.RedString("Error:"), err)
		os.Exit(1)
	}
}

// Command returns cobra command instance inside the application.
func (app *App) Command() *cobra.Command {
	return app.cmd
}

func (app *App) runCommand(cmd *cobra.Command, _ []string) error {
	printWorkingDir()
	printFlags(cmd.Flags())
	if !app.noVersion {
		verflag.PrintAndExitIfRequested()
	}

	if !app.noConfig {
		var err error
		err = viper.BindPFlags(cmd.Flags())
		if err != nil {
			return err
		}

		err = viper.Unmarshal(app.options)
		if err != nil {
			return err
		}
	}

	if !app.silence {
		log.Infof("%v Starting %s ...", progressMessage, app.name)
		if !app.noVersion {
			log.Infof("%v Version: `%s`", progressMessage, version.Get().ToJSON())
		}
		if !app.noConfig {
			log.Infof("%v Config file used: `%s`", progressMessage, viper.ConfigFileUsed())
		}
	}

	if app.options != nil {
		err := app.applyOptionRules()
		if err != nil {
			return err
		}
	}

	if app.runFunc != nil {
		return app.runFunc(app.basename)
	}

	return nil
}

// applyOptionRules completes options if completable and prints options if printable
func (app *App) applyOptionRules() error {
	completableOptions, ok := app.options.(CompletableOptions)
	if ok {
		err := completableOptions.Complete()
		if err != nil {
			return err
		}
	}

	// validate options
	errs := app.options.Validate()
	if len(errs) != 0 {
		return errors.NewAggregate(errs)
	}

	printableOptions, ok := app.options.(PrintableOptions)
	if ok && !app.silence {
		log.Infof("%v Config: `%s`", progressMessage, printableOptions.String())
	}

	return nil
}

// AddCommand adds sub command to the application.
func (app *App) AddCommand(cmd *Command) {
	app.commands = append(app.commands, cmd)
}

// AddCommands adds multiple sub commands to the application.
func (app *App) AddCommands(cmds ...*Command) {
	app.commands = append(app.commands, cmds...)
}

// Option defines optional parameters for initializing the application structure.
type Option func(app *App)

// WithOptions to open the application's function to read from the command line
// or read parameters from the configuration file.
func WithOptions(opt CliOptions) Option {
	return func(app *App) {
		app.options = opt
	}
}

// RunFunc defines the application's startup callback function.
type RunFunc func(basename string) error

// WithRunFunc is used to set the application startup callback function option.
func WithRunFunc(run RunFunc) Option {
	return func(a *App) {
		a.runFunc = run
	}
}

// WithDescription is used to set the description of the application.
func WithDescription(desc string) Option {
	return func(a *App) {
		a.description = desc
	}
}

// WithSilence sets the application to silent mode, in which the program startup
// information, configuration information, and version information are not
// printed in the console.
func WithSilence() Option {
	return func(a *App) {
		a.silence = true
	}
}

// WithNoVersion set the application does not provide version flag.
func WithNoVersion() Option {
	return func(app *App) {
		app.noVersion = true
	}
}

// WithNoConfig set the application does not provide configs flag.
func WithNoConfig() Option {
	return func(app *App) {
		app.noConfig = true
	}
}

// WithValidArgs set the validation function to valid non-flag arguments.
func WithValidArgs(args cobra.PositionalArgs) Option {
	return func(app *App) {
		app.args = args
	}
}

// WithDefaultValidArgs set default validation function to valid non-flag arguments.
func WithDefaultValidArgs() Option {
	return func(app *App) {
		app.args = func(cmd *cobra.Command, args []string) error {
			for _, arg := range args {
				if len(arg) > 0 {
					return fmt.Errorf("%q does not take any arguments, got %q", cmd.CommandPath(), args)
				}
			}

			return nil
		}
	}
}

func (app *App) buildCommand() {
	cmd := cobra.Command{
		Use:   app.basename,
		Short: app.name,
		Long:  app.description,

		SilenceUsage:  true,
		SilenceErrors: true,
		Args:          app.args,
	}

	cmd.SetOut(os.Stdout)
	cmd.SetErr(os.Stderr)
	cmd.Flags().SortFlags = true
	initFlags(cmd.Flags())

	if len(app.commands) > 0 {
		for _, command := range app.commands {
			cmd.AddCommand(command.cobraCommand())
		}
		cmd.SetHelpCommand(helpCommand(app.basename))
	}

	if app.runFunc != nil {
		cmd.RunE = app.runCommand
	}

	var fss NamedFlagSets
	if app.options != nil {
		fss = app.options.Flags()
		fs := cmd.Flags()
		for _, f := range fss.FlagSets {
			fs.AddFlagSet(f)
		}
	}

	if !app.noVersion {
		verflag.AddFlags(fss.FlagSet("global"))
	}

	if !app.noConfig {
		addConfigFlag(app.basename, fss.FlagSet("global"))
	}

	addGlobalFlags(fss.FlagSet("global"), cmd.Name())
	cmd.Flags().AddFlagSet(fss.FlagSet("global"))

	addCmdTemplate(&cmd, fss)
	app.cmd = &cmd
}

func addCmdTemplate(cmd *cobra.Command, fss NamedFlagSets) {
	usageFmt := "Usage:\n  %s\n"
	cols, _, _ := terminalSize(cmd.OutOrStdout())
	cmd.SetUsageFunc(func(cmd *cobra.Command) error {
		fmt.Fprintf(cmd.OutOrStderr(), usageFmt, cmd.UseLine())
		printSections(cmd.OutOrStderr(), fss, cols)

		return nil
	})
	cmd.SetHelpFunc(func(cmd *cobra.Command, args []string) {
		fmt.Fprintf(cmd.OutOrStdout(), "%s\n\n"+usageFmt, cmd.Long, cmd.UseLine())
		printSections(cmd.OutOrStdout(), fss, cols)
	})
}

func terminalSize(w io.Writer) (int, int, error) {
	outFd, isTerminal := term.GetFdInfo(w)
	if !isTerminal {
		return 0, 0, fmt.Errorf("given writer is no terminal")
	}
	winsize, err := term.GetWinsize(outFd)
	if err != nil {
		return 0, 0, err
	}
	return int(winsize.Width), int(winsize.Height), nil
}

func printWorkingDir() {
	wd, _ := os.Getwd()
	log.Infof("%v WorkingDir: %s", progressMessage, wd)
}
