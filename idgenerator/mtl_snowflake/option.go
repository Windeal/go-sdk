package mtl_snowflake

type Option interface {
	apply(gen *Generator)
}

// optionFunc wraps a func so it satisfies the Option interface.
type optionFunc func(generator *Generator)

func (f optionFunc) apply(gen *Generator) {
	f(gen)
}

// WithConfig :
func WithConfig(config *Config) Option {
	return optionFunc(func(gen *Generator) {
		// gen.config = config
		if config.TimeBit > 0 {
			gen.config.TimeBit = config.TimeBit
		}

		if config.MachineIDBit > 0 {
			gen.config.MachineIDBit = config.MachineIDBit
		}

		if config.TimelineBit > 0 {
			gen.config.TimelineBit = config.TimelineBit
		}

		if config.SeqBit > 0 {
			gen.config.SeqBit = config.SeqBit
		}
		if config.Epoch > 0 {
			gen.config.Epoch = config.Epoch
		}
		if config.MachineID > 0 {
			gen.config.MachineID = config.MachineID
		}
	})
}
