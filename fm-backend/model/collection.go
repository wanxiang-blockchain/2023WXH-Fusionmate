package model

type Collection struct {
	//gorm.Model

	CollectionID int64 `gorm:"primary_key;AUTO_INCREMENT;not null;column:collectionID" json:"collectionID"`

	// AssistantFactory.sol
	Name      string `gorm:"column:name" json:"name"`
	Symbol    string `gorm:"column:symbol" json:"symbol"`
	BaseURI   string `gorm:"column:baseURI" json:"baseURI"`
	MaxSupply int64  `gorm:"column:maxSupply;" json:"maxSupply"` // 最大可mint数量
	MintPrice string `gorm:"column:mintPrice;" json:"mintPrice"` // mint cost

	// 剩余表单字段
	Description string `gorm:"column:description" json:"description"` // 描述，创建NFT表单中的选填字段
	Derive      int64  `gorm:"column:derive" json:"derive"`           // 衍生来源，值为创建者拥有的NFT合约的地址
	Prompts     string `gorm:"column:prompts" json:"prompts"`         // 提示词，创建NFT表单中的必填字段

	//其他字段
	ContractAddr string `gorm:"column:contractAddr" json:"contractAddr"` // NFT合约地址
	Maker        string `gorm:"column:maker" json:"maker"`               // 创建者地址
	ImgURI       string `gorm:"column:imgURI" json:"imgURI"`             // 图像URI
	Typ          string `gorm:"column:typ" json:"type"`                  // Assistant类别
}

func (c *Collection) TableName() string {
	return "collectionTbl"
}
