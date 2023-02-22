package logger

import (
	"github.com/sirupsen/logrus"
	"runtime"
	"strconv"
)

func LoggerJson() {
	logrus.SetLevel(logrus.TraceLevel)
	logrus.SetFormatter(&logrus.JSONFormatter{
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			fileName := frame.File + ":" + strconv.Itoa(frame.Line)
			//return frame.Function, fileName
			return "", fileName
		},
	})
	logrus.SetReportCaller(true)
}

func NewLoggerField(p Field) *Field {
	x := Field{
		EventName: p.EventName,
	}

	return &x
}

func (f *Field) Append(eventName string, value string) *Field {
	f.EventName = eventName
	f.Value = value

	return f
}

type Field struct {
	EventName string `json:"event_name,omitempty"`
	Value     string `json:"value,omitempty"`
}
