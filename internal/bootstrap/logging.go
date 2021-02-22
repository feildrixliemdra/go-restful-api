package bootstrap

import log "github.com/sirupsen/logrus"

func InitLogger() {
	log.SetFormatter(&log.JSONFormatter{
		FieldMap: log.FieldMap{
			log.FieldKeyTime:  "timestamp",
			log.FieldKeyLevel: "level",
			log.FieldKeyMsg:   "message",
		},
	})

	log.SetReportCaller(true)
}
