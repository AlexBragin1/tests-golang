package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/protobuf/types/known/timestamppb"

	"test/internal/domain"
)

type ClickRepository struct {
	collection *mongo.Collection
}

func NewClickRepository(db *mongo.Database) *ClickRepository {
	return &ClickRepository{
		collection: db.Collection("banner"),
	}
}

func (r *ClickRepository) Update(bannerID string) error {
	_, err := r.collection.UpdateOne(context.TODO(),
		bson.M{"_id": bannerID},
		bson.M{"$push": bson.M{"clicks": domain.Click{CreatedAt: *timestamppb.Now()}}})

	return err
}

func (r *ClickRepository) GetStats(bannerID string, tsFrom, tsTo timestamppb.Timestamp) ([]domain.Click, error) {
	var stats []domain.Click

	cursor, err := r.collection.Find(context.Background(),
		bson.M{"banner_id": bannerID, "create_at": bson.M{"$gte": tsFrom, "$lte": tsTo}})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var stat domain.Click

		if err = cursor.Decode(&stat); err != nil {
			return nil, err
		}
		stats = append(stats, stat)
	}

	return stats, nil
}

func (r *ClickRepository) Save(banner domain.Banner) error {
	_, err := r.collection.InsertOne(context.TODO(), banner)
	if err != nil {
		return err
	}

	return nil
}
