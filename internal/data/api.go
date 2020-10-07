package data

type UserCreateRequest struct {
	Name string
}

type UserCreateResponse struct {
	Token string
}

type UserGetResponse struct {
	Name string
}

type UserUpdateRequest struct {
	Name string
}

type GachaDrawRequest struct {
	Times int
}

type GachaDrawResponse struct {
	Results []GachaResult
}

type GachaResult struct {
	CharacterID string
	Name        string
}

type UserCharacter struct {
	UserCharacterID string
	CharacterID     string
	Name            string
}

type CharacterListResponse struct {
	Characters []UserCharacter
}
