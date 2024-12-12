package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/diamondburned/gotk4/pkg/gio/v2"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
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

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Configure settings",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			allKeys := viper.AllKeys()
			for _, key := range allKeys {
				log.Printf("%s: %v\n", key, viper.Get(key))
			}
			log.Printf("You can set configuration values with the syntax of %s %s key=value", cmd.Root().Name(), cmd.Use)
			return
		}

		for _, arg := range args {
			parts := strings.SplitN(arg, "=", 2)
			if len(parts) != 2 {
				log.Printf("Invalid argument format: %s. Use key=value.", arg)
				continue
			}

			key, initialValue := parts[0], parts[1]
			var value interface{}

			if strings.HasPrefix(initialValue, "[") && strings.HasSuffix(initialValue, "]") {
				initialValue = initialValue[1 : len(initialValue)-1]
				items := strings.Split(initialValue, ",")
				var values []interface{}

				for _, item := range items {
					item = strings.TrimSpace(item)
					values = append(values, parseValue(item))
				}
				value = values
			} else {
				value = parseValue(initialValue)
			}

			viper.Set(key, value)
			if err := viper.WriteConfig(); err != nil {
				log.Printf("Error writing config: %v", err)
				continue
			}
			log.Printf("Set %s=%v\n", key, value)
		}
	},
}

func parseValue(value string) interface{} {
	if v, err := strconv.ParseBool(value); err == nil {
		return v
	} else if v, err := strconv.Atoi(value); err == nil {
		return v
	} else if v, err := strconv.ParseFloat(value, 64); err == nil {
		return v
	} else {
		return value
	}
}


func activateCommandLine(app *gtk.Application, commandLine *gio.ApplicationCommandLine) int {
	args := commandLine.Arguments()
	// cobraArgs := append([]string{rootCmd.Use}, args...)
	// rootCmd.SetArgs(cobraArgs)
	err := viper.BindPFlags(pflag.CommandLine)
	if (err != nil) {
		log.Println("Viper failed to bind the flags to Rokon's configuration.")
		log.Fatal(err.Error())
	}
	initCLI()

	cmd, remainingArgs, err := rootCmd.Find(args[1:])
	if err != nil {
		log.Println("Error finding command:", err)
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
		log.Println("Error executing command:", err)
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
	rootCmd.AddCommand(configCmd)
}
