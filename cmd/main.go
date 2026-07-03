package main

import (
	"fmt"

	"github.com/zhanxuefeng/otelcol/cmd/app"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/confmap"
	"go.opentelemetry.io/collector/confmap/provider/envprovider"
	"go.opentelemetry.io/collector/confmap/provider/fileprovider"
	"go.opentelemetry.io/collector/otelcol"
	"go.uber.org/zap"
)

var (
	logger = zap.Must(zap.NewProduction()).With(zap.String("collector", "collector"))
)

func main() {

	info := component.BuildInfo{
		Command:     "otelcorecol",
		Description: "Local OpenTelemetry Collector binary, testing only.",
		Version:     "0.155.0",
	}

	set := otelcol.CollectorSettings{
		BuildInfo: info,
		Factories: app.Components,
		ConfigProviderSettings: otelcol.ConfigProviderSettings{
			ResolverSettings: confmap.ResolverSettings{
				DefaultScheme: "file",
				ProviderFactories: []confmap.ProviderFactory{
					envprovider.NewFactory(),
					fileprovider.NewFactory(),
				},
			},
		},
		ProviderModules: map[string]string{
			envprovider.NewFactory().Create(confmap.ProviderSettings{}).Scheme():  "go.opentelemetry.io/collector/confmap/provider/envprovider v1.61.0",
			fileprovider.NewFactory().Create(confmap.ProviderSettings{}).Scheme(): "go.opentelemetry.io/collector/confmap/provider/fileprovider v1.61.0",
		},
		ConverterModules: []string{},
	}

	cmd := otelcol.NewCommand(set)
	if err := cmd.Execute(); err != nil {
		logger.Fatal("fatal", zap.Error(fmt.Errorf("app server run finished with error: %v", err)))
	}
}
