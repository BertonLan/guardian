package main

import (
	"os"

	"code.cloudfoundry.org/guardian/gqt/cmd/fake_runc/args"
	"github.com/Sirupsen/logrus"
)

func main() {
	f, err := os.Create(args.GetCommandArg("log"))
	if err != nil {
		os.Exit(1)
	}

	logrus.SetFormatter(new(logrus.JSONFormatter))
	logrus.SetOutput(f)
	logrus.Info("guardian-runc-logging-test-info")
	logrus.Warn("guardian-runc-logging-test-warn")
	logrus.Error("guardian-runc-logging-test-error")
	logrus.Print("guardian-runc-logging-test-print")
}
