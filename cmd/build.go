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

// buildCmd represents the patch command
var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "generates a build version number based on a given minor version",
	Long: `for example: given the following git tags:
	 v0.1, v0.1.b-1, v0.1.2  if v0.1 is given as argument will print v0.1.b-2 as output.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println(`please provide a minor version number to generate build number. e.g: v0.1`)
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
		result := GetBuilderVersion(tags, version)
		fmt.Println(result)
	},
}

func GetBuilderVersion(tags []string, version string) (result string) {
	result = fmt.Sprintf("%v.b-1", version)
	if len(tags) == 0 || (len(tags) == 1 && tags[0] == version) {
		return
	}
	prefix := fmt.Sprintf("%v.b-", version)
	tags = pkg.FilterTags(prefix, tags, strings.HasPrefix)
	if len(tags) == 0 || (len(tags) == 1 && tags[0] == version) {
		return
	}
	highest := pkg.GetHighestPatch(tags, func(build string) string {
		return strings.Replace(build, "b-", "", -1)
	})
	highest++
	result = fmt.Sprintf("%v.b-%d", version, highest)
	return
}

func init() {
	rootCmd.AddCommand(buildCmd)
}
