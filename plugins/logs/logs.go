package logs

import (
	"fmt"

	"github.com/dokku/dokku/plugins/common"
)

var (
	// DefaultProperties is a map of all valid logs properties with corresponding default property values
	DefaultProperties = map[string]string{
		"vector-sink": "",
	}

	// GlobalProperties is a map of all valid global logs properties
	GlobalProperties = map[string]bool{
		"vector-sink": true,
	}
)

// VectorImage contains the default vector image to run
const VectorImage = "timberio/vector:0.11.X-debian"

// GetFailedLogs outputs failed deploy logs for a given app
func GetFailedLogs(appName string) error {
	common.LogInfo2Quiet(fmt.Sprintf("%s failed deploy logs", appName))
	s := common.GetAppScheduler(appName)
	if err := common.PlugnTrigger("scheduler-logs-failed", s, appName); err != nil {
		return err
	}
	return nil
}
