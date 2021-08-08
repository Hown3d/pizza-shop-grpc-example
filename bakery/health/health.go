package health

import (
	"context"

	"github.com/hown3d/pizza-shop-grpc-example/bakery/storage"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/health/grpc_health_v1"
)

type HealthChecker struct {
	repo *storage.Repository
}

//Checks the Database for access and response SERVING when connectable
// and NOT_SERVING when no connection can be made
func (s *HealthChecker) Check(ctx context.Context, req *grpc_health_v1.HealthCheckRequest) (*grpc_health_v1.HealthCheckResponse, error) {
	logrus.Info("Serving the Check request for health check")

	response := &grpc_health_v1.HealthCheckResponse{
		Status: grpc_health_v1.HealthCheckResponse_NOT_SERVING,
	}
	// check for connectivity, respond with serving if alive
	if err := s.repo.Ping(); err == nil {
		response.Status = grpc_health_v1.HealthCheckResponse_SERVING
	}

	return response, nil

}

func (s *HealthChecker) Watch(req *grpc_health_v1.HealthCheckRequest, server grpc_health_v1.Health_WatchServer) error {
	logrus.Info("Serving the Watch request for health check")
	return server.Send(&grpc_health_v1.HealthCheckResponse{
		Status: grpc_health_v1.HealthCheckResponse_SERVING,
	})
}

func NewHealthChecker(repo *storage.Repository) *HealthChecker {
	return &HealthChecker{
		repo: repo,
	}
}
