// Package domain application models and interfaces
package domain

import "context"

// Region
type Region string

const (
	// East east
	East Region = "east"
	// West west
	West Region = "west"
)

// Policy
type Policy struct{}

// ACL
type ACL struct {
	Policy     Policy   `json:"policy"`
	Signatures []string `json:"signatures"`
}

// ComputeParams
type ComputeParams struct {
	CPU    int `json:"cpu"`
	Memory int `json:"memory"`
}

// AvailableParams
type AvailableParams struct {
	Regions       []Region      `json:"regions"`
	ComputeParams ComputeParams `json:"compute_params"`
	ACL           ACL           `json:"acl"`
}

// Available
type Available struct{}

// Provision
type Provision struct{}

// AvailableService
//
//go:generate mockery --name AvailableService
type AvailableService interface {
	// ListRegions
	ListRegions(ctx context.Context) ([]Region, error)
}

// ProvisionService
//
//go:generate mockery --name ProvisionService
type ProvisionService interface{}
