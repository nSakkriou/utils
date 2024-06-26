package agent

type AgentFile struct {
	Port              int          `json:"port"`
	ScriptsFolderPath string       `json:"scripts_folder_path"`
	EndCommands       []EndCommand `json:"endcommands"`

	UseCors    bool       `json:"use_cors"`
	CorsOption CorsOption `json:"cors_option"`
}

type EndCommand struct {
	EndpointName      string   `json:"endpoint_name"`
	Method            string   `json:"method"`
	ScriptsFilesNames []string `json:"scripts_files_names"`
}

type CorsOption struct {
	AllowedOrigins []string `json:"allowed_origins"`
}
