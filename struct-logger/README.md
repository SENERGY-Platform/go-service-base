go-service-base/structured-logger
=======

![Image](https://img.shields.io/github/v/tag/SENERGY-Platform/go-service-base?filter=struct-logger%2A&label=latest)

### Levels

    "debug"
	"info"
	"warn"
	"error"

### Handlers

    "text"
    "json"
    "discard"

### Example

	var Logger *slog.Logger

	useUTC, _ := strconv.ParseBool(os.Getenv("LOG_TIME_UTC"))

	recordTime := NewRecordTime(os.Getenv("LOG_TIME_FORMAT"), useUTC)

	options := &slog.HandlerOptions{
		AddSource:   false,
		Level:       GetLevel(os.Getenv("LOG_LEVEL"), slog.LevelWarn),
		ReplaceAttr: recordTime.ReplaceAttr,
	}

	handler := GetHandler(os.Getenv("LOG_HANDLER"), os.Stdout, options, slog.Default().Handler())
	handler = handler.WithAttrs([]slog.Attr{
        slog.String(attribute.ProjectAttrKey, "my-project"),
	})

	Logger = slog.New(handler)

	Logger.Info("hello", slog.String("user", os.Getenv("USER")))