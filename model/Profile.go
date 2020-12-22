package model

import "ginblog/utils/errmsg"

type Profile struct {
	ID int `gorm:"primaryKey" json:"id"`
	Img      string `gorm:"type:varchar(100)" json:"img"`
	Avatar   string `gorm:"type:varchar(100)" json:"avatar"`
	Name     string `gorm:"type:varchar(100)" json:"name"`
	Desc     string `gorm:"type:varchar(150)" json:"desc"`
	QqChat   string `gorm:"type:varchar(100)" json:"qq_chat"`
	Email    string `gorm:"type:varchar(100)" json:"email"`
	Weibo    string `gorm:"type:varchar(100)" json:"weibo"`
	Bili string `gorm:"type:varchar(100)" json:"bili"`
	Wechat string `gorm:"type:varchar(100)" json:"wechat"`
}

// 获取个人设置的信息
func GetProfile(id int) (Profile,int) {
	var profile Profile
	err = db.Where("ID = ?",id).First(&profile).Error
	if err != nil {
		return profile, errmsg.ERROR
	}
	return profile, errmsg.SUCCSE

}

func UpdateProfile(data *Profile) int {
	var profile Profile
	err = db.Where("id = ?",1).Model(&profile).Updates(data).Error
	if err != nil{
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}
