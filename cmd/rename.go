/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"errors"
	"os"

	"github.com/boomsquared/tap/pkg/file"
	"github.com/spf13/cobra"
)

var prefix string

// renameCmd represents the rename command
var renameCmd = &cobra.Command{
	Use:   "rename",
	Short: "rename multiple files",
	Long:  `tap rename [folder] --by [size] --prefix[whatever]`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires a path argument")
		}
		if _, err := os.Stat(args[0]); os.IsNotExist(err) {
			return errors.New("path does not exist")
		}

		fs, err := os.Stat(args[0])
		if err != nil {
			return err
		}
		if !fs.IsDir() {
			return errors.New("path is not a directory")
		}

		return nil

	},
	Run: func(cmd *cobra.Command, args []string) {
		rn := file.NewRenamer(file.Operation{})
		rn.Load(args[0], by)
		rn.Rename(prefix)
	},
}

func init() {
	rootCmd.AddCommand(renameCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// renameCmd.PersistentFlags().String("foo", "", "A help for foo")

	renameCmd.Flags().StringVar(&by, "by", "size", "file will be assigned to folder base on this flag")
	renameCmd.Flags().StringVar(&prefix, "prefix", "file", "prefix of the file")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// renameCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
