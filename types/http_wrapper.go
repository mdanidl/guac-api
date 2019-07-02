package types

type ConnectResponse struct {
	AuthToken            string   `json:"authToken"`
	Username             string   `json:"username"`
	Datasource           string   `json:"dataSource"`
	Availabledatasources []string `json:"availableDataSource"`
}

type ErrorResponse struct {
	Message string `json:"message"`
	// TranslatableMessage []map[string]string `json:"translatableMessage"`
	StatusCode string `json:"statusCode"`
	Expected   string `json:"expected"`
	Type       string `json:"type"`
}
