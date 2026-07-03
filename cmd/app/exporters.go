package app

import (
	"go.opentelemetry.io/collector/exporter"
	"go.opentelemetry.io/collector/exporter/debugexporter"

	"go.opentelemetry.io/collector/otelcol"
)

func BuildExporters(factories *otelcol.Factories) error {
	exporters := []exporter.Factory{
		debugexporter.NewFactory(),
	}
	items, err := otelcol.MakeFactoryMap(exporters...)
	if err != nil {
		return err
	}
	factories.Exporters = items
	return nil
}
