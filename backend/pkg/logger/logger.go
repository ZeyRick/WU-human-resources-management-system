package logger

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const (
	infoColor    = "\033[1;34m%+v\033[0m"
	noticeColor  = "\033[1;36m%+v\033[0m"
	warningColor = "\033[1;33m%+v\033[0m"
	errorColor   = "\033[1;31m%+v\033[0m"
	debugColor   = "\033[0;36m%+v\033[0m"
)

func InitLogger() {
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.DateTime}
	output.FormatLevel = func(i interface{}) string {
		return strings.ToUpper(fmt.Sprintf("| %-6s|", i))
	}
	output.FormatMessage = func(i interface{}) string {
		if i == "" {
			return ""
		}
		return fmt.Sprintf("%s", i)
	}
	log.Logger = zerolog.New(output).With().Timestamp().CallerWithSkipFrameCount(3).Logger()
}

// Print out in console
func Console(msg any) {
	log.Info().Msg(fmt.Sprintf(debugColor, msg))
}

// Print out error with tracing (Showing which code casusing the error)
func Trace(err error) {
	errStack := errors.WithStack(err)
	log.Error().Msg(fmt.Sprintf(errorColor, errStack))
}

// Print out error in red color
func Error(msg string) {
	log.Error().Msg(fmt.Sprintf(errorColor, msg))
}

// Print out in blue color
func Success(msg string) {
	log.Info().Msg(fmt.Sprintf(infoColor, msg))
}

func PrettyJson(v interface{}) {
	// Marshal the struct to JSON
	jsonBytes, err := json.Marshal(v)
	if err != nil {
		Trace(err)
	}
	// Indent the JSON
	var prettyJSON bytes.Buffer
	err = json.Indent(&prettyJSON, jsonBytes, "", "  ")
	if err != nil {
		Trace(err)
	}

	fmt.Println(prettyJSON.String())
}
