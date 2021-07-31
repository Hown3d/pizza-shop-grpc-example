package bakery

import (
	"log"
	"time"

	"github.com/google/uuid"
	bakerypb "github.com/hown3d/pizza-shop-grpc-example/bakery/proto/bakery/v1"
	"google.golang.org/grpc/grpclog"
)

type Server struct {
	log grpclog.LoggerV2
}

func (s *Server) BakePizza(req *bakerypb.BakePizzaRequest, stream bakerypb.BakeryService_BakePizzaServer) error {
	uuid := uuid.Must(uuid.NewRandom())
	var counter int
	//Sending Preparing Response to Stream
	for {
		switch counter {
		case 0:
			log.Printf("Preparing Pizza with id %v", uuid.String())
			err := stream.Send(&bakerypb.BakePizzaResponse{Id: uuid.String(), Status: bakerypb.BakePizzaResponse_BAKE_STATUS_PREPARING})
			if err != nil {
				return err
			}
		case 1:
			err := stream.Send(&bakerypb.BakePizzaResponse{Id: uuid.String(), Status: bakerypb.BakePizzaResponse_BAKE_STATUS_BAKING})
			if err != nil {
				return err
			}
		case 2:
			err := stream.Send(&bakerypb.BakePizzaResponse{Id: uuid.String(), Status: bakerypb.BakePizzaResponse_BAKE_STATUS_READY})
			if err != nil {
				return err
			}
		default:
			return nil
		}
		counter++
		time.Sleep(time.Second * 2)
	}
}
