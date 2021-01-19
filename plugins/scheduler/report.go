package scheduler

import (
	"strconv"

	"github.com/dokku/dokku/plugins/common"
)

// ReportSingleApp is an internal function that displays the app report for one or more apps
func ReportSingleApp(appName, infoFlag string) error {
	if err := common.VerifyAppName(appName); err != nil {
		return err
	}

	flags := map[string]common.ReportFunc{
		"--scheduler-task-count": reportTasks,
	}

	flagKeys := []string{}
	for flagKey := range flags {
		flagKeys = append(flagKeys, flagKey)
	}

	trimPrefix := false
	uppercaseFirstCharacter := true
	infoFlags := common.CollectReport(appName, infoFlag, flags)
	return common.ReportSingleApp("scheduler", appName, infoFlag, infoFlags, flagKeys, trimPrefix, uppercaseFirstCharacter)
}

func reportTasks(appName string) string {
	c, _ := fetchCronEntries(appName)
	return strconv.Itoa(len(c))
}
