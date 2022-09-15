package config

type GoogleOAuth struct {
	ClientId                string `json:"client_id"`
	ProjectId               string `json:"project_id"`
	AuthUri                 string `json:"auth_uri"`
	TokenUri                string `json:"token_uri"`
	AuthProviderX509CertUrl string `json:"auth_provider_x509_cert_url"`
	ClientSecret            string `json:"client_secret"`
}

type GoogleConfig struct {
	Web *GoogleOAuth `json:"web"`
}
