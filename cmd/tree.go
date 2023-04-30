package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var treeCmd = &cobra.Command{
	Use:   "tree [path-to-folder]",
	Short: "Prints the directory tree for the specified folder",
	Args: func(cmd *cobra.Command, args []string) error {
		if err := cobra.ExactArgs(1)(cmd, args); err != nil {
			return err
		}

		if stat, err := os.Stat(args[0]); err != nil {
			return fmt.Errorf("unexisting or invalid path specified: %s", args[0])
		} else if !stat.IsDir() {
			return fmt.Errorf("a directory is expected: %s", args[0])
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		printTree(args[0], 0)
	},
}

func init() {
	rootCmd.AddCommand(treeCmd)
}

func printTree(path string, depth int) {
	entries, err := os.ReadDir(path)
	if err != nil {
		fmt.Printf("error reading %s: %s\n", path, err.Error())
		return
	}

	printLeaf(path, depth)
	for _, entry := range entries {
		if (entry.Type() & os.ModeSymlink) == os.ModeSymlink {
			fullPath, err := os.Readlink(filepath.Join(path, entry.Name()))
			if err != nil {
				fmt.Printf("error reading link: %s\n", err.Error())
			} else {
				printLeaf(entry.Name()+" -> "+fullPath, depth+1)
			}
		} else if entry.IsDir() {
			printTree(filepath.Join(path, entry.Name()), depth+1)
		} else {
			// we do not print files
		}
	}
}

func printLeaf(entry string, depth int) {
	indent := ""
	if depth > 0 {
		entry = filepath.Base(entry)
		indent = strings.Repeat("|   ", depth-1) + "|-- "
	}
	fmt.Print(indent + entry + "\n")
}
