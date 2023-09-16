package model

// AIMessage 计算单个NFT的使用次数，用于计数
type AIMessage struct {
	Id           int    `gorm:"primary_key;AUTO_INCREMENT;not null;column:id" json:"id"`
	Timestamp    string `gorm:"type:varchar(40);not null;column:timestamp" json:"timestamp"`
	CollectionId int    `gorm:"type:int;not null;column:assistant_id;uniqueIndex:assistant_token_idx,priority:1;" json:"collectionID"`
	TokenId      int    `gorm:"type:int;not null;column:token_id;uniqueIndex:assistant_token_idx,priority:2;" json:"tokenID"`
	Harvested    int    `gorm:"type:tinyint unsigned;not null;column:harvested" json:"harvested"`
}

func (*AIMessage) TableName() string {
	return "aiMessageTbl"
}
