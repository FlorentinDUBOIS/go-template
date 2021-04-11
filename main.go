package main

import (
	"github.com/FlorentinDUBOIS/go-template/pkg/cmd"
	"github.com/sirupsen/logrus"
)

func main() {
	cmd, err := cmd.RootCmd()
	if nil != err {
		logrus.WithError(err).Fatal("could not instantiate the command line interface")
	}

	if err = cmd.Execute(); nil != err {
		logrus.WithError(err).Fatal("could not execute the application")
	}
}
