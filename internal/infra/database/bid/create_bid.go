package bid

import (
	"context"
	"sync"

	"github.com/dpcamargo/fullcycle-auction/configuration/logger"
	"github.com/dpcamargo/fullcycle-auction/internal/entity/auction_entity"
	"github.com/dpcamargo/fullcycle-auction/internal/entity/bid_entity"
	"github.com/dpcamargo/fullcycle-auction/internal/infra/database/auction"
	"github.com/dpcamargo/fullcycle-auction/internal/internal_error"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

type BidEntityMongo struct {
	ID        string  `bson:"_id"`
	UserID    string  `bson:"user_id"`
	AuctionID string  `bson:"auction_id"`
	Amount    float64 `bson:"amount"`
	Timestamp int64   `bson:"timestamp"`
}

type BidRepository struct {
	Collection        *mongo.Collection
	AuctionRepository *auction.AuctionRepository
}

func (br *BidRepository) CreateBid(
	ctx context.Context,
	bidEntities []bid_entity.Bid) *internal_error.InternalError {
	var wg sync.WaitGroup

	for _, bid := range bidEntities {
		wg.Add(1)

		go func(bidValue bid_entity.Bid) {
			defer wg.Done()

			auctionEntity, err := br.AuctionRepository.FindAuctionByID(ctx, bidValue.AuctionID)
			if err != nil {
				logger.Error("Error trying to find auction by id", err)
				return
			}

			if auctionEntity.Status != auction_entity.Active {
				return
			}

			bidEntityMongo := &BidEntityMongo{
				ID:        bidValue.ID,
				UserID:    bidValue.UserID,
				AuctionID: bidValue.AuctionID,
				Amount:    bidValue.Amount,
				Timestamp: bidValue.Timestamp.Unix(),
			}
			if _, err := br.Collection.InsertOne(ctx, bidEntityMongo); err != nil {
				logger.Error("Error trying to insert bid", err, zap.Any("bid", bidEntityMongo))
				return
			}

		}(bid)
	}

	wg.Wait()
	return nil
}
