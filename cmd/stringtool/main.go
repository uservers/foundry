package main

import (
	"github.com/sirupsen/logrus"
	"github.com/uservers/foundry/pkg/stringtool"
)

func main() {
	err := stringtool.IsValidDomain("user vers.net")
	if err != nil {
		logrus.Fatalf("Fallo %v", err)
	}

	logrus.Infof("Es dominio valido")
}
