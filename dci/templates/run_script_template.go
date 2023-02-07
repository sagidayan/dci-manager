package templates

import (
	_ "embed"
)

//go:embed run_script_openshift_app_template.sh
var RunScriptTemplateOpenShiftApp string

//go:embed run_script_openshift_template.sh
var RunScriptTemplateOpenShift string

//go:embed run_script_rhel_template.sh
var RunScriptTemplateRHEL string
