/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/yellow-canary/dora/internal/auth"
	"github.com/yellow-canary/dora/pkg/dora"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// calculateCmd represents the calculate command
var calculateCmd = &cobra.Command{
	Use:   "calculate",
	Short: "Calculate all DORA four keys metrics",
	Long: `Calculate a snapshot of the four keys DORA software delivery metrics:
Deployment Frequency, Lead Time for Changes, Time to Restore Services and Change Failure Rate
`,
	Run: handleCalculate,
}

func init() {
	rootCmd.AddCommand(calculateCmd)

	calculateCmd.PersistentFlags().StringP("repository", "r", "", `Github repository in the "OWNER/REPO" format.`)

}

func handleCalculate(cmd *cobra.Command, args []string) {

	viper.BindPFlags(cmd.Flags())

	// Generate a Github client
	ctx := context.Background()
	token, isSet := os.LookupEnv("GH_TOKEN")
	if isSet == false {
		log.Fatal("\"GH_TOKEN\" environment variable is not set.")
		os.Exit(1)
	}
	if token == "" {
		log.Fatal("\"GH_TOKEN\" environment variable is empty.")
		os.Exit(1)
	}
	client := auth.GetGithubClient(ctx, token)

	// Specify the owner and repository name
	ghRepo := viper.GetString("repository")

	owner := strings.Split(ghRepo, "/")[0]
	repo := strings.Split(ghRepo, "/")[1]

	// Calculate DORA 4 keys metrics
	deploymentFrequency, err := dora.CalculateDeploymentFrequency(client, owner, repo)
	if err != nil {
		log.Fatal("Error calculating Deployment Frequency:", err)
	}

	leadTimeToChange, err := dora.CalculateLeadTimeToChange(client, owner, repo)
	if err != nil {
		log.Fatal("Error calculating Lead Time to Change:", err)
	}

	changeFailureRate, err := dora.CalculateChangeFailureRate(client, owner, repo)
	if err != nil {
		log.Fatal("Error calculating Change Failure Rate:", err)
	}

	timeToRestoreService, err := dora.CalculateTimeToRestoreService(client, owner, repo)
	if err != nil {
		log.Fatal("Error calculating Time to Restore Service:", err)
	}

	// Print the DORA metrics
	fmt.Printf("Deployment Frequency: %.2f releases per week\n", deploymentFrequency)
	fmt.Printf("Lead Time to Change: %.2f hours\n", leadTimeToChange.Hours())
	fmt.Printf("Change Failure Rate: %.2f%%\n", changeFailureRate*100)
	fmt.Printf("Time to Restore Service: %.2f hours\n", timeToRestoreService.Hours())
}
