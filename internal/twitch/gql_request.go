package twitch

type GQLRequest struct {
	Query     string                 `json:"query"`
	Variables map[string]interface{} `json:"variables"`
}

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
