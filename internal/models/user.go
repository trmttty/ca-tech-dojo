package models

import "time"

type User struct {
	ID        int       `json:"id"`
	UserName  string    `json:"name"`
	Token     string    `json:"token"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// type UserCreateRequest struct {
// 	UserName string `json:"name"`
// }

// type UserCreateResponse struct {
// 	Token string `json:"token"`
// }

// type UserGetResponse struct {
// 	UserName string `json:"name"`
// }

// type UserUpdateRequest struct {
// 	UserName string `json:"name"`
// }

// type GachaDrawRequest struct {
// 	GachaTimes int `json:"times"`
// }

// type GachaDrawResponse struct {
// 	GachaResults []GachaResult `json:"results"`
// }

// type GachaResult struct {
// 	CharacterID   string `json:"characterID"`
// 	CharacterName string `json:"name"`
// }

// type CharacterListResponse struct {
// 	UserCharacters []UserCharacter `json:"characters"`
// }

// type UserCharacter struct {
// 	UserCharacterID string `json:"userCharacterID"`
// 	CharacterID     string `json:"characterID"`
// 	UserName        string `json:"name"`
// }
