package api

import (
	"context"
	"fmt"

	otel "go.opentelemetry.io/otel"
)

func PrintThing(message string) {
	_, span := otel.GetTracerProvider().Tracer("tracer-name").Start(context.Background(), "span-name")
	fmt.Println(message)
	span.End()
}
