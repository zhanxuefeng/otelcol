package app

import (
	"go.opentelemetry.io/collector/extension"

	"go.opentelemetry.io/collector/otelcol"
)

func BuildExtensions(factories *otelcol.Factories) error {
	// extensions 数组
	extensions := []extension.Factory{}
	items, err := otelcol.MakeFactoryMap(extensions...)
	if err != nil {
		return err
	}
	factories.Extensions = items
	return nil
}
