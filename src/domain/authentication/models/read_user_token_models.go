package models

type ReadUserTokenPayload struct {
	GUID         string `json:"guid"`
	UserGUID     string `json:"user_guid"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
