package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	version = "dev" // Will be replaced by GoReleaser with actual version
	commit  = "unknown"
	date    = "unknown"
)

var rootCmd = &cobra.Command{
	Use:     "bw",
	Short:   "Bitwarden CLI",
	Long:    "Bitwarden CLI - Official command-line tool for Bitwarden password manager",
	Version: version,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Bitwarden CLI %s\n", version)
		fmt.Println("Official command-line tool for Bitwarden password manager")
		fmt.Println("\nUsage:")
		fmt.Println("  bw [command]")
		fmt.Println("\nAvailable Commands:")
		fmt.Println("  login        Log in to your Bitwarden account")
		fmt.Println("  logout       Log out of your Bitwarden account")
		fmt.Println("  sync         Sync your vault")
		fmt.Println("  unlock       Unlock your vault")
		fmt.Println("  lock         Lock your vault")
		fmt.Println("  list         List items in your vault")
		fmt.Println("  get          Get an item from your vault")
		fmt.Println("  generate     Generate a password")
		fmt.Println("\nUse 'bw [command] --help' for more information about a command.")
	},
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Bitwarden CLI %s\n", version)
		fmt.Printf("Commit: %s\n", commit)
		fmt.Printf("Date: %s\n", date)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}