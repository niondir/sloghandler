# sloghandler
Handler for go 1.21 slog package

Supported Logging libraries:
* [Logrus](https://github.com/sirupsen/logrus)

## Usage

```go get github.com/niondir/sloghandler/logrus```

```
slogger := slog.New(logrusHandler.NewHandler(logrusLogger))
slogger.Info("hello, Info log")
```
