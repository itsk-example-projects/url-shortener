package models

type Link struct {
	Short  string `json:"short" gorm:"primaryKey"`
	Long   string `json:"long"`
	UserID string `json:"user_id" gorm:"index"`
}
