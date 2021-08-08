package service

import (
	"log"
	"net/url"
	"time"

	"github.com/hown3d/pizza-shop-grpc-example/bakery/health"
	bakerypb "github.com/hown3d/pizza-shop-grpc-example/bakery/proto/bakery/v1"
	"github.com/hown3d/pizza-shop-grpc-example/bakery/storage"
	"github.com/sirupsen/logrus"
)

type Service struct {
	Health *health.HealthChecker
	log    *logrus.Logger
	repo   *storage.Repository
}

// NewService creates a new Service, connecting it to the postgres server on
// the URL provided.
func NewService(logger *logrus.Logger, pgURL *url.URL) (*Service, error) {
	repo, err := storage.NewRepository(logger, pgURL)
	if err != nil {
		return nil, err
	}
	return &Service{
		Health: health.NewHealthChecker(repo),
		log:    logger,
		repo:   repo,
	}, nil
}

func (s *Service) BakePizza(req *bakerypb.BakePizzaRequest, stream bakerypb.BakeryService_BakePizzaServer) error {
	var counter int
	var bake *storage.Bake
	var err error
	var bakeID string
	//Sending Preparing Response to Stream
	for {

		switch counter {
		case 0:
			// Create new Bake in Database
			bake, err = s.repo.CreateBake(stream.Context(), req.Name)
			if err != nil {
				return err
			}
			err = (*bake).ID.AssignTo(&bakeID)
			if err != nil {
				return err
			}
			log.Printf("Preparing Pizza with id %v", bakeID)

			err = stream.Send(&bakerypb.BakePizzaResponse{
				Id:     bakeID,
				Status: bakerypb.BakePizzaResponse_BAKE_STATUS_PREPARING,
			},
			)
			if err != nil {
				return err
			}
		case 1:
			// Update bake in database
			log.Printf("Updating Pizza with id %v", bakeID)

			// postgresBakeID, err := storage.IdProtoToPostgres(bakeID)
			// if err != nil {
			// 	return err
			// }
			err = s.repo.UpdateBake(stream.Context(), bake.ID, storage.StatusEnumBaking)
			if err != nil {
				return err
			}
			err = stream.Send(&bakerypb.BakePizzaResponse{
				Id:     bakeID,
				Status: bakerypb.BakePizzaResponse_BAKE_STATUS_BAKING,
			},
			)
			if err != nil {
				return err
			}
		case 2:
			postgresBakeID, err := storage.IdProtoToPostgres(bakeID)
			if err != nil {
				return err
			}
			// Update bake in database, pizza is done baking

			err = s.repo.UpdateBake(stream.Context(), postgresBakeID, storage.StatusEnumDone)
			if err != nil {
				return err
			}
			err = stream.Send(&bakerypb.BakePizzaResponse{
				Id:     bakeID,
				Status: bakerypb.BakePizzaResponse_BAKE_STATUS_READY,
			},
			)
			if err != nil {
				return err
			}
		default:
			return nil
		}
		counter++
		time.Sleep(time.Second * 5)
	}
}
