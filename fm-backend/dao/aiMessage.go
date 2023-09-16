package dao

import (
	"github.com/FusionMate/fm-backend/common/log"
	"github.com/FusionMate/fm-backend/model"
)

// 创建表
func CreateAIMessageTbl() error {
	if !fmdb.Migrator().HasTable(&model.AIMessage{}) {
		err := fmdb.Migrator().CreateTable(&model.AIMessage{})
		if err != nil {
			panic(err)
		}
	}

	err := fmdb.AutoMigrate(&model.AIMessage{})
	if err != nil {
		panic(err)
	}
	return err
}

func InsertAIMessage(c *model.AIMessage) error {
	if err := fmdb.Create(c).Error; err != nil {
		log.Error("[DB][AIMessage] insert AIMessage failed: %s", err)
		return err
	}
	return nil
}

func UptAIMessageHarvested(c *model.AIMessage) error {
	oldC := model.AIMessage{}
	if err := fmdb.First(&oldC, "id = ?", c.Id).Error; err != nil {
		log.Error("[DB][AIMessage] update AIMessageHarvested failed,id %d not exist, err: %s",
			c.Id, err.Error())
		return err
	}

	if oldC.Harvested == c.Harvested {
		return nil
	}

	if err := fmdb.Model(c).Update("harvested", c.Harvested).Error; err != nil {
		log.Error("[DB][AIMessage] update AIMessageHarvested failed: %s", err.Error())
		return err
	}
	return nil
}

func CountAndUptAIMessageHarvest(collectionID, tokenID int) int64 {
	var cnt int64
	err := fmdb.Model(&model.AIMessage{}).Where("assistant_id = ? and token_id = ? and harvested = ?",
		collectionID, tokenID, 0).Count(&cnt).Update("harvested", 1).Error
	if err != nil {
		log.Error("[DB][AIMessage] CountAndUptAIMessageHarvest failed: %s", err.Error())
	}
	return cnt
}
