package cli

import (
	"fmt"

	"github.com/Ayobami0/phoenix/internal/app"
	"github.com/Ayobami0/phoenix/internal/config"
	"github.com/urfave/cli/v2"
)

const (
	VERSION = "1.0.0"
)

func New() *cli.App {
	n := cli.App{
		Name:        "phoenix",
		Usage:       "system configuration management tool",
		UsageText:   "phoenix COMMAND [OPTIONS] <ashname>",
		Version:     VERSION,
		Description: `Phoenix configures systems based on "ash" configuration files.`,
		Commands: []*cli.Command{
			RiseCmd(),
			SpawnCmd(),
		},
	}

	return &n
}

func RiseCmd() *cli.Command {
	slient := cli.BoolFlag{Name: "silent", Aliases: []string{"s"}, Usage: "run in silent mode without printing status messages"}
	exclude := cli.StringSliceFlag{Name: "exclude", Usage: "skip specific components (comma-separated)"}

	return &cli.Command{
		Name:      "rise",
		Usage:     "Apply configuration from an ash file to the current system",
		UsageText: "rise [OPTIONS] <ashname>",
		Action:    RiseAction,
		Flags: []cli.Flag{
			&slient,
			&exclude,
		},
	}
}

func RiseAction(ctx *cli.Context) error {
	if ctx.NArg() != 1 {
		return fmt.Errorf("Command line args needed")
	}
	ashName := ctx.Args().First()
	silent := ctx.Bool("silent")
	exclude := ctx.StringSlice("exclude")

	app.Rise(config.ConfigArgs{Silent: silent, AshFile: ashName, Excludes: exclude})
	return nil
}

func SpawnCmd() *cli.Command {
	silent := cli.BoolFlag{Name: "silent", Aliases: []string{"s"}, Usage: "generate script without printing status messages"}
	out := cli.StringFlag{Name: "out", Aliases: []string{"o"}, Usage: "write output to FILE default is name of ash provided. use - to write to stdout"}
	executable := cli.BoolFlag{Name: "executable", Usage: "make output file executable (chmod +x)"}
	compress := cli.BoolFlag{Name: "compress", Usage: "generate a compressed, minimal script"}
	shell := cli.StringFlag{Name: "shell", Usage: "output shell: bash, zsh, powershell", DefaultText: "bash"}
	exclude := cli.StringSliceFlag{Name: "exclude", Usage: "skip specific components (comma-separated)"}

	return &cli.Command{
		Name:      "spawn",
		Usage:     "Generate a portable shell script from an ash file",
		UsageText: "spawn [OPTIONS] <ashname>",
		Action:    SpawnAction,
		Flags: []cli.Flag{
			&silent,
			&shell,
			&compress,
			&executable,
			&out,
			&exclude,
		},
	}
}

func SpawnAction(ctx *cli.Context) error {
	if ctx.NArg() != 1 {
		return fmt.Errorf("Command line args needed")
	}
	ashName := ctx.Args().First()
	output := ctx.String("out")
	shell := ctx.String("shell")
	silent := ctx.Bool("silent")
	compress := ctx.Bool("compress")
	executable := ctx.Bool("executable")
	exclude := ctx.StringSlice("exclude")

	app.Spawn(config.ConfigArgs{
		Silent: silent, AshFile: ashName, OutputFile: output,
		Compress: compress, Executable: executable, Shell: shell,
		Excludes: exclude})
	return nil
}
