package trace

import (
	"context"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.7.0"
	"go.opentelemetry.io/otel/trace"
)

type ProviderConfig struct {
	JaegerEndpoint string
	ServiceName    string
	ServiceVersion string
	Environment    string
	Disabled       bool
}

type Provider struct {
	provider trace.TracerProvider
}

func NewProvider(config ProviderConfig) (Provider, error) {
	if config.Disabled {
		return Provider{provider: trace.NewNoopTracerProvider()}, nil
	}

	exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(config.JaegerEndpoint)))
	if err != nil {
		return Provider{}, nil
	}

	p := sdk.NewTracerProvider(
		sdk.WithBatcher(exp),
		sdk.WithResource(
			resource.NewWithAttributes(
				semconv.SchemaURL,
				semconv.ServiceNameKey.String(config.ServiceName),
				semconv.ServiceVersionKey.String(config.ServiceVersion),
				attribute.String("environment", config.Environment),
			),
		),
	)
	otel.SetTracerProvider(p)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{},
		propagation.Baggage{}),
	)
	return Provider{provider: p}, nil
}

func (p Provider) Close(ctx context.Context) error {
	if prov, ok := p.provider.(*sdk.TracerProvider); ok {
		return prov.Shutdown(ctx)
	}
	return nil
}
