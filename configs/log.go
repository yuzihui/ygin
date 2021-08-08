package configs

import (
	"ecloudsystem/pkg/logger"
	"fmt"
	"go.uber.org/zap"
	"time"
)

var Logger *zap.Logger

func InitLogs(){

	config := Get().App
	ProjectLogFile := fmt.Sprintf("./logs/%s-ecs.log", time.Now().Format("2006-01-02"))

	// 初始化 logger
	loggers, err := logger.NewJSONLogger(
		logger.WithDisableConsole(),
		logger.WithField("domain", fmt.Sprintf("%s[%s]", config.Name, config.Env)),
		logger.WithTimeLayout("2006-01-02 15:04:05"),
		logger.WithFileP(ProjectLogFile),
	)
	
	if err != nil {
		panic(err)
	}
	Logger = loggers
}
