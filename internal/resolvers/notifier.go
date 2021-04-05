package resolvers

import "github.com/graphql-go/graphql"

var NotifierMessageResponse = graphql.NewObject(graphql.ObjectConfig{
	Name: "Message",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.String,
		},
		"text": &graphql.Field{
			Type: graphql.String,
		},
		"channels": &graphql.Field{
			Type: graphql.String,
		},
	},
	Description: "Message object with text and channels",
})

var NotifierMessageRequest = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "MessageRequest",
	Fields: graphql.InputObjectConfigFieldMap{
		"text": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"channels": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
	},
	Description: "Message object with text and channels",
})

func CreateNotifierResolvers(queryRoot, mutationRoot *graphql.Fields) {
	resolveName := "notifier"

	query := graphql.NewObject(graphql.ObjectConfig{
		Name: "NotifierQuery",
		Fields: graphql.Fields{
			"getMessage": &graphql.Field{
				Type:    graphql.NewList(NotifierMessageResponse),
				Resolve: queryGetMessageResolver,
			},
		},
	})

	(*queryRoot)[resolveName] = &graphql.Field{
		Type: query,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return query, nil
		},
	}

	mutation := graphql.NewObject(graphql.ObjectConfig{
		Name: "NotifierMutation",
		Fields: graphql.Fields{
			"addMessage": &graphql.Field{
				Type: NotifierMessageResponse,
				Args: graphql.FieldConfigArgument{
					"request": &graphql.ArgumentConfig{Type: NotifierMessageRequest},
				},
				Resolve: mutationAddMessgageResolver,
			},
		},
	})

	(*mutationRoot)[resolveName] = &graphql.Field{
		Type: mutation,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return mutation, nil
		},
	}
}

func mutationAddMessgageResolver(p graphql.ResolveParams) (interface{}, error) {
	return nil, nil
}

func queryGetMessageResolver(p graphql.ResolveParams) (interface{}, error) {
	return nil, nil
}
