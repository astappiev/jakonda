package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
	"strings"
)

var printFlag bool
var confirmFlag bool
var sourceFormatFlag string
var imageFormatFlag string

// rmRawCmd represents the rmRaw command
var rmRawCmd = &cobra.Command{
	Use:   "rm-raw [path-to-folder] [OPTIONS]",
	Short: "Remove raw files if the corresponding jpg file exists",
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
		var sourceFormats []string
		for _, rawFormat := range strings.Split(sourceFormatFlag, ",") {
			sourceFormats = append(sourceFormats, strings.TrimSpace(rawFormat))
		}

		var imageFormats []string
		for _, imageFormat := range strings.Split(imageFormatFlag, ",") {
			imageFormats = append(imageFormats, strings.TrimSpace(imageFormat))
		}

		countRaw, candidates := FindSourceIfImageExists(args[0], sourceFormats, imageFormats)
		fmt.Printf("Found %d raw files.\n", countRaw)

		if len(candidates) > 0 && (confirmFlag || printFlag) {
			for _, candidate := range candidates {
				if confirmFlag {
					err := os.Remove(candidate)
					if err != nil {
						fmt.Printf("error removing %s: %s\n", candidate, err.Error())
					}
					if printFlag {
						fmt.Printf("%s removed\n", candidate)
					}
				} else if printFlag {
					fmt.Printf("%s can be removed\n", candidate)
				}
			}
		}

		if countRaw > 0 && (len(candidates) == 0 || !confirmFlag) {
			fmt.Printf("%d files can be removed.\n", len(candidates))
		} else if (len(candidates) > 0) && confirmFlag {
			fmt.Printf("%d files removed.\n", len(candidates))
		}

		if len(candidates) > 0 && !confirmFlag {
			fmt.Printf("\nUse --confirm to delete the files\n")
		}
	},
}

func init() {
	rootCmd.AddCommand(rmRawCmd)

	rmRawCmd.Flags().StringVarP(&sourceFormatFlag, "source-format", "s", ".raw, .dng, .raf, .cr2, .nef, .arw", "Specify the source format to look for")
	rmRawCmd.Flags().StringVarP(&imageFormatFlag, "image-format", "i", ".jpg, .jpeg", "Specify the image format to look for")

	rmRawCmd.Flags().BoolVarP(&printFlag, "print", "p", false, "Print the files to be deleted")
	rmRawCmd.Flags().BoolVarP(&confirmFlag, "confirm", "c", false, "Delete the found files")
}

func FindSourceIfImageExists(path string, sourceFormats []string, imageFormats []string) (int, []string) {
	countRaw := 0
	candidates := make([]string, 0)
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if contains(sourceFormats, strings.ToLower(filepath.Ext(info.Name()))) {
			countRaw++

			image, err := lookupImage(path, imageFormats)
			if err != nil {
				fmt.Printf("error reading %s: %s\n", info.Name(), err.Error())
			}
			if image != nil {
				candidates = append(candidates, path)
			}
		}

		return nil
	})

	if err != nil {
		fmt.Printf("error reading %s: %s\n", path, err.Error())
	}

	return countRaw, candidates
}

func lookupImage(prefix string, imageFormats []string) (*string, error) {
	paths, err := filepath.Glob(strings.TrimSuffix(prefix, filepath.Ext(prefix)) + ".*")
	if err != nil {
		return nil, err
	}
	if len(paths) > 1 {
		for _, path := range paths {
			fileName := filepath.Base(path)
			if contains(imageFormats, strings.ToLower(filepath.Ext(fileName))) {
				return &fileName, nil
			}
		}
	}
	return nil, nil
}

func contains[T comparable](slice []T, needle T) bool {
	for _, v := range slice {
		if v == needle {
			return true
		}
	}
	return false
}
