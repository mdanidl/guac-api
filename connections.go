package guacamole

type GuacConnection struct {
	Name              string                   `json:"name"`
	Identifier        string                   `json:"identifier,omitempty"`
	ParentIdentifier  string                   `json:"parentIdentifier"`
	Protocol          string                   `json:"protocol"`
	Attributes        GuacConnectionAttributes `json:"attributes"`
	Properties        GuacConnectionParameters `json:"parameters"`
	ActiveConnections int                      `json:"activeConnections,omitempty"`
	LastActive        int                      `json:"lastActive,omitempty"`
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
	// Telnet & SSH & VNC & RDP parameters
	Backspace               string `json:"backspace"`
	ClipboardEncoding       string `json:"clipboard-encoding"`
	ColorDepth              string `json:"color-depth"`
	CreateRecordingPath     string `json:"create-recording-path"`
	CreateTypescriptPath    string `json:"create-typescript-path"`
	Cursor                  string `json:"cursor"`
	DestPort                string `json:"dest-port"`
	EnableAudio             string `json:"enable-audio"`
	EnableSftp              string `json:"enable-sftp"`
	FontSize                string `json:"font-size"`
	Hostname                string `json:"hostname"`
	Port                    string `json:"port"`
	Password                string `json:"password"`
	ReadOnly                string `json:"read-only"`
	RecordingExcludeMouse   string `json:"recording-exclude-mouse"`
	RecordingExcludeOutput  string `json:"recording-exclude-output"`
	RecordingIncludeKeys    string `json:"recording-include-keys"`
	SftpPort                string `json:"sftp-port"`
	SftpServerAliveInterval string `json:"sftp-server-alive-interval"`
	SwapRedBlue             string `json:"swap-red-blue"`
	TerminalType            string `json:"terminal-type"`
	Username                string `json:"username"`
	// VNC Parameters
	// RDP Parameters
	Console                  string `json:"console"`
	ConsoleAudio             string `json:"console-audio"`
	CreateDrivePath          string `json:"create-drive-path"`
	DisableAudio             string `json:"disable-audio"`
	DisableAuth              string `json:"disable-auth"`
	DisableBitmapCaching     string `json:"disable-bitmap-caching"`
	DisableGlyphCaching      string `json:"disable-glyph-caching"`
	DisableOffscreenCaching  string `json:"disable-offscreen-caching"`
	Domain                   string `json:"domain"`
	Dpi                      string `json:"dpi"`
	EnableAudioInput         string `json:"enable-audio-input"`
	EnableDesktopComposition string `json:"enable-desktop-composition"`
	EnableDrive              string `json:"enable-drive"`
	EnableFontSmoothing      string `json:"enable-font-smoothing"`
	EnableFullWindowDrag     string `json:"enable-full-window-drag"`
	EnableMenuAnimations     string `json:"enable-menu-animations"`
	EnablePrinting           string `json:"enable-printing"`
	EnableTheming            string `json:"enable-theming"`
	EnableWallpaper          string `json:"enable-wallpaper"`
	GatewayPort              string `json:"gateway-port"`
	Height                   string `json:"height"`
	IgnoreCert               string `json:"ignore-cert"`
	PreconnectionID          string `json:"preconnection-id"`
	ResizeMethod             string `json:"resize-method"`
	Security                 string `json:"security"`
	ServerAliveInterval      string `json:"server-alive-interval"`
	ServerLayout             string `json:"server-layout"`
	Width                    string `json:"width"`
	// SSH Parameters
	Command    string `json:"command"`
	HostKey    string `json:"HostKey"`
	Passphrase string `json:"passphrase"`
	PrivateKey string `json:"private-key"`
	// Telnet Parameters
	ColorScheme    string `json:"color-scheme"`
	PasswordRegex  string `json:"password-regex"`
	TypescriptName string `json:"typescript-name"`
	TypescriptPath string `json:"typescript-payh"`
}
