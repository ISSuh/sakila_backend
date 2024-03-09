package model

type FilmText struct {
	FilmId int `gorm:"primaryKey"`

	Title       string
	Description string
}

func (FilmText) TableName() string {
	return "film_text"
}
