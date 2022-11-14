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
		gen.config = config
	})
}
