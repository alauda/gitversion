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
		tags, err := pkg.GetAllTags()
		if err != nil {
			fmt.Println(`Error while fetching tags:`, err)
			os.Exit(1)
			return
		}
		version, err := pkg.PatchVersion(args, tags)
		if err != nil {
			fmt.Println(`Error while genering patch version:`, err)
			os.Exit(1)
			return
		}
		fmt.Println(version)
	},
}

func init() {
	rootCmd.AddCommand(patchCmd)
}
