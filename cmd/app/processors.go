package app

import (
	"go.opentelemetry.io/collector/otelcol"
	"go.opentelemetry.io/collector/processor"
)

func BuildProcessors(c *otelcol.Factories) error {
	processors := []processor.Factory{}
	// 填充factories
	items, err := otelcol.MakeFactoryMap(processors...)
	if err != nil {
		return err
	}
	c.Processors = items
	return nil
}
