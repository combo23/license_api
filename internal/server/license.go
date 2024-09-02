package server

import (
	"errors"
	"license-api/internal/logger"
	"license-api/internal/types"
	"license-api/internal/utils"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// getLicense returns the types.License object from the database
// admin auth required
func (s *Server) getLicense(c *gin.Context) {
	logger.LogDebug(s.logger, c, "Getting license")

	license_key := c.Param("id")

	isvalid := utils.IsUUID(license_key)
	if !isvalid {
		s.HandleError(c, InvalidLicense("not matching UUID format"))
		return
	}

	license, err := s.db.GetLicense(license_key)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			s.HandleError(c, InvalidLicense("license not found"))
			return
		}

		s.HandleError(c, InternalServerError("cannot get license: "+err.Error()))
		return
	}

	c.JSON(200, gin.H{"license": license})
}

// createLicense creates a new license in the database
// admin auth required
func (s *Server) createLicense(c *gin.Context) {
	logger.LogDebug(s.logger, c, "Creating license")

	var payload *types.CreateLicense

	err := c.BindJSON(&payload)
	if err != nil {
		s.HandleError(c, InternalServerError("cannot bind payload: "+err.Error()))
		return
	}

	isvalidpayload := utils.IsValidPayload(payload)
	if !isvalidpayload {
		s.HandleError(c, InvalidPayload("missing required fields"))
		return
	}

	license, err := s.db.CreateLicense(payload)
	if err != nil {
		s.HandleError(c, InternalServerError("cannot create license: "+err.Error()))
		return
	}

	c.JSON(200, gin.H{"message": "license created", "license": license})
}

// verifyLicense verifies the login attempt with the license key and HWID
func (s *Server) verifyLicense(c *gin.Context) {
	logger.LogDebug(s.logger, c, "Verifying license")

	var payload types.LicenseVerify
	err := c.BindJSON(&payload)
	if err != nil {
		s.HandleError(c, InternalServerError("cannot bind payload: "+err.Error()))
		return
	}

	isvalidpayload := utils.IsValidPayload(&payload)
	if !isvalidpayload {
		s.HandleError(c, InvalidPayload("missing required fields"))
		return
	}

	license, err := s.db.GetLicense(payload.LicenseKey)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			s.HandleError(c, InvalidLicense("license not found"))
			return
		}

		s.HandleError(c, InternalServerError("cannot get license: "+err.Error()))
		return
	}

	if license.Status != "active" {
		s.HandleError(c, &AppError{StatusCode: 400, Message: "License is not active", LogMessage: "License is not active"})
		return
	}

	if license.ExpiresAt < primitive.NewDateTimeFromTime(time.Now()) {
		s.HandleError(c, &AppError{StatusCode: 400, Message: "License is expired", LogMessage: "License is expired"})
		return
	}

	if license.HWID == "" {
		err = s.db.BindLicense(payload.LicenseKey, payload.HWID)
		if err != nil {
			s.HandleError(c, InternalServerError("cannot bind license: "+err.Error()))
			return
		}

		c.JSON(200, gin.H{"message": "license bound to this machine"})
		return
	}

	if license.HWID != payload.HWID {
		s.HandleError(c, &AppError{StatusCode: 400, Message: "HWID mismatch", LogMessage: "HWID mismatch"})
		return
	}

	c.JSON(200, gin.H{"message": "license verified"})
}

// unbindLicense unbinds the license from the machine
func (s *Server) unbindLicense(c *gin.Context) {
	logger.LogDebug(s.logger, c, "Unbinding license")

	license_key := c.Param("id")

	isvalid := utils.IsUUID(license_key)
	if !isvalid {
		s.HandleError(c, InvalidLicense("not matching UUID format"))
		return
	}

	license, err := s.db.GetLicense(license_key)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			s.HandleError(c, InvalidLicense("license not found"))
			return
		}

		s.HandleError(c, InternalServerError("cannot get license: "+err.Error()))
		return
	}

	if license.HWID == "" {
		s.HandleError(c, &AppError{StatusCode: 400, Message: "License is not bound to any machine", LogMessage: "License is not bound to any machine"})
		return
	}

	err = s.db.UnbindLicense(license_key)
	if err != nil {
		s.HandleError(c, InternalServerError("cannot unbind license: "+err.Error()))
		return
	}

	c.JSON(200, gin.H{"message": "license unbound"})
}

// banLicense changes the license status to banned
func (s *Server) banLicense(c *gin.Context) {
	logger.LogDebug(s.logger, c, "Banning license")

	license_key := c.Param("id")

	isvalid := utils.IsUUID(license_key)
	if !isvalid {
		s.HandleError(c, InvalidLicense("not matching UUID format"))
		return
	}

	license, err := s.db.GetLicense(license_key)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			s.HandleError(c, InvalidLicense("license not found"))
			return
		}

		s.HandleError(c, InternalServerError("cannot get license: "+err.Error()))
		return
	}

	if license.Status != "active" {
		s.HandleError(c, &AppError{StatusCode: 400, Message: "License is not active, current status: " + license.Status, LogMessage: "License is not active: " + license.Status})
		return
	}

	err = s.db.BanLicense(license_key)
	if err != nil {
		s.HandleError(c, InternalServerError("cannot ban license: "+err.Error()))
		return
	}

	c.JSON(200, gin.H{"message": "license banned"})
}

// updateLicense updates the license information based on the payload
func (s *Server) updateLicense(c *gin.Context) {
	logger.LogDebug(s.logger, c, "Updating license")

	var payload types.License

	err := c.BindJSON(&payload)
	if err != nil {
		s.HandleError(c, InvalidPayload("cannot bind payload: "+err.Error()))
		return
	}

	isvalidpayload := utils.IsValidPayload(&payload)
	if !isvalidpayload {
		s.HandleError(c, InvalidPayload("missing required fields"))
		return
	}

	_, err = s.db.GetLicense(payload.LicenseKey)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			s.HandleError(c, InvalidLicense("license not found"))
			return
		}

		s.HandleError(c, InternalServerError("cannot get license: "+err.Error()))
		return
	}

	err = s.db.UpdateLicense(payload)
	if err != nil {
		s.HandleError(c, InternalServerError("cannot update license: "+err.Error()))
		return
	}

	c.JSON(200, gin.H{"message": "license updated"})
}
