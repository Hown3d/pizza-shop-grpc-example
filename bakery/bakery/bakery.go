package bakery

import (
	"time"

	"github.com/google/uuid"
	bakerypb "github.com/hown3d/pizza-shop-grpc-example/bakery/proto/bakery/v1"
)

type Server struct {
	*bakerypb.UnimplementedBakeryServiceServer
}

func (s *Server) BakePizza(req *bakerypb.BakePizzaRequest, stream bakerypb.BakeryService_BakePizzaServer) error {
	uuid := uuid.Must(uuid.NewRandom())
	//Sending Preparing Response to Stream
	for {
		err := stream.Send(&bakerypb.BakePizzaResponse{Id: uuid.String(), Status: bakerypb.BakePizzaResponse_BAKE_STATUS_PREPARING})
		if err != nil {
			return err
		}
		time.Sleep(time.Second * 3)
	}
}
