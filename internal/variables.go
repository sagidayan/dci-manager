package internal

import (
	"fmt"
	"os"
)

var (
	ConfigFile          = getVarDefault("DCIM_CONFIG", "dcim.yaml")
	StateFolderLocation = getVarDefault("DCIM_STATE_DIR", ".state/")
	StateFilePath       = fmt.Sprintf("%s/dcim_state.json", StateFolderLocation)
	JobsFolderLocation  = getVarDefault("DCIM_JOBS_DIR", "./jobs/")
)

func getVarDefault(evar string, _default string) string {
	val := os.Getenv(evar)
	if len(val) == 0 {
		return _default
	}
	return val
}
