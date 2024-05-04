package domain

import "context"

// AccessControl
//
//go:generate mockery --name AccessControl
type AccessControl interface{}

// DDNS
//
//go:generate mockery --name DDNS
type DDNS interface {
	// GetRegions
	GetRegions(ctx context.Context)
}

// LightNode
//
//go:generate mockery --name LightNode
type LightNode interface {
	// CheckAvailability
	CheckAvailability(context.Context, AvailableParams) (Available, error)
}

// MessageBroker
//
//go:generate mockery --name MessageBroker
type MessageBroker interface{}
