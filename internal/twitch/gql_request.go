package twitch

// GQLRequest defines a generic GraphQL request with a Query and input Variables
type GQLRequest struct {
	Query     string                 `json:"query"`
	Variables map[string]interface{} `json:"variables"`
}

// PersistedGQLRequest represents a persisted GQL request
type PersistedGQLRequest struct {
	OperationName string            `json:"operationName"`
	Variables     QueryVariables    `json:"variables"`
	Query         string            `json:"query"`
	Extensions    RequestExtensions `json:"extensions"`
}

type QueryVariables struct {
	Slug string `json:"slug"`
}

type RequestExtensions struct {
	PersistedQuery PersistedQuery `json:"persistedQuery"`
}

type PersistedQuery struct {
	Version    int    `json:"version"`
	Sha256Hash string `json:"sha256Hash"`
}
