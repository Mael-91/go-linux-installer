package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"go-linux-installer/app"
)

var versionCmd = &cobra.Command{
	Use: "version",
	Short: "Print the current version of the application",
	Long: "Current version of go-linux-installer. You can find all the releases in the release tab on github",
	Run: version,
}

var config app.Config

func init() {
	config = app.LoadApplicationConf()
	rootCmd.AddCommand(versionCmd)
}

func version(cmd *cobra.Command, args []string) {
	/**
	* tag --release affiche le lien vers les release
	* tag --issue affiche le lien vers les issues
	* tag --pr affiche le lien vers les pull request
	* tag --authors affiche le nom des auteurs
	* tag --contributor affiche le lien vers la page des contributeur
	* On peut seulement faire 1 tag à la fois
	 */
	fmt.Printf("%s", config.Application.Version)
}