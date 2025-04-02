package utils

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/Ayobami0/phoenix/internal/config"
	"github.com/Ayobami0/phoenix/internal/parser"
	p "github.com/Ayobami0/phoenix/internal/utils/pretty"
)

const (
	BASH = "#!/bin/bash"
	ZSH  = "#!/bin/zsh"
)

var SHELL_MAP = map[string]string{
	"bash": BASH,
	"zsh":  ZSH,
}

func GetShell(shell string) string {
	if s, ok := SHELL_MAP[shell]; ok {
		return s
	}

	return BASH
}

func InstallPackage(manager_cmd string, packages []string, pretty p.PrettyPrint) {
	cmd := strings.Split(manager_cmd, " ")

	args := append(cmd[1:], packages...)

	RunCommand(cmd[0], args, pretty)
}

func RunCommand(cmd string, args []string, pretty p.PrettyPrint) {
	c := exec.Command(cmd, args...)
	fmt.Printf("%s%s%s Running: %s %s\n", p.MAGENTA, p.STEP2, p.RESET, cmd, strings.Join(args, " "))

	c.Stdin = os.Stdin
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr

	if err := c.Run(); err != nil {
		pretty.PrintFormat("%s%s Command failed: %v%s\n", p.BOLD, p.RED, err, p.RESET)
		os.Exit(1)
	}
}

func GenerateShellOutput(a config.ConfigArgs, pretty p.PrettyPrint, ash parser.Ash) (string, error) {

	var fileBuffer *strings.Builder

	fileBuffer = &strings.Builder{}

	pretty.Fprintf(fileBuffer, "%s\n", GetShell(a.Shell))
	pretty.Fprintln(fileBuffer)

	author := "Author: " + ash.System.Meta.Author
	target := "Target: " + ash.System.Meta.Target
	desc := "Description: " + ash.System.Meta.Description
	ashname := "Generated from ash: " + ash.System.Meta.Name
	created := "Created At: " + ash.System.Meta.Created
	version := "Version: " + ash.System.Meta.Version
	headText := "Phoenix Configuration Shell Script"

	maxLen := max(
		len(author),
		len(target),
		len(desc),
		len(ashname),
		len(created),
		len(version),
		len(headText),
	)

	var line strings.Builder

	for i := 0; i < maxLen+3; i++ {
		line.WriteString("=")
	}

	pretty.FprintComment(fileBuffer, line.String())
	pretty.FprintComment(fileBuffer, headText)
	pretty.FprintComment(fileBuffer, ashname)
	pretty.FprintComment(fileBuffer, line.String())
	pretty.FprintComment(fileBuffer, desc)
	pretty.FprintComment(fileBuffer, version)
	pretty.FprintComment(fileBuffer, author)
	pretty.FprintComment(fileBuffer, created)
	pretty.FprintComment(fileBuffer, target)
	pretty.FprintComment(fileBuffer, line.String())

	pretty.Fprintln(fileBuffer, "set -e")
	pretty.Fprintln(fileBuffer, "set -u")

	pretty.FprintEcho(fileBuffer, "🔥 Phoenix Configuration Script 🔥", fmt.Sprintf("Setting up %s environment...", ashname))

	pretty.FprintComment(fileBuffer, "===== Pre Setup Script =====")
	pretty.FprintEcho(fileBuffer, "Running Pre setup script...")

	if len(ash.Workflow.PreSetup) == 0 {
		pretty.FprintEcho(fileBuffer, "Skipping: No pre setup script!")
	} else {
		var preScriptString strings.Builder

		for _, v := range ash.Workflow.PreSetup {
			preScriptString.WriteString(fmt.Sprintf("./%s %s", v.Script, strings.Join(v.Args, " ")))
		}
		pretty.Fprintln(fileBuffer, preScriptString.String())
		pretty.FprintEcho(fileBuffer, "✅ Pre setup completed!", "")
	}

	pretty.FprintComment(fileBuffer, "===== Package Installation =====")
	pretty.FprintEcho(fileBuffer, "📦 Installing packages...")

	for _, v := range ash.Install.Packages {
		pretty.FprintComment(fileBuffer, fmt.Sprintf("%s group (%s)", v.Group, v.Manager))
		pretty.FprintEcho(fileBuffer, fmt.Sprintf("Installing %s packages...", v.Group))

		pretty.Fprintln(fileBuffer, fmt.Sprintf(`%s %s`, v.ManagerCommand, strings.Join(v.Packages, " ")))
	}

	pretty.FprintEcho(fileBuffer, "✅ Package installation complete!", "")

	pretty.FprintComment(fileBuffer, "===== Service Configuration =====")
	pretty.FprintEcho(fileBuffer, "🔧 Configuring services...")

	pretty.FprintComment(fileBuffer, "Enable Services")

	if ash.Services.EnableCommand == "" {
		pretty.FprintEcho(fileBuffer, "Ignoring services: no enable command provided")
	} else {
		pretty.FprintEcho(fileBuffer, fmt.Sprintf("Enabling services: %s", strings.Join(ash.Services.Enable, " ")))

		for _, v := range ash.Services.Enable {
			pretty.Fprintln(fileBuffer, fmt.Sprintf("%s %s", ash.Services.EnableCommand, v))
		}
	}

	if ash.Services.DisableCommand == "" {
		pretty.FprintEcho(fileBuffer, "Ignoring services: no disable command provided")
	} else {
		pretty.FprintEcho(fileBuffer, fmt.Sprintf("Disabling services: %s", strings.Join(ash.Services.Disable, " ")))

		for _, v := range ash.Services.Disable {
			pretty.Fprintln(fileBuffer, fmt.Sprintf("%s %s", ash.Services.DisableCommand, v))
		}
	}

	pretty.FprintEcho(fileBuffer, "✅ Service configuration complete!", "")

	pretty.FprintComment(fileBuffer, "===== Environment Configuration =====")
	pretty.FprintEcho(fileBuffer, "🌍 Setting up environment variables...")

	if len(ash.Environment) == 0 {
		pretty.FprintEcho(fileBuffer, "No environment variable to set")
	} else {
		var environmentString strings.Builder

		for _, v := range ash.Environment {
			environmentString.WriteString(fmt.Sprintf("export %s=%s\n", v.Name, v.Value))
		}
		pretty.FprintComment(fileBuffer, "Add environment variables to ~/.profile")
		pretty.Fprintf(fileBuffer, `cat << 'EOF' >> ~/.profile

# Phoenix configured environment variables
%s
EOF
`, environmentString.String())

		pretty.FprintEcho(fileBuffer, "Environment variables added to ~/.profile",
			"Run 'source ~/.profile' to apply changes to current session",
			"✅ Environment configuration complete!",
			"",
		)
	}
	pretty.FprintComment(fileBuffer, "===== User Configuration =====")
	pretty.FprintEcho(fileBuffer, "👤 Adding users...")
	if len(ash.User) == 0 {
		pretty.FprintEcho(fileBuffer, "No users to add")
	} else {
		var userString strings.Builder

		for _, v := range ash.User {
			userString.WriteString(fmt.Sprintf("sudo useradd -m -s %s -G %s %s", v.Shell, strings.Join(v.Groups, ","), v.Username))
		}
		pretty.FprintComment(fileBuffer, "Add users to the system")
		pretty.Fprintln(fileBuffer, userString.String())
		pretty.FprintEcho(fileBuffer, "✅ Users added!", "")
	}

	pretty.FprintComment(fileBuffer, "===== Git Resources =====")
	pretty.FprintEcho(fileBuffer, "📥 Fetching Git resources...")

	if len(ash.Git) == 0 {
		pretty.FprintEcho(fileBuffer, "No git repositories to clone")
	} else {
		var gitString strings.Builder

		for _, v := range ash.Git {
			gitString.WriteString(fmt.Sprintln("git clone -b %s %s %s", v.Branch, v.Source, v.Destination))
		}
		pretty.Fprintln(fileBuffer, gitString.String())
		pretty.FprintEcho(fileBuffer, "✅ Git resources fetched!", "")
	}

	pretty.FprintComment(fileBuffer, "===== Post Setup Script =====")
	pretty.FprintEcho(fileBuffer, "Running Post setup script...")

	if len(ash.Workflow.PostSetup) == 0 {
		pretty.FprintEcho(fileBuffer, "Skipping: No post setup script!")
	} else {
		var postScriptString strings.Builder

		for _, v := range ash.Workflow.PostSetup {
			postScriptString.WriteString(fmt.Sprintf("./%s %s", v.Script, strings.Join(v.Args, " ")))
		}
		pretty.Fprintln(fileBuffer, postScriptString.String())
		pretty.FprintEcho(fileBuffer, "✅ Post setup completed!", "")
	}

	pretty.FprintEcho(fileBuffer, "🎉 Configuration complete!",
		"Some changes may require a system restart to take effect.",
		"",
		"Phoenix setup finished! Your system has risen from the ashes.")

	return fileBuffer.String(), nil
}
