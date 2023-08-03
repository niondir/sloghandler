# sloghandler
Handler for go 1.21 slog package

## Use cases
**Migrate to slog**  
Many projects use a 3rd party structured logging framework like logrus. The slog handler for logrus helps to introduce slog while keeping the logrus logging configuration as it is. One can register the logrus handler and just use both APIs. Once the project is only using the slog api anymore, it's possible to use different slog handlers as well.

**Use 3rd party library that supports slog**  
When a 3rd party library accepts a slog handler one can reuse the existing logger (e.g. logrus) from the main application.

## Supported Logging libraries
* [logrus](https://github.com/sirupsen/logrus)

Candidates:
* [zerolog](https://github.com/rs/zerolog)
* [zap](https://github.com/uber-go/zap)
* [go-kit](https://github.com/go-kit/log)
* [apex/log](https://github.com/apex/log)
* [log15](https://github.com/inconshreveable/log15)

Contributions for other libraries are welcome.

Please report any issues you find, pull requests for bugfixes and tests are also welcome.

## Usage

```go get github.com/niondir/sloghandler/logrus```

```
import "github.com/niondir/sloghandler/logrus/handler"

logrusLogger := logrus.StandardLogger()
slogger := slog.New(handler.New(logrusLogger))
slogger.Info("hello, Info log")
```
