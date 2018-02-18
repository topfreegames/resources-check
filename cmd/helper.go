package cmd

import "github.com/sirupsen/logrus"

func configureLogger() *logrus.Logger {
	ll := logrus.InfoLevel
	switch Verbose {
	case 0:
		ll = logrus.InfoLevel
	case 1:
		ll = logrus.WarnLevel
	case 3:
		ll = logrus.DebugLevel
	}

	var log = logrus.New()
	if json {
		log.Formatter = new(logrus.JSONFormatter)
	}
	log.Level = ll

	return log
}
