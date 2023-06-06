package command

import (
	"fmt"
	"github.com/spf13/cobra"
	"ordent-assessment/config"
	"ordent-assessment/database/seeder"
)

var seedCmd = &cobra.Command{
	Use:   "seeder",
	Short: "Seed database with initial data",
	Long:  "This command will seed the database with initial data.",
	Run: func(cmd *cobra.Command, args []string) {
		configuration := config.New()
		database := config.NewMongoDatabase(configuration)

		//SEEDER REGISTER
		err := seeder.UserSeeder(database)
		if err != nil {
			fmt.Println("Seeder error:", err)
		}

		fmt.Println("Seeder completed!")
	},
}
