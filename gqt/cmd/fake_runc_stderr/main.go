package main

import (
	"io/ioutil"
	"os"

	"code.cloudfoundry.org/guardian/gqt/cmd/fake_runc/args"
)

func main() {
	ioutil.WriteFile(args.GetCommandArg("pid-file"), []byte("13"), 0777)

	stderrContents := ""
	for i := 0; i < 5000; i++ {
		stderrContents += "I am a bad runC\n"
	}
	os.Stderr.WriteString(stderrContents)
	os.Exit(100)
}
