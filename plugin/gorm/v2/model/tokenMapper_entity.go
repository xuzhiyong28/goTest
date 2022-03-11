package model

import "time"

type TokenMappers []TokenMapper

type TokenMapper struct {
	ID         int       `json:"id" gorm:"autoIncrement;primaryKey"`
	RootToken  string    `json:"root_token",gorm:"column:root_token;not null"`
	ChildToken string    `json:"child_token",gorm:"column:child_token;not null"`
	Mintable   string    `json:"mintable",gorm:"column:mintable;not null"`
	MapType    string    `json:"map_type",gorm:"column:map_type;not null"`
	Name       string    `json:"name",gorm:"column:name"`
	Symbol     string    `json:"symbol",gorm:"column:symbol"`
	Deleted    string    `json:"deleted",gorm:"column:deleted"`
	ChainId    int       `json:"chain_id",gorm:"column:chain_id;not null"`
	CreatedAt  time.Time `json:"created_at",gorm:"column:created_at"`
	UpdatedAt  time.Time `json:"updated_at",gorm:"column:updated_at"`
}

func (tm *TokenMapper) TableName() string {
	return "token_mapper"
}
