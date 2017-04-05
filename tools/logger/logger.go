package logger

import (
	"encoding/json"
	"os"
	"syscall"

	"github.com/Sirupsen/logrus"
	"github.com/finalist736/seabattle_resultslogapi/config"
)

var stdout, stderr *logrus.Logger

//var log2 *logrus.New()

func JsonStdOut(name string, jsn interface{}) {
	ba, _ := json.MarshalIndent(jsn, "", "	")
	StdOut().Printf(name+": %v", ba)
}

func StdOut() *logrus.Logger {
	return stdout
}

func StdErr() *logrus.Logger {
	return stderr
}

func ReloadLogs() {
	if config.GetConfiguration().Logpath != "" {

		err := os.MkdirAll(config.GetConfiguration().Logpath, 0764)
		if err != nil {
			panic(err)
		}

		stdoutfile, err := os.OpenFile(config.GetConfiguration().Logpath+"/stdout.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0764)
		if err != nil {
			panic(err)
		}
		_, err = stdoutfile.Seek(0, os.SEEK_END)
		if err != nil {
			panic(err)
		}
		err = syscall.Dup2(int(stdoutfile.Fd()), int(os.Stdout.Fd()))
		if err != nil {
			panic(err)
		}

		stderrfile, err := os.OpenFile(config.GetConfiguration().Logpath+"/stderr.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0764)
		if err != nil {
			panic(err)
		}
		_, err = stderrfile.Seek(0, os.SEEK_END)
		if err != nil {
			panic(err)
		}
		err = syscall.Dup2(int(stderrfile.Fd()), int(os.Stderr.Fd()))
		if err != nil {
			panic(err)
		}
	}

	var err error

	stdout = logrus.New()
	stdout.Level, err = logrus.ParseLevel(config.GetConfiguration().LogLevelStdOut)
	if err != nil {
		stdout.Level = logrus.InfoLevel
	}
	stdout.Out = os.Stdout

	stderr = logrus.New()
	stderr.Level, err = logrus.ParseLevel(config.GetConfiguration().LogLevelStdErr)
	if err != nil {
		stderr.Level = logrus.InfoLevel
	}
	stderr.Out = os.Stderr
}
