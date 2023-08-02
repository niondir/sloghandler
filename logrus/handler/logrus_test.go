package handler_test

import (
	"bytes"
	"fmt"
	"github.com/niondir/sloghandler/logrus/handler"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"log/slog"
	"testing"
)

func TestLogrusViaSlog_GroupsAndAttrs(t *testing.T) {
	outBuf := &bytes.Buffer{}
	logrusLogger := logrus.New()
	logrusLogger.Out = outBuf

	logrusLogger.SetFormatter(&logrus.TextFormatter{
		DisableTimestamp: true,
	})
	// 	RFC3339     = "2006-01-02T15:04:05Z07:00"
	//logTime, err  := time.Parse(time.RFC3339, "2020-01-01T00:00:00Z")
	//require.NoError(t, err)

	logrusLogger.SetLevel(logrus.InfoLevel)
	slogger := slog.New(handler.New(logrusLogger))

	slogger.Info("info")
	assert.Equal(t, "level=info msg=info\n", outBuf.String())

	outBuf.Reset()
	slogger.With("isGrouped", false).WithGroup("grouped").With("isGrouped", true).Info("grouped attrs")
	assert.Equal(t, "level=info msg=\"grouped attrs\" grouped:isGrouped=true isGrouped=false\n", outBuf.String())

	// TODO: Test time output
	// but that needs some hack for testing in the handler since slog always takes time.Now() for the record
}

func TestLogrusViaSlog_Levels(t *testing.T) {
	outBuf := &bytes.Buffer{}
	logrusLogger := logrus.New()
	logrusLogger.Out = outBuf

	logrusLogger.SetFormatter(&logrus.TextFormatter{
		DisableTimestamp: true,
	})

	logrusLogger.SetLevel(logrus.InfoLevel)
	slogger := slog.New(handler.New(logrusLogger))

	// Logrus Trace (Log everything)
	logrusLogger.SetLevel(logrus.TraceLevel)

	outBuf.Reset()
	slogger.Debug("debug", "fieldName", "value")
	assert.Equal(t, "level=debug msg=debug fieldName=value\n", outBuf.String())

	outBuf.Reset()
	slogger.Info("info", "fieldName", "value")
	assert.Equal(t, "level=info msg=info fieldName=value\n", outBuf.String())

	outBuf.Reset()
	slogger.Warn("warn", "fieldName", "value")
	assert.Equal(t, "level=warning msg=warn fieldName=value\n", outBuf.String())

	outBuf.Reset()
	slogger.Error("error", "err", fmt.Errorf("an error"))
	assert.Equal(t, "level=error msg=error err=\"an error\"\n", outBuf.String())

	// Logrus Panic (Log nothing)
	logrusLogger.SetLevel(logrus.PanicLevel)

	outBuf.Reset()
	slogger.Debug("debug", "fieldName", "value")
	assert.Equal(t, "", outBuf.String())

	outBuf.Reset()
	slogger.Info("info", "fieldName", "value")
	assert.Equal(t, "", outBuf.String())

	outBuf.Reset()
	slogger.Warn("warn", "fieldName", "value")
	assert.Equal(t, "", outBuf.String())

	outBuf.Reset()
	slogger.Error("error", fmt.Errorf("an error"))
	assert.Equal(t, "", outBuf.String())

	// Logrus Error (Log error but no warning)
	logrusLogger.SetLevel(logrus.ErrorLevel)

	outBuf.Reset()
	slogger.Warn("warn", "fieldName", "value")
	assert.Equal(t, "", outBuf.String())

	outBuf.Reset()
	slogger.Error("error", "err", fmt.Errorf("an error"))
	assert.Equal(t, "level=error msg=error err=\"an error\"\n", outBuf.String())
}
