package cmd

import (
	"fmt"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:     "list [directory]",
	Aliases: []string{"ls"},
	Short:   "List files and directories",
	Long:    "List all files and folders available in the specified directory with optional detailed information",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		directory := args[0]
		detailed, _ := cmd.Flags().GetBool("all")

		fileList, err := List(directory)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Printf("Files and directories in \"%s\":\n", directory)
		if detailed {
			printDetailedList(fileList)
		} else {
			printList(fileList)
		}
	},
}

func init() {
	listCmd.Flags().BoolP("all", "a", false, "Show detailed information")
	rootCmd.AddCommand(listCmd)
}

func printList(fileList []FileInfo) {
	maxLen := maxNameLength(fileList)

	header := fmt.Sprintf("%-*s | %s", maxLen, "Name", "Type")
	separator := strings.Repeat("-", len(header))
	fmt.Println(header)
	fmt.Println(separator)

	for _, file := range fileList {
		fmt.Printf("%-*s | %s\n", maxLen, file.Name, file.Type)
	}
}

func printDetailedList(fileList []FileInfo) {
	maxNameLen := maxNameLength(fileList)
	maxSizeLen := maxSizeLength(fileList)
	maxExtLen := maxExtensionLength(fileList)

	header := fmt.Sprintf("%-*s | %-*s | %-*s | %-*s | %s",
		maxNameLen, "Name",
		maxSizeLen, "Size",
		maxExtLen, "Ext",
		25, "Creation Time",
		"Type")
	separator := strings.Repeat("-", len(header))
	fmt.Println(header)
	fmt.Println(separator)

	for _, file := range fileList {
		fmt.Printf("%-*s | %-*d | %-*s | %s | %s\n",
			maxNameLen, file.Name,
			maxSizeLen, file.Size,
			maxExtLen, file.FileExtension,
			file.CreationTime.Format(time.RFC3339),
			file.Type)
	}
}

func maxNameLength(fileList []FileInfo) int {
	maxLen := 0
	for _, file := range fileList {
		if len(file.Name) > maxLen {
			maxLen = len(file.Name)
		}
	}
	return maxLen
}

func maxSizeLength(fileList []FileInfo) int {
	maxLen := 0
	for _, file := range fileList {
		sizeLen := len(fmt.Sprintf("%d", file.Size))
		if sizeLen > maxLen {
			maxLen = sizeLen
		}
	}
	return maxLen
}

func maxExtensionLength(fileList []FileInfo) int {
	maxLen := 0
	for _, file := range fileList {
		if len(file.FileExtension) > maxLen {
			maxLen = len(file.FileExtension)
		}
	}
	return maxLen
}
