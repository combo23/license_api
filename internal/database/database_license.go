package database

import (
	"context"
	"license-api/internal/types"
	"log"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	DB_NAME       = "License-API"
	DB_COLLECTION = "licenses"
)

func (s *service) CreateLicense(license *types.CreateLicense) (*types.License, error) {
	timestamp := time.Now()
	dblicense := new(types.License)
	dblicense.Username = license.Username
	dblicense.ExpiresAt = primitive.NewDateTimeFromTime(time.Unix(license.ExpiresAt, 0))
	dblicense.UpdatedAt = primitive.NewDateTimeFromTime(timestamp)
	dblicense.CreatedAt = primitive.NewDateTimeFromTime(timestamp)
	dblicense.LicenseKey = uuid.NewString()
	dblicense.Status = "active"

	_, err := s.db.Database(DB_NAME).Collection(DB_COLLECTION).InsertOne(context.Background(), dblicense)
	if err != nil {
		return nil, err
	}

	return dblicense, nil
}

func (s *service) GetLicense(key string) (*types.License, error) {
	filter := bson.M{"license": key}

	var license *types.License
	result := s.db.Database(DB_NAME).Collection(DB_COLLECTION).FindOne(context.Background(), filter)
	if result.Err() != nil {
		return nil, result.Err()
	}

	err := result.Decode(&license)
	if err != nil {
		return nil, err
	}

	return license, nil
}

func (s *service) BindLicense(key, hwid string) error {
	filter := bson.M{"license": key}
	update := bson.M{"$set": bson.M{"hwid": hwid}}

	_, err := s.db.Database(DB_NAME).Collection(DB_COLLECTION).UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) UnbindLicense(key string) error {
	filter := bson.M{"license": key}
	update := bson.M{"$set": bson.M{"hwid": "", "updated_at": primitive.NewDateTimeFromTime(time.Now())}}

	_, err := s.db.Database(DB_NAME).Collection(DB_COLLECTION).UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) BanLicense(key string) error {
	filter := bson.M{"license": key}
	update := bson.M{"$set": bson.M{"status": "banned", "updated_at": primitive.NewDateTimeFromTime(time.Now())}}

	_, err := s.db.Database(DB_NAME).Collection(DB_COLLECTION).UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) UpdateLicense(license types.License) error {
	filter := bson.M{"license": license.LicenseKey}
	update := bson.M{"$set": bson.M{"username": license.Username, "expires_at": license.ExpiresAt, "hwid": license.HWID, "status": license.Status, "updated_at": primitive.NewDateTimeFromTime(time.Now())}}

	_, err := s.db.Database(DB_NAME).Collection(DB_COLLECTION).UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) setupDB() {
	result, err := s.db.ListDatabases(context.Background(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	for _, db := range result.Databases {
		if db.Name == DB_NAME {
			return
		}
	}

	err = s.db.Database(DB_NAME).CreateCollection(context.Background(), DB_COLLECTION)
	if err != nil {
		log.Fatal(err)
	}

}
