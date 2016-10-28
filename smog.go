package smog

import (
    "github.com/jcelliott/lumber"
    "SmallCang/module/env"
)


type LoggerInterface interface{
    AddLoggers(...lumber.Logger)
    Trace(string, ...interface{})
    Debug(string, ...interface{})
    Info(string, ...interface{})
    Warn(string, ...interface{})
    Error(string, ...interface{})
    Fatal(string, ...interface{})
}


func NewLogger(filepath string, maxLine int, backups int) LoggerInterface {
    logFile, err := lumber.NewFileLogger(
        //env.Conf.Log.File,
        filepath,
        lumber.INFO,
        lumber.ROTATE,
        maxLine,
        backups,
        //env.Conf.Log.Max_line,
        //env.Conf.Log.Backups,
        256,
    )

    if err != nil {
        panic(err)
    }

    var log LoggerInterface
    if env.Conf.Env.Mode == env.ENV_MODE_PROD {
        log = lumber.NewMultiLogger()
        log.AddLoggers(logFile)
    } else {
        log = &DebugLogger{
            multiLogger: lumber.NewMultiLogger(),
        }
        logConsole := lumber.NewConsoleLogger(lumber.DEBUG)
        log.AddLoggers(logFile, logConsole)
    }
    return log
}



