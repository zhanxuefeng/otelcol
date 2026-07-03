package app

import (
	"go.opentelemetry.io/collector/otelcol"
	"go.opentelemetry.io/collector/service/telemetry/otelconftelemetry"
)

func Components() (otelcol.Factories, error) {
	factories := otelcol.Factories{
		Telemetry: otelconftelemetry.NewFactory(),
	}

	if err := BuildExtensions(&factories); err != nil {
		return otelcol.Factories{}, err
	}

	if err := BuildReceivers(&factories); err != nil {
		return otelcol.Factories{}, err
	}

	if err := BuildProcessors(&factories); err != nil {
		return otelcol.Factories{}, err
	}

	if err := BuildExporters(&factories); err != nil {
		return otelcol.Factories{}, err
	}
	return factories, nil
}
