package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/diamondburned/gotk4/pkg/gio/v2"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   getExecutableName(),
	Short: "Control your Roku with style.",
	Long:  `GTK4 application that controls your Roku TV. Whether that be with your keyboard, mouse, or controller.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Default behavior when no subcommand or flags are provided
		log.Println(cmd.UsageString())
	},
	Version: version,
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(applicationInfo())
	},
}

func activateCommandLine(app *gtk.Application, commandLine *gio.ApplicationCommandLine) int {
	args := commandLine.Arguments()
	cobraArgs := append([]string{rootCmd.Use}, args...)
	rootCmd.SetArgs(cobraArgs)

	cmd, remainingArgs, err := rootCmd.Find(args)
	if err != nil {
		fmt.Println("Error finding command:", err)
		log.Println(rootCmd.UsageString())
		return 1
	}
	if cmd == nil {
		log.Println("Unknown usage")
		rootCmd.UsageFunc()
		return 1
	}
	cmd.SetArgs(remainingArgs)
	if err := cmd.Execute(); err != nil {
		fmt.Println("Error executing command:", err)
		return 1
	}
	if len(args) > 1 {
		return 0
	}
	app.SetFlags(gio.ApplicationFlagsNone)
	app.Activate()
	return 0
}

func getExecutableName() string {
	executablePath := os.Args[0]

	_, executableName := filepath.Split(executablePath)

	return executableName
}

func initCLI() {
	rootCmd.AddCommand(versionCmd)
	rootCmd.Use = getExecutableName()
}
