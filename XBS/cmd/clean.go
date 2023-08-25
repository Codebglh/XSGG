package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	targetFolder  string
	defaultFolder = "middlefile"
)

// deleteFolderCmd represents the deleteFolder command
var deleteFolderCmd = &cobra.Command{
	Use:   "clean",
	Short: "Delete a specific target folder and its contents",
	RunE:  deleteFolder,
}

func init() {
	rootCmd.AddCommand(deleteFolderCmd)

	// Here you will define your flags and configuration settings.
	deleteFolderCmd.Flags().StringVarP(&targetFolder, "folder", "f", defaultFolder, "Target folder to delete")
}

func deleteFolder(cmd *cobra.Command, args []string) error {
	// Check if target folder exists
	if _, err := os.Stat(targetFolder); os.IsNotExist(err) {
		return errors.New("target folder does not exist")
	}

	// Delete the target folder and its contents
	err := os.RemoveAll(targetFolder)
	if err != nil {
		return fmt.Errorf("failed to delete folder %s: %w", targetFolder, err)
	}

	fmt.Printf("Deleted folder: %s\n", targetFolder)

	return nil
}
