package command

import (
	"flag"
	"fmt"
	"github.com/mitchellh/cli"
	"strings"
)

// JoinCommand is a Command implementation that tells a running Serf
// agent to join another.
type RegisterCommand struct {
	Ui cli.Ui
}

func (c *RegisterCommand) Help() string {
	helpText := `
Usage: watchdog register [options] path/to/config.json ...

  Registers a new process with the Watchdog agent for monitoring.

  The configuration file can be specified in either JSON or TOML format
  depending on your preference and ease of integration.

  If the process is configured to start on load (start_on_load=true) it will be
  started immediately after registering.

  The path to this config file will be watched for changes and automatically
  reloaded if it changes.

  NOTE: This command is idempotent and will return success if the process is
  already registered.

Options:

  -no-start                 Don't start the process even if configured to do so.
  -no-watch                 Don't watch this configuration file for changes.
  -rpc-addr=127.0.0.1:6673  RPC address of the Watchdog agent.
`
	return strings.TrimSpace(helpText)
}

func (c *RegisterCommand) Run(args []string) int {
	var noStartOnLoad bool = false
	var noWatch bool = false

	cmdFlags := flag.NewFlagSet("join", flag.ContinueOnError)
	cmdFlags.Usage = func() { c.Ui.Output(c.Help()) }
	cmdFlags.BoolVar(&noStartOnLoad, "no-start", false, "no-start")
	cmdFlags.BoolVar(&noWatch, "no-watch", false, "no-watch")
	if err := cmdFlags.Parse(args); err != nil {
		return 1
	}

	configPaths := cmdFlags.Args()
	fmt.Println(configPaths, noStartOnLoad)
	if len(configPaths) == 0 {
		c.Ui.Error("At least one configuration file must be specified.")
		c.Ui.Error("")
		c.Ui.Error(c.Help())
		return 1
	}

	return 0
}

func (c *RegisterCommand) Synopsis() string {
	return "Register a new process with the Watchdog agent"
}