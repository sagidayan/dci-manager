package templates

import (
	_ "embed"
)

//go:embed settings_template.yaml
var SettingsTemplate string
