package logger

import (
	"fmt"
	"github.com/logrusorgru/aurora"
	"strings"
	"time"
)

type LogLevel string

const (
	GameLogLevel       = "GameLogLevel"
	MatchMakingLogLevel   = "MatchMakingLogLevel"
	ConnectionLogLevel = "ConnectionLogLevel"
	BaseLogLevel       = "BaseLogLevel"
	WarnLogLevel       = "WarnLogLevel"
	ErrorLogLevel      = "ErrorLogLevel"
)

var colors = map[LogLevel]aurora.Value{
	GameLogLevel:       aurora.Magenta(GameLogLevel),
	MatchMakingLogLevel:   aurora.Cyan(MatchMakingLogLevel),
	ConnectionLogLevel: aurora.Blue(ConnectionLogLevel),
	BaseLogLevel:       aurora.Gray(15, BaseLogLevel),
	WarnLogLevel:       aurora.Yellow(WarnLogLevel),
	ErrorLogLevel:      aurora.Red(ErrorLogLevel),
}

type EmojigambleLogger struct {
	ActiveLogLevels         []LogLevel
	LogToDatabase           bool
	MongoDBConnectionString string
}

func AllLogLevels() []LogLevel {
	return []LogLevel{
		GameLogLevel,
		MatchMakingLogLevel,
		ConnectionLogLevel,
		BaseLogLevel,
		WarnLogLevel,
		ErrorLogLevel,
	}
}

func (logger *EmojigambleLogger) Log(message string, level LogLevel) {
	logger.log(message, level, logger.LogToDatabase)
}

func (logger *EmojigambleLogger) log(message string, level LogLevel, logToDb bool) {
	for _, l := range logger.ActiveLogLevels {
		if l == level {
			fmt.Println(
				aurora.BgWhite(aurora.Gray(3, time.Now().Format(" 02.01.2006|15:04:05 "))),
				strings.Replace(colors[level].String(), "LogLevel", "", 1), " ",
				message)

			if logToDb {
				// TODO: save log into database
			}
		}
	}
}

func (logger *EmojigambleLogger) SampleLoggerOutput(logToDb bool) {
	fmt.Println(aurora.Red("+START+ Emojigamble logger sample\n"))
	logger.log("Game started!", GameLogLevel, logToDb)
	logger.log("Connected!", ConnectionLogLevel, logToDb)
	logger.log("Base log", BaseLogLevel, logToDb)
	logger.log("Warning!", WarnLogLevel, logToDb)
	logger.log("Error at nothing :)", ErrorLogLevel, logToDb)
	fmt.Println(aurora.Red("\n+END+ Emojigamble logger sample"))
}
