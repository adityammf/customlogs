package services

import (
	"customlogs/utility"

	"github.com/go-openapi/errors"
)

var logger = utility.Log

// IsAuthentic : auth token
func IsAuthentic(token string) (*bool, error) {
	true := true
	// logger.Debugf("Secret Key: %s", token)
	if token == "8yhZ34W9AEZPm40oslDDPpWBey5wNF55" {
		return &true, nil
	}
	logger.ErrorMsg("nrtapi-authentication", errors.New(401, "Unauthorized: Invalid Credential"))
	logger.Debug("Bad credentials")
	return nil, errors.New(401, "Unauthorized: Invalid Credential")
}
