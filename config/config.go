package config

import (
	"errors"
	"os"
)

type Config struct {
	RoleID       string
	SecretID     string
	DestFilePath string
	SecretKey    string
	SecretPath   string
}

func GetConfig(c *Config) error {
	c.RoleID = os.Getenv("ROLE_ID")
	c.SecretID = os.Getenv("SECRET_ID")
	c.DestFilePath = os.Getenv("DESTINATION_FILEPATH")
	c.SecretKey = os.Getenv("SECRET_KEY")
	c.SecretPath = os.Getenv("SECRET_PATH")
	if c.RoleID == "" || c.SecretID == "" {
		return errors.New("ROLE_ID and SECRET_ID both need to be set to authenticate with Vault")
	}
	if c.DestFilePath == "" {
		return errors.New("DESTINATION_FILEPATH needs to be set in order to write the secret to the filesystem")
	}
	if c.SecretKey == "" {
		return errors.New("SECRET_KEY needs to be set to identify the Key:Value secret to be retrieved")
	}
	if c.SecretPath == "" {
		return errors.New("SECRET_PATH needs to be set in order to request the secret")
	}
	return nil
}

func getTLSConfig(c *Config) error {
	return nil
}
