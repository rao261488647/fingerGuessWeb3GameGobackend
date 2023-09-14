package model

type Games struct {
	Id           int    `json:"id" gorm:"primaryKey ;autoIncrement"`
	GameId       int    `json:"game_id"`
	TokenAddress string `json:"token_address"`
	Bonus        int    `json:"bonus"`
}

func (game *Games) TableName() string {
	return "games"
}
