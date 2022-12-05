package config

type Kpotato struct {
HelloWorldMessage string `hcl:"hello_world_message,attr"`

// ... your config here
}

// DefaultKpotatoConfig returns default config values
func DefaultKpotatoConfig() Kpotato {
	return Kpotato{
	HelloWorldMessage:
		"hello from the default config",
	}
}
