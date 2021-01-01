package chat

import (
	"os"

	log "github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

// init prepare Logrus logging system to perform logging for this Chat Server
func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&prefixed.TextFormatter{
		ForceColors:     false,
		ForceFormatting: true,
		FullTimestamp:   true,
		TimestampFormat: "02/01/2006 15:04:05",
	})

	// Output to stdout instead of the default stderr
	log.SetOutput(os.Stdout)

	// Level for debugging
	log.SetLevel(log.DebugLevel)
}
