package models

import "github.com/google/uuid"

type Device struct {
	ID              uuid.UUID
	Hostname        string
	Loopback        string
	Model           uint
	SoftwareVersion uint
}
