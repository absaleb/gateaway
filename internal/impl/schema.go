package impl

import (
	"github.com/absaleb/gateway/internal/resolvers"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"log"
	"net/http"
)

func getGraphqlHandler() http.Handler {
	schema, err := graphql.NewSchema(compileSchema())
	if err != nil {
		log.Panic(err)
	}

	h := handler.New(&handler.Config{
		Schema:     &schema,
		Pretty:     true,
		GraphiQL:   false,
		Playground: true,
	})

	return h
}

func createSchema(query, mutation, subscription graphql.Fields) graphql.SchemaConfig {
	var q *graphql.Object
	if len(query) > 0 {
		q = graphql.NewObject(graphql.ObjectConfig{
			Name:   "RootQuery",
			Fields: query,
		})
	}

	var m *graphql.Object
	if len(mutation) > 0 {
		m = graphql.NewObject(graphql.ObjectConfig{
			Name:   "RootMutation",
			Fields: mutation,
		})
	}

	var s *graphql.Object
	if len(subscription) > 0 {
		s = graphql.NewObject(graphql.ObjectConfig{
			Name:   "RootSubscription",
			Fields: subscription,
		})
	}

	schemaConfig := graphql.SchemaConfig{
		Query:        q,
		Mutation:     m,
		Subscription: s,
	}

	return schemaConfig
}

func compileSchema() graphql.SchemaConfig {
	queryRoot := make(graphql.Fields)
	mutationRoot := make(graphql.Fields)
	subscription := make(graphql.Fields)

	resolvers.CreateNotifierResolvers(&queryRoot, &mutationRoot)

	return createSchema(queryRoot, mutationRoot, subscription)
}
