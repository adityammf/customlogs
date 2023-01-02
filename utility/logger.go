//TO EDIT

package utility

import (
	"fmt"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

const (
	stdout = "stdout"
)

/*
Logger is the structure to use to produce logs.
Country is mandatory.
It is best to provide an app version to help in correlate logs with application upgrade.
UseCase is for identifying the app that is producing the logs.
*/
type (
	Logger struct {
		Country string
		// UseCase    string
		ShowDebug  bool
		AppVersion string
		logger     *zap.Logger
	}
)

var Log Logger

func (log *Logger) newLogger() {
	dyn := zap.NewAtomicLevel()
	dyn.SetLevel(zap.InfoLevel)
	if log.ShowDebug {
		dyn.SetLevel(zap.DebugLevel)
	}
	var outLogFile, outLogDir string
	dir, err := getBinDir()
	if err != nil {
		outLogFile = stdout
	} else {
		_ = dir
		outLogDir = "/logs"
		outLogFile = outLogDir + "/preservedLogs.log"
		err = os.Mkdir(outLogDir, 0755)
		if err != nil && os.IsNotExist(err) {
			outLogFile = stdout
		} else {
			err = nil
			file, err := os.OpenFile(outLogFile, os.O_RDONLY|os.O_CREATE, 0666)
			if err != nil {
				outLogFile = stdout
			}
			file.Close()
		}
	}

	w := zapcore.AddSync(&lumberjack.Logger{
		Filename:   outLogFile,
		MaxSize:    50, // megabytes
		MaxBackups: 5,
	})
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(zapcore.EncoderConfig{
			TimeKey:        "dateTime",
			LevelKey:       "level",
			NameKey:        "logger",
			MessageKey:     "useCase",
			StacktraceKey:  "stacktrace",
			EncodeLevel:    zapcore.LowercaseLevelEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		}),
		w,
		dyn,
	)

	var initialFields []zap.Field

	initialFields = append(initialFields, zap.String("country", log.Country),
		zap.String("appVersion", log.AppVersion))

	core = core.With(initialFields)

	log.logger = zap.New(core)
	defer log.logger.Sync()
}

/*
WithSessionID is to be used to update logger and add a sessionId for all subsequent invocations.
The logger is cloned to avoid concurrency issues.
*/
func (log *Logger) WithSessionID(sessionID string) Logger {
	log.newLogger()
	log.logger = log.logger.With(zap.String("sessionId", sessionID))
	return *log

}

/*
Debug is for all logs that are for developers.
*/
func (log *Logger) Debug(msg string, fields ...zapcore.Field) {
	if log.logger == nil {
		log.newLogger()
	}
	log.logger.Debug(msg, fields...)
}

/*
Debug is for all formatted logs that are for developers.
*/
func (log *Logger) Debugf(msg string, vars ...interface{}) {
	if log.logger == nil {
		log.newLogger()
	}
	logMsg := fmt.Sprintf(msg, vars...)
	log.Debug(logMsg)
}

/*
Info is for all logs that are for informative purpose.
Incoming requests for example, or outgoing results.
*/
func (log *Logger) Info(msg string, fields ...zapcore.Field) {
	if log.logger == nil {
		log.newLogger()
	}
	log.logger.Info(msg, fields...)
}

/*
Infof is for all formatted logs that are for informative purpose.
Incoming requests for example, or outgoing results.
*/
func (log *Logger) Infof(msg string, vars ...interface{}) {
	if log.logger == nil {
		log.newLogger()
	}
	logMsg := fmt.Sprintf(msg, vars...)
	log.Info(logMsg)
}

// func (log *Logger) APIhit(v []byte, msg string) {
// 	var data PlainText
// 	err := json.Unmarshal(v, &data)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	log.Info(msg, zap.String("eventType", "incoming request"), zap.String("errorDetails", ""), zap.Any("parameters", data))
// }

// func (log *Logger) OutgoingResp(v []byte, msg string) {
// 	var data OutgoingResponse
// 	err := json.Unmarshal(v, &data)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	log.Info(msg, zap.String("eventType", "outgoing response"), zap.String("errorDetails", ""), zap.Any("parameters", data))
// }

// func (log *Logger) DbReading(v []byte, msg string) {
// 	var data DBRead
// 	err := json.Unmarshal(v, &data)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	log.Info(msg, zap.String("eventType", "database reading"), zap.String("errorDetails", ""), zap.Any("parameters", data))
// }

/*
Error level is for any error that is not recoverable : a retry will not solve the problem.
The administrator should have action to fix the problem.
It is for inconsistency in data for example.
*/
func (log *Logger) Error(msg string, err error, fields ...zapcore.Field) {
	if log.logger == nil {
		log.newLogger()
	}
	// fields = append(fields, zap.Error(err))
	log.logger.Error(msg, fields...)
}

/*
Errorf is for errors. Executable should be stopped after this error.
It is itended to cover cases like network error, database overload, lack of resources, etc.
*/
func (log *Logger) Errorf(msg string, vars ...interface{}) {
	if log.logger == nil {
		log.newLogger()
	}
	//logMsg := fmt.Sprintf(msg, vars...)
	log.Error(msg, nil)
}

/*
ErrorMsg is for errors. Executable should be stopped after this error.
It is itended to cover cases like network error, database overload, lack of resources, etc.
*/
func (log *Logger) ErrorMsg(msg string, err error) {
	log.Error(msg, err, zap.String("errorDetails", err.Error()), zap.String("eventType", "error occurred"))
}

/*
Fatal is for recoverable errors. Executable should be stopped after this error.
It is itended to cover cases like network error, database overload, lack of resources, etc.
*/
func (log *Logger) Fatal(msg string, err error, fields ...zapcore.Field) {
	if log.logger == nil {
		log.newLogger()
	}
	log.logger.Fatal(msg, zap.Error(err))
}

/*
Fatalf is for recoverable errors. Executable should be stopped after this error.
It is itended to cover cases like network error, database overload, lack of resources, etc.
*/
func (log *Logger) Fatalf(msg string, err error) {
	log.Fatal(msg, err, zap.String("errorDetails", err.Error()), zap.String("eventType", "error occurred"))
}

//InitLogger : function to populate and initialize the Logger structure
// func InitLogger(country string, debugParam bool, version string) Logger {
func InitLogger(country string) Logger {
	Log = Logger{
		Country: country,
		// ShowDebug:  debugParam,
		// AppVersion: version,
	}
	return Log
}
