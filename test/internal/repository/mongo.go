package repository

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

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

func (r *ClickRepository) Update(bannerID string, click domain.Click) error {
	_, err := r.collection.UpdateOne(context.TODO(),
		bson.M{"_id": bannerID},
		bson.M{"$push": bson.M{"clicks": click}})

	if err != nil {
		return err
	}

	return nil
}

func (r *ClickRepository) GetStats(bannerID string, tsFrom, tsTo time.Time) ([]domain.Stat, error) {
	pipeline := mongo.Pipeline{
		{{
			"$match", bson.M{
				"_id": bannerID,
				"clicks.created_at": bson.M{
					"$gte": tsFrom,
					"$lt":  tsTo,
				},
			},
		}},
		{{
			"$unwind", "$clicks",
		}},
		{{
			"$match", bson.M{
				"clicks.created_at": bson.M{
					"$gte": tsFrom,
					"$lt":  tsTo,
				},
			},
		}},
		{{
			"$group", bson.M{
				"_id": bson.M{
					"$dateTrunc": bson.M{
						"date": "$clicks.created_at",
						"unit": "minute",
					},
				},
				"count": bson.M{"$sum": 1},
			},
		}},
		{{
			"$sort", bson.M{"_id": 1},
		}},
	}
	cursor, err := r.collection.Aggregate(context.TODO(), pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var result struct {
		Interval time.Time `bson:"_id"`
		Count    int       `bson:"count"`
	}

	var stats []domain.Stat

	for cursor.Next(context.TODO()) {

		if err := cursor.Decode(&result); err != nil {
			continue
		}
		ts := primitive.NewDateTimeFromTime(result.Interval)
		stats = append(stats, domain.Stat{
			Ts: ts,
			V:  result.Count,
		})
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
