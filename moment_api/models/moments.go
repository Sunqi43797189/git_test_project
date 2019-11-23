package models

type Moment struct {
	Model

	MemberID int `json:"member_id" gorm:"index"`
	LikeCount int `json:"like_count"`
	Content string `json:"content"`
	Status int `json:"status"`
}

func GetAllMoments() (moments []Moment){
	Db.Model(&Moment{}).Find(&moments)
	return
}

func GetMomentsCount() (num int) {
	Db.Model(&Moment{}).Count(&num)
	return
}

func MomentExistsById(id int ) bool {
	var moment Moment
	Db.Model(&Moment{}).Where("id = ?", id).First(&moment)
	if moment.ID > 0 {
		return true
	}
	return false
}
func GetMomentById(id string) (moment Moment) {
	Db.Model(&Moment{}).Where("id = ?", id).First(&moment)
	return
}
func UpdateMomentsById(id int, maps map[string]interface{}) bool {
	var moment Moment
	count := Db.Model(&Moment{}).Where("id = ?", id).First(&moment).Update(maps).RowsAffected
	if count > 0{
		return true
	}
	return false
}

func DeleteMomentById(id int) bool{
	var moment Moment
	Db.Model(&Moment{}).Where("id = ?", id).Delete(&moment)
	return true
}
