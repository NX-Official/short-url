package model

type Url struct {
	Model
	ShortURL string `gorm:"unique;type:varchar(255);not null;comment:短链接;"`
	LongURL  string `gorm:"type:varchar(255);not null;comment:原始链接"`
}
