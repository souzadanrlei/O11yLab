package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/metric/global"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

type Item struct {
	Name string `json:"name"`
}

func initTracer(ctx context.Context, serviceName string) (*sdktrace.TracerProvider, error) {
	endpoint := os.Getenv("OTEL_EXPORTER_OTLP_ENDPOINT")
	if endpoint == "" {
		endpoint = "otel-gateway-collector-headless.default.svc.cluster.local:4317"
	}

	exporter, err := otlptracegrpc.New(ctx,
		otlptracegrpc.WithEndpoint(endpoint),
		otlptracegrpc.WithInsecure(),
	)
	if err != nil {
		return nil, fmt.Errorf("creating OTLP exporter: %w", err)
	}

	res, err := resource.New(ctx,
		resource.WithAttributes(
			attribute.String("service.name", serviceName),
		),
	)
	if err != nil {
		return nil, fmt.Errorf("creating resource: %w", err)
	}

	bsp := sdktrace.NewBatchSpanProcessor(exporter)
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithResource(res),
		sdktrace.WithSpanProcessor(bsp),
	)
	otel.SetTracerProvider(tp)
	return tp, nil
}

func main() {
	ctx := context.Background()
	serviceName := os.Getenv("OTEL_SERVICE_NAME")
	if serviceName == "" {
		serviceName = "fastapi-go-service"
	}

	tp, err := initTracer(ctx, serviceName)
	if err != nil {
		log.Fatalf("failed to init tracer: %v", err)
	}
	defer func() { _ = tp.Shutdown(ctx) }()

	// Metrics
	meter := global.Meter("fastapi.meter")
	itemCounter, err := meter.Int64Counter("app.items.count")
	if err != nil {
		log.Printf("warning: couldn't create counter: %v", err)
	}

	tr := otel.Tracer("fastapi.tracer")
	r := mux.NewRouter()

	// Root
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ctx, span := tr.Start(r.Context(), "root")
		defer span.End()

		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]string{"message": "API Go com OpenTelemetry"})
	}).Methods(http.MethodGet)

	// CRUD endpoints
	r.HandleFunc("/items", func(w http.ResponseWriter, r *http.Request) {
		ctx, span := tr.Start(r.Context(), "post_create_item")
		defer span.End()

		var it Item
		if err := json.NewDecoder(r.Body).Decode(&it); err != nil {
			http.Error(w, "invalid body", http.StatusBadRequest)
			return
		}

		span.SetAttributes(attribute.String("item.name", it.Name))
		if itemCounter != nil {
			itemCounter.Add(ctx, 1, attribute.Key("operation").String("create"))
		}

		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]interface{}{"status": "create", "item": it.Name})
	}).Methods(http.MethodPost)

	r.HandleFunc("/item/{id:[0-9]+}", func(w http.ResponseWriter, r *http.Request) {
		ctx, span := tr.Start(r.Context(), "get_item")
		defer span.End()

		vars := mux.Vars(r)
		id, _ := strconv.Atoi(vars["id"])
		span.SetAttributes(attribute.Int("item.id", id))
		if itemCounter != nil {
			itemCounter.Add(ctx, 1, attribute.Key("operation").String("read"))
		}

		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]interface{}{"id": id, "name": fmt.Sprintf("item%d", id)})
	}).Methods(http.MethodGet)

	r.HandleFunc("/item/{id:[0-9]+}", func(w http.ResponseWriter, r *http.Request) {
		ctx, span := tr.Start(r.Context(), "update_item")
		defer span.End()

		vars := mux.Vars(r)
		id, _ := strconv.Atoi(vars["id"])
		var it Item
		if err := json.NewDecoder(r.Body).Decode(&it); err != nil {
			http.Error(w, "invalid body", http.StatusBadRequest)
			return
		}

		span.SetAttributes(attribute.Int("item.id", id))
		if itemCounter != nil {
			itemCounter.Add(ctx, 1, attribute.Key("operation").String("update"))
		}

		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]interface{}{"status": "update", "id": id, "name": it.Name})
	}).Methods(http.MethodPut)

	r.HandleFunc("/item/{id:[0-9]+}", func(w http.ResponseWriter, r *http.Request) {
		ctx, span := tr.Start(r.Context(), "delete_item")
		defer span.End()

		id, _ := strconv.Atoi(mux.Vars(r)["id"])
		span.SetAttributes(attribute.Int("item.id", id))
		if itemCounter != nil {
			itemCounter.Add(ctx, 1, attribute.Key("operation").String("delete"))
		}

		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]interface{}{"status": "delete", "id": id})
	}).Methods(http.MethodDelete)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	srv := &http.Server{
		Handler:      r,
		Addr:         ":" + port,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Printf("listening on %s (service=%s)", srv.Addr, serviceName)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("server failed: %v", err)
	}
}
