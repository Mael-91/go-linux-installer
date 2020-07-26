/*
Copyright © 2020 Mael Constantin <mael.constantin@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	release bool
	issue bool
	pr bool
	github bool
	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Print the current version of the application",
		Long: "Current version of go-linux-installer. You can find all the releases in the release tab on github",
		Run: version,
	}
)

func init() {
	rootCmd.AddCommand(versionCmd)
	versionCmd.Flags().BoolVarP(&github, "github", "g", false, "Affiche le lien vers le github")
	versionCmd.Flags().BoolVarP(&release, "release", "r", false, "Affiche le lien vers les release")
	versionCmd.Flags().BoolVarP(&issue, "issue", "i", false, "Affiche le lien vers les issues")
	versionCmd.Flags().BoolVarP(&pr, "pull-request", "p", false, "Affiche le lien vers les pull request")
}

func version(cmd *cobra.Command, args []string) {
	var isTrue = false
	if github {
		isTrue = true
		fmt.Printf("Gitub Repository : %s\n", application.Application.Github.Repo)
	}
	if release {
		isTrue = true
		fmt.Printf("Application release : %s\n", application.Application.Github.Release)
	}
	if issue {
		isTrue = true
		fmt.Printf("You have a problem, open a issue : %s\n", application.Application.Github.Issue)
	}
	if pr {
		isTrue = true
		fmt.Printf("Check pull request : %s\n", application.Application.Github.PullRequest)
	}
	if !isTrue {
		fmt.Printf("Current version : %s", application.Application.Version)
	}
}