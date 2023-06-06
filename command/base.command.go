package command

import "github.com/spf13/cobra"

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(seedCmd)
}

var rootCmd = &cobra.Command{
	Use:   "myapp",
	Short: "My App short description",
	Long:  `My App long description`,
}
