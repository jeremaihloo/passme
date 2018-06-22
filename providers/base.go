package providers

import "github.com/jeremaihloo/passme"

type ProviderInterface interface {
	Name() error
	Push(storage passme.Storage) error
	Fetch(storage passme.Storage) error
	Merge(storage passme.Storage) error
}
