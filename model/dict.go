package model

type Dict struct {
	ID uint `json:"id" gorm:"primaryKey"`
	From string `json:"from"`
	To string `json:"to"`
	Time string `json:"time"`
	Count uint `json:"count"`
}

func NewDict(from string,to string,time string,count uint) *Dict {
	return &Dict{
		From: from,
		To: to,
		Time:    time,
		Count:   0,
	}
}
