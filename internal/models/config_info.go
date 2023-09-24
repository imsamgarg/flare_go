package models

type ConfigInfo struct {
	Host           string `json:"host"`
	Port           string `json:"port"`
	SaveFolderPath string `json:"saveFolderPath"`
}
