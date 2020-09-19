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
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	osGo "os"
	"os/exec"
)

const (
	DEBIAN = "debian"
	UBUNTU = "ubuntu"
)

var (
	apt = "install -y "
	inlinePackage string
	os string
)

func DebianInstaller(updateBefore bool, upgradeBefore bool, packages []string, currentOs string) error {
	os = currentOs
	manager, err := setPackageManager()
	if err != nil {
		return err
	}
	command := manager + apt
	if err = updateDistrib(updateBefore); err != nil {
		return err
	}
	if err = upgradeDistrib(upgradeBefore); err != nil {
		return err
	}
	command = command + getInlinePackages(packages)
	cmd := exec.Command("echo", "hello")
	if err := printShellResult(cmd); err != nil {
		log.Fatal(err)
	}
	return nil
}

func setPackageManager() (string, error) {
	switch os {
		case DEBIAN:
			return "apt-get ", nil
			break
		case UBUNTU:
			return "apt-get ", nil
			break
		default:
			return "", errors.New("Impossible de déterminer le package manager pour cette os")
	}
	return "", errors.New("Une erreur est survenue setPackageManager")
}

func updateDistrib(update bool) error {
	if update {
		switch os {
			case DEBIAN:
				cmd := exec.Command("apt-get", "update")
				if err := printShellResult(cmd); err != nil {
					log.Fatal(err)
				}
				break
			default:
				return errors.New("Aucun OS défini")

		}
	}
	return nil
}

func upgradeDistrib(upgrade bool) error {
	if upgrade {
		switch os {
			case DEBIAN:
				cmd := exec.Command("apt-get", "upgrade -y")
				if err := printShellResult(cmd); err != nil {
					log.Fatal(err)
				}
				break
			default:
				return errors.New("Aucun OS défini")
		}
	}
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

func printShellResult(cmd *exec.Cmd) error {
	var stdBuffer bytes.Buffer
	mw := io.MultiWriter(osGo.Stdout, &stdBuffer)
	cmd.Stdout = mw
	//cmd.Stderr = mw

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
	fmt.Printf(stdBuffer.String())
	return nil
}