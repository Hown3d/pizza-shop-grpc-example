package main

import (
	"errors"
	"flag"
	"fmt"
	"net"
	"net/url"
	"os"
	"time"

	bakerypb "github.com/hown3d/pizza-shop-grpc-example/bakery/proto/bakery/v1"
	"github.com/hown3d/pizza-shop-grpc-example/bakery/service"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

type pgURL url.URL

func (p *pgURL) Set(in string) error {
	u, err := url.Parse(in)
	if err != nil {
		return err
	}

	switch u.Scheme {
	case "psql", "postgresql":
	default:
		return errors.New("unexpected scheme in URL")
	}

	*p = pgURL(*u)
	return nil
}

func (p pgURL) String() string {
	return (*url.URL)(&p).String()
}

var (
	port = flag.Int("port", 10000, "The server port")
	u    pgURL
)

func main() {
	flag.Var(&u, "postgres-url", "URL formatted address of the postgres server to connect to")
	flag.Parse()

	log := logrus.New()

	log.Formatter = &logrus.TextFormatter{
		TimestampFormat: time.StampMilli,
		FullTimestamp:   true,
		ForceColors:     true,
	}

	if u.String() == "" {
		log.Fatal("Flag postgres-url is required")
	}

	service, err := service.NewService(log, (*url.URL)(&u))
	if err != nil {
		log.Fatalf("Can't create bakery Service: %v", err)
	}

	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%v", *port))
	if err != nil {
		log.Fatalf("Failed to listen on port %v %v", port, err)
	}

	grpcServer := grpc.NewServer()

	bakerypb.RegisterBakeryServiceServer(grpcServer, service)
	grpc_health_v1.RegisterHealthServer(grpcServer, service.Health)
	reflection.Register(grpcServer)

	log.Infof("Starting grpc server on %v ...", lis.Addr().String())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to start grpc Server %v", err)
	}
}

func mustGetEnv(key string) (string, error) {
	env := os.Getenv(key)
	if env == "" {
		return "", errors.New(fmt.Sprintf("Can't get environment variable %v", key))
	}
	return env, nil
}
