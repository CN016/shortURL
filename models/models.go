package models

type SU struct {
	ID       int64 `gorm:"primary_key;AUTO_INCREMENT"`
	ShortUrl string
	Url      string
}

type Tabler interface {
	TableName() string
}

func (SU) TableName() string {
	return "t_short_url"
}
