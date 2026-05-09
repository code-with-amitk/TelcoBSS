package main

import (
	"context"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	"telcobss/internal/common/db"
	"telcobss/internal/common/kafka"
	"telcobss/internal/common/metrics"
	"telcobss/internal/billing/handlers"
	"telcobss/internal/billing/service"
)

func main() {
	logrus.Info("starting billing-service")

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	mysqlDSN := envOrDefault("MYSQL_DSN", "root:password@tcp(localhost:3306)/telco_bss")
	kafkaBrokers := envOrDefault("KAFKA_BROKERS", "localhost:9092")

	sqlDB, err := db.NewMySQLClient(mysqlDSN)
	if err != nil {
		logrus.Fatalf("failed to connect MySQL: %v", err)
	}

	producer, err := kafka.NewProducer([]string{kafkaBrokers})
	if err != nil {
		logrus.Fatalf("failed to create kafka producer: %v", err)
	}

	metrics.InitPrometheus()

	repo := service.NewBillingRepository(sqlDB)
	billingService := service.NewBillingService(repo, producer)
	handler := handlers.NewBillingHandler(billingService)

	router := gin.New()
	router.Use(gin.Recovery())
	router.GET("/v1/invoices/:id", handler.GetInvoice)
	router.GET("/metrics", gin.WrapH(promhttp.Handler()))

	go func() {
		if err := router.Run(":8083"); err != nil && err != http.ErrServerClosed {
			logrus.Fatalf("failed to run HTTP server: %v", err)
		}
	}()

	grpcListener, err := net.Listen("tcp", ":9093")
	if err != nil {
		logrus.Fatalf("failed to listen grpc: %v", err)
	}

	grpcServer := grpc.NewServer()
	// TODO: Register gRPC services after proto generation.

	go func() {
		logrus.Infof("starting grpc server on :9093")
		if err := grpcServer.Serve(grpcListener); err != nil {
			logrus.Fatalf("grpc server failed: %v", err)
		}
	}()

	<-ctx.Done()
	logrus.Info("shutting down billing-service")
	grpcServer.GracefulStop()
}

func envOrDefault(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}
