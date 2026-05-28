package entity

type BunnyCommandResolver string

type BunnyCommandOpts struct {
	Name      string               `yaml:"name"`
	Aliases   []string             `yaml:"aliases"`
	BaseURL   string               `yaml:"base_url"`
	QueryPath string               `yaml:"query_path"`
	Resolver  BunnyCommandResolver `yaml:"resolver"`
}

type BunnyConfig struct {
	BunnyCommands []BunnyCommandOpts `yaml:"commands"`
}

const (
	QueryBunnyCommandResolver BunnyCommandResolver = "QUERY_ONLY"
)
