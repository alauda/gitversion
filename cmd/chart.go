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

// chartCmd represents the chart command
var chartCmd = &cobra.Command{
	Use:   "chart",
	Short: "generates a new version number for a chart based on a current version and a given minor version",
	Long: `usage: gitversion <chart current version> <target minor version>
	example1: gitversion chart v0.1.2 v0.2 will print v0.2.0
	example2: gitversion chart v0.1.2 v0.1 will print v0.1.3
	example3: gitversion chart v0.1.2 v0.1.3 will print v0.1.3
	example4: gitversion chart v0.1.2 v0.1.2 will print v0.1.2
	example5: gitversion chart v0.1.2 v0.1.0 will print v0.1.3 
	`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 2 && len(args) != 1 {
			fmt.Println(`please provide at least one full current version and an optional target minor version.`)
			os.Exit(1)
			return
		}
		currentVersion, targetVersion := args[0], ""
		if len(args) == 2 {
			targetVersion = args[1]
		}
		res, err := pkg.GetNextChartVersion(currentVersion, targetVersion)
		if err != nil {
			fmt.Println(`Error generating version number: `, err)
			os.Exit(1)
			return
		}
		fmt.Println(res)
	},
}

func init() {
	rootCmd.AddCommand(chartCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// chartCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// chartCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
