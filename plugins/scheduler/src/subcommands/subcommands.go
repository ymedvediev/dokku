package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/dokku/dokku/plugins/common"
	"github.com/dokku/dokku/plugins/scheduler"

	flag "github.com/spf13/pflag"
)

// main entrypoint to all subcommands
func main() {
	parts := strings.Split(os.Args[0], "/")
	subcommand := parts[len(parts)-1]

	var err error
	switch subcommand {
	case "list":
		args := flag.NewFlagSet("scheduler:list", flag.ExitOnError)
		args.Parse(os.Args[2:])
		appName := args.Arg(0)
		err = scheduler.CommandList(appName)
	case "report":
		args := flag.NewFlagSet("scheduler:report", flag.ExitOnError)
		osArgs, infoFlag, flagErr := common.ParseReportArgs("scheduler", os.Args[2:])
		if flagErr == nil {
			args.Parse(osArgs)
			appName := args.Arg(0)
			err = scheduler.CommandReport(appName, infoFlag)
		}
	default:
		common.LogFail(fmt.Sprintf("Invalid plugin subcommand call: %s", subcommand))
	}

	if err != nil {
		common.LogFail(err.Error())
	}
}
