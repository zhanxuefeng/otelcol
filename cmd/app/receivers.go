package app

import (
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/filelogreceiver"
	"go.opentelemetry.io/collector/receiver"
	"go.opentelemetry.io/collector/receiver/otlpreceiver"

	"go.opentelemetry.io/collector/otelcol"
)

func BuildReceivers(factories *otelcol.Factories) error {
	receivers := []receiver.Factory{
		filelogreceiver.NewFactory(),
		otlpreceiver.NewFactory(),
	}
	items, err := otelcol.MakeFactoryMap(receivers...)
	if err != nil {
		return err
	}
	factories.Receivers = items
	return nil
}
