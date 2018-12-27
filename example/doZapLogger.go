package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"pdriver/zapLogger"
)

type Configr struct {
	Path 	string
}

var Logger *zap.Logger

func main() {

	var c = &zapLogger.Configuration{
		LogPath: "./zap_output.log",
	}

	logger := c.InitLogger2("info")
	defer logger.Sync()


	infoMsg := "redis connected"
	errMsg := "connect refused"

	logger.Error("error", zap.String("detail", infoMsg))
	logger.Info("info", zap.String("detail", infoMsg))
	logger.Error("redis error", zap.String("detail", errMsg))
}

func (c *Configr)InitLogger() *zap.Logger {
	path := c.Path
	level := "DEBUG"

	isDebug := true

	logger := initLogger(path, level, isDebug)

	log.SetFlags(log.Lmicroseconds)

	return logger
}

func initLogger(path, level string, isDebug bool) *zap.Logger {
	//var jsonformat string
	//if isDebug {
	//	jsonformat = fmt.Sprintf(`{
	//		"level": "%s",
	//		"encoding": "json",
    //  		"outputPaths": ["stdout"],
    //  		"errorOutputPaths": ["stdout"]
	//	}`, level)
	//} else {
	//	jsonformat = fmt.Sprintf(`{
	//		"level": "%s",
	//		"encoding": "json",
    //  		"outputPaths": ["%s"],
    //  		"errorOutputPaths": ["%s"]
	//	}`, level, path, path)
	//}

	var logConf  = zap.Config {
		//Level: nil,
		Encoding: 	"json",
		OutputPaths: 	[]string{path},
		ErrorOutputPaths: []string{path},
	}

	//if err := json.Unmarshal([]byte(jsonformat), &logConf); err != nil {
	//	panic(err)
	//}

	logConf.EncoderConfig = zap.NewProductionEncoderConfig()
	logConf.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	var err error
	logger, err := logConf.Build()
	if err != nil {
		panic(err)
	}

	return logger
}
