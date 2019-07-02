package types

type GuacUser struct {
	Username   string             `json:"username"`
	Password   string             `json:"password"`
	Attributes GuacUserAttributes `json:"attributes"`
	LastActive int                `json:"lastActive"`
}

type GuacUserAttributes struct {
	GuacOrganizationalRole string `json:"guac-organizational-role"`
	GuacFullName           string `json:"guac-full-name"`
	Expired                string `json:"expired"`
	Timezone               string `json:"timezone"`
	AccessWindowStart      string `json:"access-window-start"`
	AccessWindowEnd        string `json:"access-window-end"`
	Disabled               string `json:"disabled"`
	ValidFrom              string `json:"valid-from"`
	ValidUntil             string `json:"valid-until"`
}
