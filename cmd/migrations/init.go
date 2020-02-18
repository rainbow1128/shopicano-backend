package migration

import (
	"github.com/jaswdr/faker"
	"github.com/shopicano/shopicano-backend/app"
	"github.com/shopicano/shopicano-backend/log"
	"github.com/shopicano/shopicano-backend/models"
	"github.com/shopicano/shopicano-backend/utils"
	"github.com/shopicano/shopicano-backend/values"
	"github.com/spf13/cobra"
	"time"
)

var MigInitCmd = &cobra.Command{
	Use:   "init",
	Short: "init creates initially required data",
	Run:   initCmd,
}

func initCmd(cmd *cobra.Command, args []string) {
	tx := app.DB().Begin()

	upAdmin := models.UserPermission{
		ID:         values.AdminGroupID,
		Permission: models.AdminPerm,
	}
	if err := tx.Table(upAdmin.TableName()).Create(&upAdmin).Error; err != nil {
		tx.Rollback()
		log.Log().Errorln(err)
		return
	}
	upManager := models.UserPermission{
		ID:         values.ManagerGroupID,
		Permission: models.ManagerPerm,
	}
	if err := tx.Table(upManager.TableName()).Create(&upManager).Error; err != nil {
		tx.Rollback()
		log.Log().Errorln(err)
		return
	}
	upUser := models.UserPermission{
		ID:         values.UserGroupID,
		Permission: models.UserPerm,
	}
	if err := tx.Table(upUser.TableName()).Create(&upUser).Error; err != nil {
		tx.Rollback()
		log.Log().Errorln(err)
		return
	}

	stAdmin := models.StorePermission{
		ID:         values.AdminGroupID,
		Permission: models.AdminPerm,
	}
	if err := tx.Table(stAdmin.TableName()).Create(&stAdmin).Error; err != nil {
		tx.Rollback()
		log.Log().Errorln(err)
		return
	}
	stManager := models.StorePermission{
		ID:         values.ManagerGroupID,
		Permission: models.ManagerPerm,
	}
	if err := tx.Table(stManager.TableName()).Create(&stManager).Error; err != nil {
		tx.Rollback()
		log.Log().Errorln(err)
		return
	}

	password, _ := utils.GeneratePassword("admin")

	u := models.User{
		ID:             utils.NewUUID(),
		Name:           "Shopicano Admin",
		Status:         models.UserActive,
		Phone:          nil,
		ProfilePicture: nil,
		Password:       password,
		PermissionID:   upAdmin.ID,
		Email:          "admin@example.com",
		CreatedAt:      time.Now().UTC(),
		UpdatedAt:      time.Now().UTC(),
	}
	if err := tx.Table(u.TableName()).Create(&u).Error; err != nil {
		tx.Rollback()
		log.Log().Errorln(err)
		return
	}

	a := models.Address{
		ID:        utils.NewUUID(),
		Name:      "Company Address",
		Address:   faker.New().Address().Address(),
		Email:     "example@shopicano.com",
		Phone:     "8801710333333",
		Postcode:  "1209",
		Country:   "Bangladesh",
		City:      "Dhaka",
		UserID:    u.ID,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}
	if err := tx.Table(a.TableName()).Create(&a).Error; err != nil {
		tx.Rollback()
		log.Log().Errorln(err)
		return
	}

	s := models.Settings{
		ID:                     "1",
		Name:                   "Shopicano Marketplace",
		Website:                "http://shopicano.com",
		IsActive:               true,
		CompanyAddressID:       a.ID,
		DefaultCommissionRate:  0,
		IsSignUpEnabled:        false,
		IsStoreCreationEnabled: false,
		TagLine:                "Do it",
		CreatedAt:              time.Now().UTC(),
		UpdatedAt:              time.Now().UTC(),
	}
	if err := tx.Table(s.TableName()).Create(&s).Error; err != nil {
		tx.Rollback()
		log.Log().Errorln(err)
		return
	}

	if err := tx.Commit().Error; err != nil {
		log.Log().Infoln("Migration init failed with err : ", err)
		return
	}
	log.Log().Infoln("Migration init completed")
}
