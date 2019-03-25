// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
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
	"strings"

	"github.com/alauda/gitversion/pkg"
	"github.com/spf13/cobra"
)

// patchCmd represents the patch command
var patchCmd = &cobra.Command{
	Use:   "patch",
	Short: "generates a patch version number based on a given minor version",
	Long: `for example: given the following git tags:
	 v0.1, v0.1.1, v0.1.2  if v0.1 is given as argument will print v0.1.3 as output.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println(`please provide a minor version number to generate patch number. e.g: v0.1`)
			os.Exit(1)
			return
		}
		version := args[0]
		tags, err := pkg.GetAllTags()
		if err != nil {
			fmt.Println("an error occurred while fetching tags:", err)
			os.Exit(1)
			return
		}
		if len(tags) == 0 || (len(tags) == 1 && tags[0] == version) {
			fmt.Println(fmt.Sprintf("%v.0", version))
			return
		}
		tags = pkg.FilterTags(version, tags, strings.HasPrefix)
		if len(tags) == 0 || (len(tags) == 1 && tags[0] == version) {
			fmt.Println(fmt.Sprintf("%v.0", version))
			return
		}
		highest := pkg.GetHighestPatch(tags)
		highest++
		fmt.Println(fmt.Sprintf("%v.%d", version, highest))
	},
}

func init() {
	rootCmd.AddCommand(patchCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// patchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// patchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
