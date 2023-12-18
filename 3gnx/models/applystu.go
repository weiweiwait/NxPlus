package models

import (
	"3gnx/dao"
	"fmt"
)

type Student struct {
	ID        uint   `gorm:"id"`
	Username  string `gorm:"username"`
	Xuehao    string `gorm:"xuehao"`
	Class     string `gorm:"class"`
	Direction string `gorm:"direction"`
	Status    int    `gorm:"status"`
}

// 增加报名用户
func CreateApplyUser(student *Student) (err error) {
	err = dao.DB.Table("ApplyStu").Create(&student).Error
	return
}

// 查询一面通过学生，并返回他们的信息
func GetAllApplyUserOne() (studentList []*Student, err error) {
	var status int = 1
	if err = dao.DB.Table("ApplyStu").Where("status >= ?", status).Find(&studentList).Error; err != nil {
		return nil, err
	}
	return
}

// 查询二面通过学生，并返回他们的信息
func GetAllApplyUserTwo() (studentList []*Student, err error) {
	var status int = 2
	if err = dao.DB.Table("ApplyStu").Where("status >= ?", status).Find(&studentList).Error; err != nil {
		return nil, err
	}
	return
}

// 学生自己查询自己的一面面试通过情况
func GetStatusByXuehaoAndStatusOne(xuehao string) (int, error) {
	var student Student
	if err := dao.DB.Table("ApplyStu").Where("xuehao = ?", xuehao).Find(&student).Error; err != nil {
		return -1, err
	}
	fmt.Println(student.Status)
	return student.Status, nil
}

// 学生自己查询自己的二面面试通过情况
func GetStatusByXuehaoAndStatusTwo(xuehao string) (int, error) {
	var student Student
	if err := dao.DB.Table("ApplyStu").Where("xuehao = ?", xuehao).Find(&student).Error; err != nil {
		return -1, err
	}
	fmt.Println(student.Status)
	return student.Status, nil
}

// 设置一面未通过
func SetFailure() error {
	err := dao.DB.Table("ApplyStu").
		Where("status = ?", 0).
		Update("status", -10).
		Error
	if err != nil {
		return err
	}

	return nil
}

// 设置二面未通过
func SetTwoFailure() error {
	err := dao.DB.Table("ApplyStu").
		Where("status = ?").
		Update("status", -10).
		Error
	if err != nil {
		return err
	}

	return nil
}

// 设置一面通过
func SetOneSuccess(xuehao string) error {
	err := dao.DB.Table("ApplyStu").
		Where("xuehao = ?", xuehao).
		Update("status", 1).
		Error
	if err != nil {
		return err
	}

	return nil
}

// 设置二面通过
func SetTwoSuccess(xuehao string) error {
	err := dao.DB.Table("ApplyStu").
		Where("xuehao = ?", xuehao).
		Update("status", 2).
		Error
	if err != nil {
		return err
	}

	return nil
}
