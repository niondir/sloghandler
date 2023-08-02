# sloghandler
Handler for go 1.21 slog package

Supported Logging libraries:
* [Logrus](https://github.com/sirupsen/logrus)

## Usage

```go get github.com/niondir/sloghandler/logrus```

```
import "github.com/niondir/sloghandler/logrus/handler"

slogger := slog.New(handler.NewHandler(logrusLogger))
slogger.Info("hello, Info log")
```
