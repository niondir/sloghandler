# sloghandler
Handler for go 1.21 slog package

Supported Logging libraries:
* [Logrus](https://github.com/sirupsen/logrus)

Contributions for other libraries are welcome.

Please report any issues you find, pull requests for bugfixes and tests are also welcome.

## Usage

```go get github.com/niondir/sloghandler/logrus```

```
import "github.com/niondir/sloghandler/logrus/handler"

slogger := handler.New(handler.NewHandler(logrusLogger))
slogger.Info("hello, Info log")
```
