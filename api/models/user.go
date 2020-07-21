package models

import (
	orm "gin-demo/api/database"
)

type User struct {
	ID       int64  `json:"id"`       // 列名为 `id`
	Username string `json:"username"` // 列名为 `username`
	Password string `json:"password"` // 列名为 `password`
}

var Users []User

// 添加
func (user User) Insert(users User) (id int64, err error) {
	//添加数据
	result := orm.Eloquent.Create(&users)
	id = user.ID
	if result.Error != nil {
		err = result.Error
		return
	}
	return
}

// 列表
func (user *User) Users() (users []User, err error) {
	if err = orm.Eloquent.Find(&users).Error; err != nil {
		return
	}
	return
}

// get
func (user *User) GetUser(username string) (users []User, err error) {
	if err = orm.Eloquent.Where("username = ?", username).Find(&users).Error; err != nil {
		return
	}
	return
}

// 修改
func (user *User) Update(data User) (updateUser User, err error) {

	//参数1:是要修改的数据
	//参数2:是 修改的数据
	if err = orm.Eloquent.Model(&updateUser).Where("id = ?", data.ID).Updates(&data).Error; err != nil {
		return
	}
	return
}

// 删除数据
func (user *User) Destroy(id int64) (rows int64, err error) {
	user.ID = id
	db := orm.Eloquent.Unscoped().Delete(&user)
	if err = db.Error; err != nil {
		return
	}
	//if err = orm.Eloquent.Where("id = ? ", id).Delete(&user).Error; err != nil {
	//	return
	//}
	rows = db.RowsAffected
	return
}
