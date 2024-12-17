package logger

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"os"
	"sync"
	"time"
)

var PROD = "prod"

var logMutex = sync.Mutex{}

const LogColorReset = "\033[0m"

const LogColorRed = "\033[38;2;255;0;0m"
const LogColorGreen = "\033[38;2;0;255;0m"
const LogColorBlue = "\033[38;2;50;90;255m"

const LogColorYellow = "\033[38;2;255;210;0m"
const LogColorPurple = "\033[38;2;160;0;250m"
const LogColorCyan = "\033[38;2;0;240;250m"
const LogColorLime = "\033[38;2;220;255;0m"

const LogColorGray = "\033[38;2;180;180;180m"
const LogColorWhite = "\033[38;2;255;255;255m"

// LogW is used for printing warnings
func LogW(input ...interface{}) {

	if os.Getenv("env") == PROD {
		nexDevLog(LogColorYellow, "[WARN]", input...)
	} else {
		log.Warn().Any("log", input).Send()
	}
}

// LogWf is used for printing warnings/*
func LogWf(format string, input ...any) {
	if os.Getenv("env") == PROD {
		nexDevLog(LogColorYellow, "[WARN]", fmt.Sprintf(format, input...))
	} else {
		log.Warn().Any("log", []interface{}{format, input}).Send()
	}
}

// LogS indicates success
func LogS(input ...interface{}) {
	if os.Getenv("env") == PROD {
		nexDevLog(LogColorGreen, "[SUCCESS]", input...)
	} else {
		log.Log().Any("level", "success").Any("log", input).Send()
	}

}

// LogSf indicates success formatted
func LogSf(format string, input ...any) {
	if os.Getenv("env") == PROD {
		nexDevLog(LogColorGreen, "[SUCCESS]", fmt.Sprintf(format, input...))
	} else {
		log.Log().Any("level", "success").Any("log", []interface{}{format, input}).Send()
	}
}

// LogInfo indicates info
func LogInfo(input ...interface{}) {
	if os.Getenv("env") == PROD {
		nexDevLog(LogColorBlue, "[INFO]", input...)
	} else {
		log.Info().Any("log", input).Send()
	}
}

// LogInfoF indicates info formatted
func LogInfoF(format string, input ...any) {
	if os.Getenv("env") == PROD {
		nexDevLog(LogColorBlue, "[INFO]", fmt.Sprintf(format, input...))
	} else {
		log.Info().Any("log", []interface{}{format, input}).Send()
	}

}

// LogDebug indicates debug
func LogDebug(input ...interface{}) {
	if os.Getenv("env") == PROD {
		nexDevLog(LogColorCyan, "[DEBUG]", input...)
	} else {
		log.Log().Any("level", "notice").Any("debug", input).Send()
	}
}

// LogDebugF indicates debug formatted
func LogDebugF(format string, input ...any) {
	if os.Getenv("env") == PROD {
		nexDevLog(LogColorCyan, "[DEBUG]", fmt.Sprintf(format, input...))
	} else {
		log.Log().Any("level", "notice").Any("debug", []interface{}{format, input}).Send()
	}

}

// LogE indicates error
func LogE(input ...interface{}) {
	if os.Getenv("env") == PROD {
		nexDevLog(LogColorRed, "[ERROR]", input...)
	} else {
		log.Error().Any("log", input).Send()
		os.Exit(1)
	}
}

// LogEf indicates error formatted
func LogEf(format string, input ...any) {
	if os.Getenv("env") == PROD {
		nexDevLog(LogColorRed, "[ERROR]", fmt.Sprintf(format, input...))
		os.Exit(1)
	} else {
		log.Error().Any("log", []interface{}{format, input}).Send()
		os.Exit(1)
	}
}

// LogD indicates internal data
func LogD(input ...interface{}) {
	if os.Getenv("env") == PROD {
		nexDevLog(LogColorPurple, "[DATA]", input...)
	} else {
		log.Log().Any("level", "data").Any("log", input).Send()
	}
}

// LogDf indicates internal data formatted
func LogDf(format string, input ...any) {
	if os.Getenv("env") == PROD {
		nexDevLog(LogColorPurple, "[DATA]", fmt.Sprintf(format, input...))
	} else {
		log.Log().Any("level", "data").Any("log", []interface{}{format, input}).Send()
	}
}

// nexDevLog log helper function
func nexDevLog(colorCode string, keyWord string, input ...interface{}) {
	logMutex.Lock()
	defer logMutex.Unlock()

	cTime := time.Now()
	fmt.Print(colorCode)
	fmt.Print(cTime.Format("2006/01/02 15:04:05"), " ")
	fmt.Print(keyWord, " ")
	fmt.Print(input...)
	fmt.Print(LogColorReset, "\n")
}
