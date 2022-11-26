package guacamole

type GuacConnectionGroup struct {
	Name              string                        `json:"name"`
	Identifier        string                        `json:"identifier"`
	ParentIdentifier  string                        `json:"parentIdentifier"`
	Type              string                        `json:"type"`
	ActiveConnections int                           `json:"activeConnections"`
	ChildConnections  []GuacConnection              `json:"childConnections"`
	ChildGroups       []GuacConnectionGroup         `json:"childConnectionGroups"`
	Attributes        GuacConnectionGroupAttributes `json:"attributes"`
}

type GuacConnectionGroupAttributes struct {
	MaxConnections        string `json:"max-connections"`
	MaxConnectionsPerUser string `json:"max-connections-per-user"`
	EnableSessionAffinity string `json:"enable-session-affinity"`
}
