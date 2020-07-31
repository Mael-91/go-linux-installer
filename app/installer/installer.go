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
package installer

import (
	"errors"
	"fmt"
)

const (
	DEBIAN = "debian"
	UBUNTU = "ubuntu"
)

var (
	apt = "install -y "
	inlinePackage string
	os string
	sudo bool
)

func DebianInstaller(updateBefore bool, upgradeBefore bool, packages []string, withSudo bool, currentOs string) error {
	os = currentOs
	sudo = withSudo
	manager, err := setPackageManager()
	if err != nil {
		return err
	}
	command := manager + apt
	if sudo {
		command = manager + apt
	}
	err = updateDistrib(manager, updateBefore)
	if err != nil {
		return err
	}
	err = upgradeDistrib(manager, upgradeBefore)
	if err != nil {
		return err
	}
	command = command + getInlinePackages(packages)
	fmt.Printf("%s", command)
	return nil
}

func setPackageManager() (string, error) {
	var cmd string
	switch os {
		case DEBIAN:
			cmd = "apt-get "
			break
		case UBUNTU:
			cmd = "apt-get "
			break
		default:
			return "", errors.New("Impossible de déterminer le package manager pour cette os")
	}
	switch sudo {
		case true:
			return "sudo " + cmd, nil
		case false:
			return cmd, nil
		default:
			return "", errors.New("Impossible de déterminer si l'action doit être effectué en mode sudo")
	}
}

func updateDistrib(cmd string, update bool) error {
	if !update {
		return nil
	}
	// Vérification pour les mise à jour
	return nil
}

func upgradeDistrib(cmd string, upgrade bool) error {
	if !upgrade {
		return nil
	}
	// Mise à jour de la distribution
	return nil
}

func getInlinePackages(packages []string) string {
	for index, value := range packages {
		switch index {
			case 0:
				inlinePackage = value
				break
			default:
				inlinePackage += " " + value
		}
	}
	return inlinePackage
}