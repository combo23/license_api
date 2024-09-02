package types

import "go.mongodb.org/mongo-driver/bson/primitive"

type License struct {
	LicenseKey string             `bson:"license" json:"license"`
	Username   string             `bson:"username" json:"username"`
	CreatedAt  primitive.DateTime `bson:"created_at" json:"created_at"`
	ExpiresAt  primitive.DateTime `bson:"expires_at" json:"expires_at"`
	UpdatedAt  primitive.DateTime `bson:"updated_at" json:"updated_at"`
	HWID       string             `bson:"hwid" json:"hwid"`
	Status     string             `bson:"status" json:"status"`
}

type CreateLicense struct {
	Username  string `json:"username"`
	ExpiresAt int64  `json:"expires_at"`
}

type LicenseVerify struct {
	LicenseKey string `json:"license_key"`
	HWID       string `json:"hwid"`
}
