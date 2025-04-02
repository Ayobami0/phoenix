package config

type ConfigArgs struct {
	Silent bool
	Executable bool
	Excludes []string
	Components []string
	Compress bool
	Shell string
	AshFile string
  OutputFile string
}

type PrintConfig struct {
	Silent bool
	Compress bool
}

func (a ConfigArgs) ToPrintConfig() PrintConfig {
	return PrintConfig{Silent: a.Silent, Compress: a.Compress}
}
