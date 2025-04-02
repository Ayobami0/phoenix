package parser

type Ash struct {
	System struct {
		Meta struct {
			Name        string
			Description string
			Version     string
			Author      string
			Created     string
			Target      string
		}
	}
	Install struct {
		Packages []struct {
			Manager        string
			ManagerCommand string `yaml:"manager_command"`
			Group          string
			Packages       []string
		}
	}
	Services struct {
		DisableCommand string
		EnableCommand  string
		Enable         []string
		Disable        []string
	}
	Environment []struct {
		Name  string
		Value string
	}

	User []struct {
		Username string
		Shell    string
		Groups   []string
	}

	FileSystem struct {
		Directories []struct {
			Path string
			Mode int
		}
		SymLinks []struct {
			Source string
			Target string
		}
	}
	Git []struct {
		Source      string
		Destination string
		Branch      string
	}

	Workflow struct {
		PreSetup []struct {
			Script string
			Args   []string
		} `yaml:"pre_setup"`
		PostSetup []struct {
			Script string
			Args   []string
		} `yaml:"post_setup"`
	}
}
