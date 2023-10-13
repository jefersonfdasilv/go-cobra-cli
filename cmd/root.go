package cmd

import (
	"database/sql"
	"github.com/jefersonfdasilv/go-cobra-cli/internal/database"
	"os"

	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/cobra"
)

type RunEFunc func(cmd *cobra.Command, args []string) error

func GetDb() *sql.DB {
	db, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		panic(err)
	}
	return db
}

func GetCategoryDB(db *sql.DB) database.Category {
	return *database.NewCategory(db)
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "go-cobra-cli",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
