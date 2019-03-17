package infrastructures

import (
	"io/ioutil"

	log "github.com/sirupsen/logrus"
	config "github.com/spf13/viper"
)

// SetLog sets log configuration
func SetLog() {

	log.AddHook(NewLogHook().
		SetFormatType(config.GetString("log.format_output")).
		SetLogLevel(config.GetInt("log.level")).
		SetRotateLog(config.GetString("log.rotate")))

	log.SetFormatter(&log.JSONFormatter{})

	if config.GetString("log.format_output") == "stdout" {
		log.SetOutput(ioutil.Discard)
	}
}
