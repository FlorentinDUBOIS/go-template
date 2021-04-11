package mod

type (
	Log struct {
		Level int32 `json:"level" yaml:"level"`
	}

	// Configuration structure is here to describe the application configuration
	Configuration struct {
		Log Log `json:"log" yaml:"log"`
	}
)
