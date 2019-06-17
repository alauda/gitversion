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

var inplace bool

// buildCmd represents the patch command
var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "generates a build version number based on a given minor version",
	Long: `for example: given the following git tags:
	 v0.1, v0.1.b-1, v0.1.2  if v0.1 is given as argument will print v0.1.b-2 as output.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		implementation := pkg.BuildGit
		if inplace {
			implementation = pkg.BuildInplace
		}
		result, err := implementation(args, pkg.GetAllTags)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		fmt.Println(result)
	},
}

func init() {
	rootCmd.AddCommand(buildCmd)

	buildCmd.Flags().BoolVarP(&inplace, "inplace", "i", false, "if added will not use the git repository as reference but only the given arguments")
}
