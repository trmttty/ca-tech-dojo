package data

type UserCreateRequest struct {
	Name string `json:"name"`
}

type UserCreateResponse struct {
	Token string `json:"token"`
}

type UserGetResponse struct {
	Name string `json:"name"`
}

type UserUpdateRequest struct {
	Name string `json:"name"`
}

type GachaDrawRequest struct {
	Times int `json:"times"`
}

type GachaDrawResponse struct {
	Results []GachaResult `json:"results"`
}

type GachaResult struct {
	CharacterID string `json:"characterID"`
	Name        string `json:"name"`
}

type UserCharacter struct {
	UserCharacterID string `json:"userCharacterID"`
	CharacterID     string `json:"characterID"`
	Name            string `json:"name"`
}

type CharacterListResponse struct {
	Characters []UserCharacter `json:"characters"`
}
