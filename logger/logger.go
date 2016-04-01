package logger

import (
	log "github.com/cihub/seelog"
)

const (
	FATALLV = iota //0
	ERRORLV
	WARNLV
	INFOLV
	DEBUGLV
	VERBOSELV
)

//配置文件字符串形式
var configstring = `
<seelog minlevel="debug">
    <outputs>
        <filter levels="debug">
            <rollingfile formatid="sended" type="date" filename="/apps/logs/akbs/akbs.log" datepattern="2006-01-02" />
        </filter>
        <filter levels="info">
            <console formatid="info"/>
            <rollingfile formatid="info" type="date" filename="/apps/logs/akbs/akbs.log"  datepattern="2006-01-02" />
        </filter>
        <filter levels="warn,error,critical">
            <console formatid="error"/>
            <rollingfile formatid="error" type="date" filename="/apps/logs/akbs/akbs_errors.log"  datepattern="2006-01-02" />
        </filter>
    </outputs>
    <formats>
        <format id="sended" format="[%Date %Time] [%Level] %Msg%n"/>
        <format id="info" format="[%Date %Time] [%Level] %Msg%n"/>
        <format id="error" format="[%Date %Time] [%Level] %Msg%n"/>
    </formats>
</seelog>
`

var mylog *Logger

type Logger struct {
	LOG log.LoggerInterface
}

func GetLogger() *Logger {
	return mylog
}

func init() {
	pkglogger, err := log.LoggerFromConfigAsString(configstring)
	if err != nil {
		log.Critical("err parsing config log file", err)
		return
	}
	//log.ReplaceLogger(pkglogger)
	mylog = &Logger{
		LOG: pkglogger,
	}
}

func (self *Logger) Infoln(v ...interface{}) {
	self.LOG.Info(v)
}

func (self *Logger) Infof(format string, params ...interface{}) {
	self.LOG.Infof(format, params)
}

func (self *Logger) Errorln(v ...interface{}) {
	self.LOG.Error(v)
}

func (self *Logger) Errorf(format string, params ...interface{}) {
	self.LOG.Errorf(format, params)
}

func (self *Logger) Warnln(v ...interface{}) {
	self.LOG.Warn(v)
}

func (self *Logger) Warnf(format string, params ...interface{}) {
	self.LOG.Warnf(format, params)
}

// -------  log interface ------ //

func (self *Logger) Fatal(v ...interface{}) {
	self.LOG.Error(v...)
}
func (self *Logger) Fatalf(format string, v ...interface{}) {
	self.LOG.Errorf(format, v...)
}
func (self *Logger) Fatalln(v ...interface{}) {
	self.LOG.Error(v...)
}
func (self *Logger) Panic(v ...interface{}) {
	self.LOG.Error(v...)
}
func (self *Logger) Panicf(format string, v ...interface{}) {
	self.LOG.Errorf(format, v...)
}
func (self *Logger) Panicln(v ...interface{}) {
	self.LOG.Error(v...)
}
func (self *Logger) Print(v ...interface{}) {
	self.LOG.Info(v...)
}
func (self *Logger) Printf(format string, v ...interface{}) {
	self.LOG.Infof(format, v...)
}
func (self *Logger) Println(v ...interface{}) {
	self.LOG.Info(v...)
}
