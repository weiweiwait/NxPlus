package models

import "3gnx/dao"

type Account struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func FindAManger(username string) (account *Account, err error) {
	account = new(Account)
	if err = dao.DB.Debug().Table("user").Where("username=?", username).First(account).Error; err != nil {
		return nil, err
	}
	return
}
