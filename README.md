# Telemetry Library

This Go library provides a comprehensive solution for handling telemetry
data using the OpenTelemetry Protocol (OTLP). It supports metrics, traces,
and logs, and comes with preconfigured modules for easy integration.

## Features

- **Metrics**: Collect and export metrics data.
- **Traces**: Capture and export trace data.
- **Logs**: Log data using `zerolog`.
- **Preconfigured Modules**:
	- **Metrics Host and Runner Collector**: Simplifies the setup for collecting and exporting metrics.
	- **Logs**: Preconfigured to use `zerolog` for efficient logging.

## Installation

To install the library, run:

```sh
go get go.globalso.dev/x/telemetry
```

## Usage

### Metrics
To use the metrics module, import the package and initialize the metrics collector:

```golang
import "go.globalso.dev/x/telemetry/metricsold"

func main() {
    metricsCollector := metrics.NewCollector()
    metricsCollector.Start()
    // Your application code
    metricsCollector.Stop()
}
```

### Traces 

To use the traces module, import the package and initialize the tracer:

```golang
import "go.globalso.dev/x/telemetry/traces"

func main() {
    tracer := traces.NewTracer()
    tracer.Start()
    // Your application code
    tracer.Stop()
}
```

### Logs

To use the logs module with zerolog, import the package and initialize the logger:

```golang
import (
    "go.globalso.dev/x/telemetry/logger"
    "github.com/rs/zerolog/log"
)

func main() {
    logger := logs.NewLogger()
    log.Info().Msg("Application started")
    // Your application code
    log.Info().Msg("Application stopped")
}
```

## Configuration

The library can be configured using environment variables or configuration
files. Refer to the documentation for detailed configuration options. 

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request. 

## Acknowledgements

- [OpenTelemetry](https://opentelemetry.io/)
- [Zerolog](https://github.com/rs/zerolog)

## Contact
For any questions or support, please open an issue on the GitHub repository.
