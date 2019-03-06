package types

type GuacConnection struct {
	Name              string                   `json:"name"`
	Identifier        string                   `json:"identifier,omitempty"`
	ParentIdentifier  string                   `json:"parentIdentifier"`
	Protocol          string                   `json:"protocol"`
	Attributes        GuacConnectionAttributes `json:"attributes"`
	Properties        GuacConnectionParameters `json:"parameters"`
	ActiveConnections int                      `json:"activeConnections,omitempty"`
}

type GuacConnectionAttributes struct {
	GuacdEncryption       string `json:"guacd-encryption"`
	FailoverOnly          string `json:"failover-only"`
	Weight                string `json:"weight"`
	MaxConnections        string `json:"max-connections"`
	GuacdHostname         string `json:"guacd-hostname,omitempty"`
	GuacdPort             string `json:"guacd-port"`
	MaxConnectionsPerUser string `json:"max-connections-per-user"`
}
type GuacConnectionParameters struct {
	Port                    string `json:"port"`
	ReadOnly                string `json:"read-only"`
	SwapRedBlue             string `json:"swap-red-blue"`
	Cursor                  string `json:"cursor"`
	ColorDepth              string `json:"color-depth"`
	ClipboardEncoding       string `json:"clipboard-encoding"`
	RecordingExcludeOutput  string `json:"recording-exclude-output"`
	RecordingExcludeMouse   string `json:"recording-exclude-mouse"`
	RecordingIncludeKeys    string `json:"recording-include-keys"`
	CreateRecordingPath     string `json:"create-recording-path"`
	DestPort                string `json:"dest-port"`
	EnableSftp              string `json:"enable-sftp"`
	SftpPort                string `json:"sftp-port"`
	SftpServerAliveInterval string `json:"sftp-server-alive-interval"`
	EnableAudio             string `json:"enable-audio"`
}
