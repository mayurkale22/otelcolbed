// Copyright 2020, OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Program otelcol is the OpenTelemetry Collector that collects stats
// and traces from OpenCensus Receiver and exports to a Stackdriver backend.
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
			log.Fatalf("Failed to run the collector: %v", err)
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
                &batchprocessor.Factory{},
		&queuedprocessor.Factory{},
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

