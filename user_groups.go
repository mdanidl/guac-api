package guacamole

type GuacUserGroup struct {
	Identifier string                  `json:"identifier"`
	Attributes GuacUserGroupAttributes `json:"attributes"`
}

type GuacUserGroupAttributes struct {
	Disabled string `json:"disabled"`
}
