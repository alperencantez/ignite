package service

import (
	"github.com/alperencantez/ignite/database"
	"github.com/alperencantez/ignite/model"
	"gorm.io/gorm"
)

func AddUser(modelInstance *model.User) *gorm.DB {
	return database.DB.Create(modelInstance)
}
