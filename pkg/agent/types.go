package agent

type AgentFile struct {
	// MANDATORY
	Port              int          `json:"port"`
	ScriptsFolderPath string       `json:"scripts_folder_path"`
	EndCommands       []EndCommand `json:"endcommands"`

	// OPTIONAL
	UseCors    bool       `json:"use_cors"`
	CorsOption CorsOption `json:"cors_option"`

	// OPTIONAL
	UseCustomNav         bool      `json:"use_custom_nav"`
	CustomNavDescription []NavLink `json:"custom_nav_links"`
}

type EndCommand struct {
	// MANDATORY
	EndpointName      string   `json:"endpoint_name"`
	Method            string   `json:"method"`
	ScriptsFilesNames []string `json:"scripts_files_names"`
}

type CorsOption struct {
	// MANDATORY
	AllowedOrigins []string `json:"allowed_origins"`
}

type NavLink struct {
	// MANDATORY
	Label string `json:"label"`
	Link  string `json:"link"`
}
