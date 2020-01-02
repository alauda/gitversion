// Copyright © 2020 NAME HERE <EMAIL ADDRESS>
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

	"github.com/alauda/gitversion/pkg"
	"github.com/spf13/cobra"
)

// buildCommitCmd represents the buildCommit command
var buildCommitCmd = &cobra.Command{
	Use:   "build-commit",
	Short: "generate sa build version that can be traced using a commit hash given a minor version",
	Long: `for example: given a minor version v2.6: 
	will generate v2.6-abcdefg given abcdefg is the short commit id for the current commit
	`,
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		if len(args) < 1 {
			err = fmt.Errorf("Needs at least one argument")
			return
		}
		var version string
		if version, err = pkg.GitDescribe("--always"); err != nil {
			return err
		}
		finalVersion := fmt.Sprintf("%s-%s", args[0], version)
		fmt.Println(finalVersion)
		return
	},
}

func init() {
	rootCmd.AddCommand(buildCommitCmd)
}
