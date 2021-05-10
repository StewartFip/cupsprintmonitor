package logging

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/rs/zerolog"
)

type Logger struct {
	Log *zerolog.Logger
}

func New() *Logger {
	logfile, err := logFile()
	if err != nil {
		fmt.Printf("error encountered: %s\n", err.Error())
	}

	logger := zerolog.New(logfile).With().Timestamp().Logger()

	return &Logger{
		Log: &logger,
	}
}

func logFile() (io.Writer, error) {
	logfile := filepath.Join(os.TempDir(), "cupsprintmonitor", fmt.Sprintf("cupsprintmonitor_%s.log", time.Now().Format("20060102")))

	if err := os.MkdirAll(filepath.Dir(logfile), 0770); err != nil {
		return nil, err
	}

	return os.OpenFile(logfile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
}
