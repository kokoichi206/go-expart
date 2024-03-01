package graph

import "graphql-github-sample/graph/services"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

// DI のための Resolver 構造体。
type Resolver struct {
	Srv services.Services
	*Loaders
}
