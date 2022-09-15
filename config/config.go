package config

type Configuration struct {
	GodTokens   []string      `json:"god_tokens"`
	DriveConfig *GoogleConfig `json:"drive_config"`
	GmailConfig *GoogleConfig `json:"gmail_config"`
}
