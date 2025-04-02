package app

import (
	"fmt"
	"os"
	"strings"

	"github.com/Ayobami0/phoenix/internal/config"
	"github.com/Ayobami0/phoenix/internal/parser"
	"github.com/Ayobami0/phoenix/internal/utils"
	p "github.com/Ayobami0/phoenix/internal/utils/pretty"
)

func Spawn(a config.ConfigArgs) error {
	pretty := p.New(a.ToPrintConfig())

	ash, err := parser.ParseAsh(a.AshFile)

	if err != nil {
		pretty.PrintError("Failed to parse configuration file")
		return err
	}
	output := a.OutputFile

	switch output {
	case "-":
		content, err := utils.GenerateShellOutput(a, pretty, *ash)
		if err != nil {
			pretty.PrintError(err.Error())
			return err
		}
		fmt.Println(content)
	default:
		if output == "" {
			output = strings.Split(a.AshFile, ".")[0] + ".sh"
		}

		pretty.PrintLogo()

		pretty.PrintStep("Parsing configuration file")
		pretty.PrintSuccess("Configuration loaded successfully")

		content, err := utils.GenerateShellOutput(a, pretty, *ash)
		if err != nil {
			pretty.PrintError(err.Error())
			return err
		}
		if err != nil {
			pretty.PrintError(err.Error())
			return err
		}

		file, err := os.Create(output)

		if err != nil {
			pretty.PrintError(err.Error())
			return err
		}

		_, err = file.WriteString(content)
		if err != nil {
			pretty.PrintError(err.Error())
			return err
		}

		if a.Executable {
			err := os.Chmod(file.Name(), 0755)
			if err != nil {
				pretty.PrintError(err.Error())
				return err
			}
		}

		err = file.Close()
		if err != nil {
			pretty.PrintError(err.Error())
			return err
		}

		pretty.PrintSuccess(fmt.Sprintf("File generated %s", output))
		pretty.PrintSilentCompletion(fmt.Sprintf("File generated %s", output))
		pretty.PrintFooter()
	}

	return nil
}
