package app

import (
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path"
	"strings"

	"github.com/Ayobami0/phoenix/internal/config"
	"github.com/Ayobami0/phoenix/internal/parser"
	"github.com/Ayobami0/phoenix/internal/utils"
	p "github.com/Ayobami0/phoenix/internal/utils/pretty"
)

func Rise(a config.ConfigArgs) error {

	pretty := p.New(a.ToPrintConfig())

	pretty.ClearScreen()

	pretty.PrintLogo()

	pretty.PrintStep("Parsing configuration file")
	ash, err := parser.ParseAsh(a.AshFile)

	if err != nil {
		pretty.PrintError("Failed to parse configuration file")
		return err
	}

	pretty.PrintSuccess("Configuration loaded successfully")

	pretty.Println()
	pretty.PrintStep("Running pre-setup scripts")
	for _, v := range ash.Workflow.PreSetup {
		utils.RunCommand(v.Script, v.Args, pretty)
	}

	pretty.Println()
	pretty.PrintStep("Installing packages..")
	for _, v := range ash.Install.Packages {
		if v.ManagerCommand == "" {
			pretty.PrintWarning(fmt.Sprintf("Skipping packages in group %s, manager_command not defined\n", v.Group))
			continue
		}
		pretty.PrintFormat("%s%s%s Installing %s%s%s packages using %s%s%s\n",
			p.MAGENTA, p.STEP1, p.RESET,
			p.BOLD, v.Group, p.RESET,
			p.BOLD, v.Manager, p.RESET)
		utils.InstallPackage(v.ManagerCommand, v.Packages, pretty)
	}
	pretty.PrintSuccess("Packages installed")

	pretty.Println()
	pretty.PrintStep("Starting services...")
	if ash.Services.EnableCommand == "" {
		pretty.PrintWarning("Services not enabled... No command provided")
	} else {
		pretty.PrintFormat("%s%s%s Enabling services: %s\n", p.MAGENTA, p.STEP1,
			p.RESET, strings.Join(ash.Services.Enable, ", "))
		utils.RunCommand(ash.Services.EnableCommand, ash.Services.Enable, pretty)
	}

	if ash.Services.DisableCommand == "" {

		pretty.PrintWarning("Services not disabled... No command provided")
	} else {
		pretty.PrintFormat("%s%s%s Disabling services: %s\n", p.MAGENTA, p.STEP1,
			p.RESET, strings.Join(ash.Services.Enable, ", "))
		utils.RunCommand(ash.Services.DisableCommand, ash.Services.Disable, pretty)
	}

	pretty.PrintFormat("%s%s%s Setting environment...", p.MAGENTA, p.STEP1, p.RESET)
	home, err := os.UserHomeDir()

	if err != nil {
		pretty.PrintError(err.Error())
		return err
	}

	pretty.Println()
	profilePath := path.Join(home, ".profile")
	proFile, err := os.OpenFile(profilePath, os.O_APPEND|os.O_WRONLY, 0644)
	defer proFile.Close()

	if err != nil {
		pretty.PrintError(err.Error())
		return err
	}

	var pBuffer strings.Builder

	pBuffer.WriteString("# Phoenix environments")

	for _, v := range ash.Environment {
		pBuffer.WriteString(fmt.Sprintf(`export %s=%s\n`, v.Name, v.Value))
	}

	_, err = proFile.WriteString(pBuffer.String())
	if err != nil {
		pretty.PrintError(fmt.Sprintf("failed to write to profile %s", err.Error()))
		return err
	}
	utils.RunCommand("source", []string{profilePath}, pretty)
	pretty.PrintSuccess("Environment set")

	pretty.Println()
	pretty.PrintStep("Setting up users...")
	for _, v := range ash.User {
		utils.RunCommand("sudo", []string{"useradd", "-m", "-s", v.Shell, "-G", strings.Join(v.Groups, ","), v.Username}, pretty)
		pretty.PrintSuccess(fmt.Sprintf("Created user %s", v.Username))
	}

	pretty.Println()
	pretty.PrintStep("Setting up file systems...")
	pretty.PrintFormat("%s%s%sCreating directories", p.MAGENTA, p.STEP1, p.RESET)
	for _, v := range ash.FileSystem.Directories {
		err := os.MkdirAll(v.Path, fs.FileMode(v.Mode))

		if err != nil {
			pretty.PrintError(fmt.Sprintf("Failed to create directory %s", v.Path))
			return err
		}

		pretty.PrintSuccess(fmt.Sprintf("Created directory %s", v.Path))
	}

	pretty.Println("Creating symbolic links")
	for _, v := range ash.FileSystem.SymLinks {
		err := os.Symlink(v.Source, v.Target)

		if err != nil {
			pretty.PrintWarning(fmt.Sprintf("Failed to create symlink from %s to %s: %v", v.Source, v.Target, err))
		} else {
			pretty.PrintSuccess(fmt.Sprintf("Created symlink from %s to %s", v.Source, v.Target))
		}
	}

	for _, v := range ash.FileSystem.SymLinks {
		err := os.Symlink(v.Source, v.Target)

		if err != nil {
			return err
		}
	}

	pretty.Println()
	pretty.PrintStep("Setting up git repositories..")
	git_path, err := exec.LookPath("git")

	if err != nil {
		pretty.PrintError("Could not find git in PATH. Install it and run the command again")
		return err
	}
	for _, v := range ash.Git {
		clone := "clone"

		if v.Branch != "" {
			clone = fmt.Sprintf("%s -b %s", clone, v.Branch)
		}

		pretty.PrintFormat("%s%s%s Cloning %s to %s\n", p.MAGENTA, p.STEP1, p.RESET, v.Source, v.Destination)
		utils.RunCommand(git_path, []string{
			clone,
			v.Source,
			v.Destination,
		},
			pretty,
		)
	}

	pretty.Println()
	pretty.PrintStep("Running post-setup scripts")
	for _, v := range ash.Workflow.PostSetup {
		utils.RunCommand(v.Script, v.Args, pretty)
	}

	pretty.Println()
	pretty.PrintFooter()
	pretty.PrintSilentCompletion("System configured successfully!")

	return nil
}
