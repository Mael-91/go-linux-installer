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
	"encoding/json"
	"errors"
	"github.com/spf13/cobra"
	"go-linux-installer/app/installer"
	"io/ioutil"
	"os"
	"strings"
)

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install packages from config file",
	Long: "Installe les paquets depuis un fichier de configuration ou depuis l'argument --install-packages",
	RunE: install,
}

var (
	packages []string
	config string
	noConfigFile bool
	distrib string
	sudo bool
	updateBeforeInstall bool
	upgradeBeforeInstall bool
)

const (
	DEBIAN = "debian"
	UBUNTU = "ubuntu"
)

func init() {
	rootCmd.AddCommand(installCmd)
	installCmd.Flags().StringVarP(&config, "config", "c", "", "Défini un fichier de configuration (Required)")
	installCmd.Flags().StringArrayVarP(&packages, "install-packages", "i", nil, "Installe une liste de paquet depuis l'argument")
	installCmd.Flags().BoolVarP(&noConfigFile,"no-config-file", "", false, "Permet d'installer des paquets en ligne de commande sans fichier de configuration")
	installCmd.Flags().BoolVarP(&sudo, "sudo", "s", false, "L'installateur doit il utiliser sudo ?")
	installCmd.Flags().BoolVarP(&updateBeforeInstall, "update-before-install", "u", true, "L'installateur doit il télécharger les mise à jours de la distribution avant l'installation des paquets (Default : true)")
	installCmd.Flags().BoolVarP(&upgradeBeforeInstall, "upgrade-before-install", "g", false, "L'installateu doit il mettre à jours la distribution avant l'installation des paquets (Default : false")
	installCmd.Flags().StringVarP(&distrib, "os", "o", "", "Permet de définir l'os sur lequel l'installation ce fait")
}

func install(cmd *cobra.Command, args []string) error {
	if config != "" {
		iConfig, err := unmarshalConfig()
		if err != nil {
			return err
		}
		err = getInstallerManager(iConfig)
		if err != nil {
			return err
		}
	}
	if noConfigFile {
		iConfig := InstallerConfigFile{
			GoLinuxInstaller: InstallerConfig{
				Os: strings.ToLower(distrib),
				Sudo: sudo,
				UpdateBeforeInstall: updateBeforeInstall,
				UpgradeBeforeInstall: upgradeBeforeInstall,
				Packages: packages,
			},
		}
		err := getInstallerManager(iConfig)
		if err != nil {
			return err
		}
	}
	return nil
}

/**
* Décode le fichier de configuration depuis un format json
*/
func unmarshalConfig() (configFile InstallerConfigFile, err error) {
	jsonFile, err := os.Open(config)
	if err != nil {
		return configFile, errors.New("Impossible d'ouvrir le fichier\n" + err.Error())
	}
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return configFile, err
	}

	var c InstallerConfigFile

	err = json.Unmarshal(byteValue, &c)
	if err != nil {
		return configFile, errors.New("Unable to decode into a struct.\n" + err.Error())
	}

	return c, nil
}

/**
* Retourne vers la méthode qui est responsable
* de l'installation des packages en fonction de la distribution
*/
func getInstallerManager(cf InstallerConfigFile) error {
	conf := cf.GoLinuxInstaller
	currentOs := strings.ToLower(conf.Os)
	if currentOs == "" {
		return errors.New("Aucun OS (nom de la distribution linux) n'a été configuré.")
	}
	switch currentOs {
		case DEBIAN:
			return installer.DebianInstaller(conf.UpdateBeforeInstall, conf.UpgradeBeforeInstall, conf.Packages, currentOs)
		case UBUNTU:
			return installer.DebianInstaller(conf.UpdateBeforeInstall, conf.UpgradeBeforeInstall, conf.Packages, currentOs)
	}
	return errors.New("Impossible de trouver le manager d'installation pour cette distribution")
}