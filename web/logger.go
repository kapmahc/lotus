package web

import log "github.com/Sirupsen/logrus"

type injectLogger struct {
}

func (p *injectLogger) Debugf(format string, args ...interface{}) {
	log.Debugf(format, args...)
}

func init() {
	log.SetLevel(log.DebugLevel)
	if IsProduction() {
		log.SetLevel(log.InfoLevel)
	}
}
