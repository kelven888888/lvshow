package service

import (
	"context"
	"errors"
	"fmt"
	"ginshop.com/admin/model"
	"ginshop.com/admin/model/common/request"
	"ginshop.com/global"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
}

func (this *User) GetAll(pageInfo request.PageInfo) ([]model.User, int64) {
	var models []model.User

	query := global.SHOP_DB.Model(model.User{})
	if pageInfo.Keyword != "" {
		query.Where("username LIKE ?  ", "%"+pageInfo.Keyword+"%")
	}
	if *pageInfo.Status != 0 {
		query.Where("status =? ", pageInfo.Status)
	}
	if pageInfo.IsTest != 0 {
		query.Where("is_test =? ", pageInfo.IsTest)
	}

	var count int64 = 0
	query.Count(&count)
	err := query.Limit(pageInfo.Limit).Offset((pageInfo.Page - 1) * pageInfo.Limit).Order(" id desc").Find(&models).Error
	if err != nil {
		fmt.Println(err.Error())
		return nil, 0
	}

	return models, count

}
func (this *User) Changepwd(user model.User) error {
	//fmt.Println(user)
	//if user.Password == "" && *user.IsActive != -1 {
	//
	//	err := global.SHOP_DB.Model(model.User{}).Where("id=?", user.Id).Update(
	//		"is_active", *user.IsActive).Error
	//	if err != nil {
	//		fmt.Println(err)
	//		return err
	//	}
	//	return nil
	//}
	//if user.Password == "" && user.IsTest != -1 {
	//
	//	err := global.SHOP_DB.Model(model.User{}).Where("id=?", user.Id).Update(
	//		"is_test", user.IsTest).Error
	//	if err != nil {
	//		fmt.Println(err)
	//		return err
	//	}
	//	return nil
	//}
	var users model.User

	if user.Password != "" {
		//users.Password = utils.EncryptPassworld(utils.MD5V(user.Password))
		hashPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		users.Password = string(hashPassword)
		if err != nil {
			//utils.Fail(ctx, "加密错误", nil)
			return errors.New("加密错误")
		}
	}

	users.IsActive = user.IsActive
	users.IsTest = user.IsTest
	users.Level = user.Level

	err := global.SHOP_DB.Where("id"+
		" = ?", user.Id).First(&model.User{}).Updates(&users).Error
	if *users.IsActive == 0 {
		key := fmt.Sprintf("login_%d", user.Id)
		ctx := context.Background()
		global.SHOP_REDIS.Del(ctx, key)
	}
	if err != nil {

		return errors.New(err.Error())

	}

	return nil

}
func (this *User) Changetradepwd(user model.User) error {
	//fmt.Println(user)
	//if user.Password == "" && *user.IsActive != -1 {
	//
	//	err := global.SHOP_DB.Model(model.User{}).Where("id=?", user.Id).Update(
	//		"is_active", *user.IsActive).Error
	//	if err != nil {
	//		fmt.Println(err)
	//		return err
	//	}
	//	return nil
	//}
	//if user.Password == "" && user.IsTest != -1 {
	//
	//	err := global.SHOP_DB.Model(model.User{}).Where("id=?", user.Id).Update(
	//		"is_test", user.IsTest).Error
	//	if err != nil {
	//		fmt.Println(err)
	//		return err
	//	}
	//	return nil
	//}
	var users model.User

	if user.TradePassword != "" {
		//users.Password = utils.EncryptPassworld(utils.MD5V(user.Password))
		hashPassword, err := bcrypt.GenerateFromPassword([]byte(user.TradePassword), bcrypt.DefaultCost)
		users.TradePassword = string(hashPassword)
		if err != nil {
			//utils.Fail(ctx, "加密错误", nil)
			return errors.New("加密错误")
		}
	}

	err := global.SHOP_DB.Where("id"+
		" = ?", user.Id).First(&model.User{}).Updates(&users).Error

	if err != nil {

		return errors.New(err.Error())

	}

	return nil

}
