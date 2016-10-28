package smog

import (
    "github.com/jcelliott/lumber"
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


func NewLogger(filepath string, maxLine int, backups int, isDebug bool) LoggerInterface {
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
    if isDebug {
        log = &DebugLogger{
            multiLogger: lumber.NewMultiLogger(),
        }
        logConsole := lumber.NewConsoleLogger(lumber.DEBUG)
        log.AddLoggers(logFile, logConsole)
    } else {
        log = lumber.NewMultiLogger()
        log.AddLoggers(logFile)
    }
    return log
}



