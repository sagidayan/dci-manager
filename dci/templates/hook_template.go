package templates

import (
	_ "embed"
)

//go:embed hook_template.yaml
var HookTemplate string
