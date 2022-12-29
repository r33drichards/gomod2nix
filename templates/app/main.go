package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "myapp",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func main() {
	err := rootCmd.Execute()
	if err != nil {
		// todo capture sentry error
		fmt.Println(err)
		os.Exit(1)
	}
}

func addCommands(cmd *cobra.Command, cmds []*cobra.Command) *cobra.Command {
	for _, c := range cmds {
		cmd.AddCommand(c)
	}
	return cmd
}

func init() {

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.rwcli.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd = addCommands(rootCmd, []*cobra.Command{
		{
			Use:   "version",
			Short: "Print the version number of myapp",
			Long:  `All software has versions. This is myapp's`,
			Run: func(cmd *cobra.Command, args []string) {
				fmt.Println("0.0.0")
			},
		},
	})
}
