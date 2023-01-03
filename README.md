# go-health-http
Health check implementation for HTTP servers - To be used with [go-health](https://github.com/pcordeiro/go-health)

#### Usage
Get the package
```bash
go get -u github.com/pcordeiro/go-health-http
```

In the code:
```go
import(
   	"github.com/pcordeiro/go-health"
	health_http "github.com/pcordeiro/go-health-http"
)

health, err := health.NewHealth(
    health.WithComponent(
        health.Component{
            Name:    app.config.Name,
            Version: app.config.Version,
        },
    ),
    health.WithChecks(
        health.Check{
            Name:      "Google",
            Timeout:   2 * time.Second,
            SkipOnErr: false,
            Check: healthhttp.NewHttpCheck(&healthhttp.Config{
                Name:    "Google",
                URL:     "https://google.com",
                Timeout: 2 * time.Second,
            }),
        },
    ),
)
``` 