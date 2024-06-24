package agent

import (
	"path/filepath"

	"github.com/nSakkriou/utils/pkg/logn"
	"github.com/nSakkriou/utils/pkg/util"
)

// Checker le fichier de config
// Tous les champs rempli
// Au moins 1 script dans la liste
// Pas de doublons dans les endpoints
// checker si les fichiers existes
func ValidConfig(config *AgentFile) bool {
	if config.Port == 0 || config.Port == -1 {
		logn.Error("le n° de port n'est pas correct")
		return false
	}

	if util.IsEmpty(config.ScriptsFolderPath) {
		logn.Error("le champ scripts_folder_path est vide")
		return false
	}

	// Check si le dossier existe
	if util.FileExist(config.ScriptsFolderPath) {
		logn.Error("le dossier %s n'existe pas", config.ScriptsFolderPath)
		return false
	}

	endpointNames := []string{}
	for _, endCommand := range config.EndCommands {
		// Checker les doublons
		if util.ArrayContains(endpointNames, endCommand.EndpointName) {
			logn.Error("le endpoint %s ne peut pas exister en double", endCommand.EndpointName)
			return false
		}

		endpointNames = append(endpointNames, endCommand.EndpointName)

		// Checker si les fichiers existe
		for _, scriptName := range endCommand.ScriptsFilesNames {
			path := filepath.Join(config.ScriptsFolderPath, scriptName)
			if util.FileExist(path) {
				logn.Error("le script %s n'existe pas", path)
				return false
			}
		}
	}

	return true
}
