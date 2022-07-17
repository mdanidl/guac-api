package guacamole

// GuacPermissionItem a
type GuacPermissionItem struct {
	Op    string `json:"op"`
	Path  string `json:"path"`
	Value string `json:"value"`
}

// GuacPermissionData a
type GuacPermissionData struct {
	ConnectionPermissions       map[string][]string `json:"connectionPermissions"`
	ConnectionGroupPermissions  map[string][]string `json:"connectionGroupPermissions"`
	SharingProfilePermissions   map[string][]string `json:"sharingProfilePermissions"`
	UserPermissions             map[string][]string `json:"userPermissions"`
	UserGroupPermissions        map[string][]string `json:"userGroupPermissions"`
	SystemPermissions           []string            `json:"systemPermissions"`
	ActiveConnectionPermissions map[string][]string `json:"activeConnectionPermissions"`
}
