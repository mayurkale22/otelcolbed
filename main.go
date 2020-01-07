package main

import (
	"log"
	
	"github.com/open-telemetry/opentelemetry-collector/config"
  	"github.com/open-telemetry/opentelemetry-collector/exporter"
  	"github.com/open-telemetry/opentelemetry-collector/extension"
  	"github.com/open-telemetry/opentelemetry-collector/internal/version"
  	"github.com/open-telemetry/opentelemetry-collector/processor"
  	"github.com/open-telemetry/opentelemetry-collector/receiver"
  	"github.com/open-telemetry/opentelemetry-collector/service"
 	"github.com/open-telemetry/opentelemetry-collector/receiver/opencensusreceiver"
 	"github.com/open-telemetry/opentelemetry-collector/extension/pprofextension"
  	"github.com/open-telemetry/opentelemetry-collector/processor/queuedprocessor"
  	"github.com/open-telemetry/opentelemetry-collector/processor/batchprocessor"
  	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/stackdriverexporter"
)

func main() {
	handleErr := func(err error) {
		if err != nil {
			log.Fatalf("Failed to run the service: %v", err)
		}
	}

	extensions, err := extension.Build(
		&pprofextension.Factory{},
	)
	handleErr(err)

	receivers, err := receiver.Build(
		&opencensusreceiver.Factory{},
	)
	handleErr(err)

	exporters, err := exporter.Build(
		&stackdriverexporter.Factory{},
	)
	handleErr(err)

	processors, err := processor.Build(
		&queuedprocessor.Factory{},
		&batchprocessor.Factory{},
	)
	handleErr(err)

	factories := config.Factories{
		Extensions: extensions,
		Receivers:  receivers,
		Processors: processors,
		Exporters:  exporters,
	}

	info := service.ApplicationStartInfo{
		ExeName:  "otelcol",
		LongName: "OpenTelemetry Collector",
		Version:  version.Version,
		GitHash:  version.GitHash,
	}

	svc, err := service.New(factories, info)
	handleErr(err)

	err = svc.Start()
	handleErr(err)
}

// package main

// import (
// 	"log"

// 	"github.com/open-telemetry/opentelemetry-collector/defaults"
// 	"github.com/open-telemetry/opentelemetry-collector/internal/version"
// 	"github.com/open-telemetry/opentelemetry-collector/service"
// )

// func main() {
// 	handleErr := func(err error) {
// 		if err != nil {
// 			log.Fatalf("Failed to run the service: %v", err)
// 		}
// 	}

// 	factories, err := defaults.Components()
// 	handleErr(err)

// 	info := service.ApplicationStartInfo{
// 		ExeName:  "otelcol",
// 		LongName: "OpenTelemetry Collector",
// 		Version:  version.Version,
// 		GitHash:  version.GitHash,
// 	}

// 	svc, err := service.New(factories, info)
// 	handleErr(err)

// 	err = svc.Start()
// 	handleErr(err)
// }
