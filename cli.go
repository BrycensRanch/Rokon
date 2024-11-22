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
		if len(args) > 0 {
			// Parse args as key=value pairs and set them in Viper
			for _, arg := range args {
				// Split on the first '=' to get the key and value
				parts := strings.SplitN(arg, "=", 2)
				if len(parts) == 2 {
					key := parts[0]
					initialValue := parts[1]

					// Determine the correct type for the value and parse it
					var value interface{}
					if strings.HasPrefix(initialValue, "[") && strings.HasSuffix(initialValue, "]") {
						// Remove the square brackets
						initialValue = initialValue[1 : len(initialValue)-1]
						// Split the values by commas
						items := strings.Split(initialValue, ",")
						// Now parse each item in the list
						var values []interface{}
						for _, item := range items {
							item = strings.TrimSpace(item) // Trim whitespace around items
							// Try parsing as bool, int, or leave as string
							if v, err := strconv.ParseBool(item); err == nil {
								values = append(values, v)
							} else if v, err := strconv.Atoi(item); err == nil {
								values = append(values, v)
							} else if v, err := strconv.ParseFloat(item, 64); err == nil {
								values = append(values, v)
							} else {
								// Otherwise treat it as a string
								values = append(values, item)
							}
						}
						viper.Set(key, values)
						fmt.Printf("Set %s=%v\n", key, values)
					} else {
						// Try to parse the value as a boolean, integer, or leave as string
						if v, err := strconv.ParseBool(initialValue); err == nil {
							value = v
						} else if v, err := strconv.Atoi(initialValue); err == nil {
							value = v
						} else {
							// If not a bool or int, treat it as a string
							value = initialValue
						}
						// Set the key-value pair in Viper
					}
					viper.Set(key, value)
					viper.WriteConfig()
					log.Printf("Set %s=%w\n", key, value)
				} else {
					log.Printf("Invalid argument format: %s. Use key=value.", arg)
				}
			}
		} else {
			allKeys := viper.AllKeys()
			for _, key := range allKeys {
				value := viper.Get(key)
				log.Printf("%s: %v\n", key, value)
			}
			log.Printf("You can set configuration values with the syntax of %s %s key=value", cmd.Root().Name(), cmd.Use)
		}
	},
}

func activateCommandLine(app *gtk.Application, commandLine *gio.ApplicationCommandLine) int {
	args := commandLine.Arguments()
	// cobraArgs := append([]string{rootCmd.Use}, args...)
	// rootCmd.SetArgs(cobraArgs)
	viper.BindPFlags(pflag.CommandLine)
	initCLI()

	cmd, remainingArgs, err := rootCmd.Find(args[1:])
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
	rootCmd.AddCommand(configCmd)
}
