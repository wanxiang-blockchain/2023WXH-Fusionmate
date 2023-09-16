package dao

import (
	"github.com/FusionMate/fm-backend/common/log"
	"github.com/FusionMate/fm-backend/model"
)

// 创建表
func CreateCollectionTbl() error {
	if !fmdb.Migrator().HasTable(&model.Collection{}) {
		err := fmdb.Migrator().CreateTable(&model.Collection{})
		if err != nil {
			panic(err)
		}
	}

	err := fmdb.AutoMigrate(&model.Collection{})
	if err != nil {
		panic(err)
	}
	return err
}

func InsertCollection(c *model.Collection) error {
	if err := fmdb.Create(c).Error; err != nil {
		log.Error("[DB][collection] insert collection failed: %s", err)
		return err
	}
	return nil
}

// 更新nft合约地址
func UptCollectionAddress(c *model.Collection) error {
	oldC := model.Collection{}
	if err := fmdb.First(&oldC, "collectionID = ?", c.CollectionID).Error; err != nil {
		log.Error("[DB][collection] update contractAddress failed,collectionID %d not exist, err: %s",
			c.CollectionID, err.Error())
		return err
	}

	if oldC.ContractAddr == c.ContractAddr {
		return nil
	}

	if err := fmdb.Model(c).Update("contractAddr", c.ContractAddr).Error; err != nil {
		log.Error("[DB][collection] update contractAddress failed: %s", err.Error())
		return err
	}
	return nil
}

func UptCollectionBaseURI(c *model.Collection) error {
	oldC := model.Collection{}
	if err := fmdb.First(&oldC, "collectionID = ?", c.CollectionID).Error; err != nil {
		log.Error("[DB][collection] update baseURI failed,collectionID %d not exist, err: %s",
			c.CollectionID, err.Error())
		return err
	}

	if oldC.BaseURI == c.BaseURI {
		return nil
	}

	if err := fmdb.Model(c).Update("baseURI", c.BaseURI).Error; err != nil {
		log.Error("[DB][collection] update baseURI failed: %s", err.Error())
		return err
	}
	return nil
}

func GetCollectionsByPage(typ string, makerAddr string, page int, pageSize int) ([]model.Collection, error) {
	var result []model.Collection
	var err error
	offset := page * pageSize

	if len(typ) != 0 && len(makerAddr) != 0 {
		err = fmdb.Where("typ = ? and maker = ?", typ, makerAddr).Offset(offset).Limit(pageSize).Find(&result).Error
	} else if len(typ) != 0 {
		err = fmdb.Where("typ = ?", typ).Offset(offset).Limit(pageSize).Find(&result).Error
	} else if len(makerAddr) != 0 {
		err = fmdb.Where("maker = ?", makerAddr).Offset(offset).Limit(pageSize).Find(&result).Error
	} else {
		err = fmdb.Offset(offset).Limit(pageSize).Find(&result).Error
	}

	if err != nil {
		log.Error("[DB][collection] query by page fail: %s", err.Error())
	}
	return result, err
}

func GetConnectionByID(collectionID int64) (model.Collection, error) {
	var c model.Collection
	err := fmdb.First(&c, "collectionID = ?", collectionID).Error
	if err != nil {
		log.Error("[DB][collection] query collection by ID %d fail: %s", collectionID, err.Error())
	}
	return c, err
}

// 获取当前最后一个数据的自增ID
func GetCollectionLastID() (int64, error) {
	var c model.Collection
	if err := fmdb.Last(&c); err != nil {
		log.Error("[DB][collection] query last Collection fail: %s", err.Error)
	}
	return c.CollectionID, nil
}
