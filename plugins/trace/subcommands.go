package trace

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/dokku/dokku/plugins/common"
)

// CommandOn turns off trace mode
func CommandOff() error {
	common.LogInfo1("Disabling trace mode")

	dokkurcPath := filepath.Join(common.MustGetEnv("DOKKU_ROOT"), ".dokkurc")
	tracefilePath := filepath.Join(dokkurcPath, "DOKKU_TRACE")
	if !common.FileExists(tracefilePath) {
		return nil
	}

	if !common.DirectoryExists(dokkurcPath) {
		if err := os.MkdirAll(dokkurcPath, 0755); err != nil {
			return fmt.Errorf("Unable to create .dokkurc directory: %s", err.Error())
		}
	}

	if err := os.Remove(tracefilePath); err != nil {
		return fmt.Errorf("Unable to remove trace file: %s", err.Error())
	}

	return nil
}

// CommandOn turns on trace mode
func CommandOn() error {
	common.LogInfo1("Enabling trace mode")

	dokkurcPath := filepath.Join(common.MustGetEnv("DOKKU_ROOT"), ".dokkurc")
	tracefilePath := filepath.Join(dokkurcPath, "DOKKU_TRACE")
	if !common.DirectoryExists(dokkurcPath) {
		if err := os.MkdirAll(dokkurcPath, 0755); err != nil {
			return fmt.Errorf("Unable to create .dokkurc directory: %s", err.Error())
		}
	}

	lines := []string{"export DOKKU_TRACE=1"}
	if err := common.WriteSliceToFile(tracefilePath, lines); err != nil {
		return fmt.Errorf("Unable to write trace file: %s", err.Error())
	}

	return nil
}
