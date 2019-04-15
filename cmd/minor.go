// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"os"

	"github.com/alauda/gitversion/pkg"
	"github.com/spf13/cobra"
)

// minorCmd represents the minor command
var minorCmd = &cobra.Command{
	Use:   "minor",
	Short: "upgrade a minor version",
	Long: `update a minor version. for example:
	given: v0.1 as argument, should print v0.2
	if none is given will return an error`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("please provide a major.minor version to upgrade minor version. e.g v0.1")
			os.Exit(1)
			return
		}
		version := args[0]
		new, err := pkg.BumpMinor(version)
		if err != nil {
			fmt.Println("error", err)
			os.Exit(1)
			return
		}
		fmt.Println(new)
	},
}

func init() {
	rootCmd.AddCommand(minorCmd)
}
