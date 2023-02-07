package common

type TestStep interface {
	Render() string
}

type TestVariables interface {
	Get(key string) string
	Len() int
	Keys() []string
	Values() []string
	YAMLString() string
}

type Test interface {
	AgentName() string
	ComponentName() string
	TargetVersion() string
	AppVersion() string
	Name() string
	Equals(Test) bool
	Modified(Test) bool
	Steps() []TestStep
	Setup() []TestStep
	Variables() TestVariables
}

type ConfigurationDiff interface {
	Removed() []Test
	Added() []Test
	Modified() []Test
	String() string
	Len() int
}

type Configuration interface {
	ConfigVersion() string
	MatrixString() string
	Plan() (ConfigurationDiff, error)
	Apply() (ConfigurationDiff, error)
	IndexSuites()
	SetPrevConfig(Configuration)
	Tests() []Test
	Sample(toJSON bool) ([]byte, error)
	Validate(data []byte, isYaml bool) error
}
