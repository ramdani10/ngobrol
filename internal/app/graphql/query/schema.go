package query

import (
	"github.com/graphql-go/graphql"
	"github.com/imansuparman/ngobrol/internal/app/graphql/resolver"
	"github.com/imansuparman/ngobrol/internal/app/graphql/types"
)

// GetQuerySchema schema for query / get record
func GetQuerySchema() graphql.ObjectConfig {
	queryFields := graphql.Fields{
		"books": &graphql.Field{
			Type:    graphql.NewList(types.BookType),
			Resolve: resolver.GetBooks,
		},
		"book": &graphql.Field{
			Type: types.BookType,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
			},
			Resolve: resolver.GetBook,
		},
	}

	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: queryFields}

	return rootQuery
}
