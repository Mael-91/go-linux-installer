package app

import (
	"errors"
	"log"
)

type system struct {
	os string
}

func init() {
	// Détection de l'os et ajout dans la struct system
	sys := &system{os: "debian"}
	err := sys.Updater()
	if err != nil {
		log.Fatal(err)
	}
}

func (u *system) Updater() error {
	if u.os == "debian" || u.os == "ubuntu" {
		err := debianUpdater()
		return err
	}
	return errors.New("Impossible de trouver la distribution sur lequel le programme s'exécute")
}

func debianUpdater() error {
	return nil
}

func detectOs() (string, error) {
	return "", nil
}