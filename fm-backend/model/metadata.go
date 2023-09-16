package model

// TokenMetaData
// NFT TokenURI对应的MetaData
type TokenMetaData struct {
	CollectionID int         `json:"collectionID"`
	TokenID      int         `json:"tokenID"`
	Description  string      `json:"description"` // 描述，创建NFT表单中的选填字段
	Image        string      `json:"image"`
	Name         string      `json:"name"`
	Attributes   []Attribute `json:"attributes"`
}

type Attribute struct {
	TraitType string `json:"traitType"`
	Value     string `json:"value"`
}
