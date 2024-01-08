package log

import (
	"github.com/rs/zerolog"
	"gopkg.in/natefinch/lumberjack.v2"

	"os"

	"copilot-gpt4-service/config"
)

type Logger struct {
	Log zerolog.Logger
}

func NewLogger() *Logger {
	logger := &lumberjack.Logger{
		Filename:   "logs/log.log",
		MaxSize:    500,
		MaxBackups: 3,
		MaxAge:     38,
		LocalTime:  true,
		Compress:   true,
	}
	consoleWriter := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: "2006-01-02 15:04:05"}
	multi := zerolog.MultiLevelWriter(logger, consoleWriter)
	zlog := zerolog.New(multi).With().Timestamp().Logger()

	if config.ConfigInstance.Debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
		zlog.Debug().Msg("Debug mode enabled")
	}

	return &Logger{Log: zlog}
}

var ZLog *Logger = NewLogger()
