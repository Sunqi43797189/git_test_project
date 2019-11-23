package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

type Moment struct {
	ID uint `json:"id" gorm:"primary_key"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	MemberID int `json:"member_id" gorm:"index"`
	LikeCount int `json:"like_count"`
	Content string `json:"content"`
	Status int `json:"status"`

	MomentVideo MomentVideo
}

type MomentVideo struct {
	ID uint `json:"id" gorm:"primary_key"`
	MemberId uint `json:"member_id"`
	MomentId uint `json:"moment_id"`
	Url string `json:"url"`
	Status int `json:"status"`
	Firstname string `json:"firstname"`
	Extra string `json:"extra"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Member struct {
	ID uint `json:"id"`
	Sex int `json:"sex"`
	Nickname string `json:"nickname"`
	Uid string `json:"uid"`
	Wealth int `json:"wealth"`
	Role string `json:"role"`
	Status int `json:"status"`
	RoseCount uint `json:"rose_count"`
	ActiveAt time.Time `json:"active_at"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Moment []Moment
	MomentVideo []MomentVideo
}

func ConnectMysql() (db *gorm.DB){
	dbType := "mysql"
	dbName := "yidui_test"
	user := "root"
	password := "root"
	host := "127.0.0.1:3306"
	db, err := gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName))
	if err != nil{
		fmt.Println("connect database error")
	}
	return db
}

func main() {
	db := ConnectMysql()
	defer db.Close()
	var res interface{}
	var moments []Moment
	res = db.HasTable(&Moment{})
	fmt.Println(res)
	//
	//res = db.Model(&Moment{}).Where("id = 221")
	//
	//moment := Moment{
	//	ID:          44444,
	//	CreatedAt:   time.Now(),
	//	UpdatedAt:   time.Now(),
	//	MemberID:    1,
	//	LikeCount:   0,
	//	Content:     "测试",
	//	Status:      0,
	//	MomentVideo: MomentVideo{},
	//}
	//res = db.Create(moment) // 创建
	//db.First(&moment)
	//db.Last(&moment)
	//db.Find(&moments)
	//db.First(&moment, 10)
	//db.Where("status = ?", 0).Find(&moments)
	//db.Where("status <> ?", 0).Find(&moments)
	//db.Where("status in (?)", []int{0, 1}).Find(&moments)
	//db.Where("content LIKE ?", "%测试%").Find(&moments)
	//db.Where("content like ? and status = ?", "%测试%", 1).Find(&moments)
	//db.Where("updated_at > ?", time.Now().AddDate(0, 0, -2)).Find(&moments)

	//db.Where(&Moment{ID: 44444}).Find(&moments) //结构体作为查询条件
	//db.Where(map[string]interface{}{"ID":44444}).Find(&moments) // map作为参数查询条件
	//db.Where([]int{12,13,15,16}).Find(&moments) // 主键分片查询

	//db.Not("content like ?", "%测试%").Find(&moments)
	//db.Not("id in (?)", []int{1,2,44444}).Find(&moments)
	//db.Not("id", []int{1,2,44444}).Find(&moments)
	//db.Not(&Moment{ID: 44444}).Find(&moments)

	//db.Find(&moments, "id = ?", 44444)
	//db.Find(&moments, &Moment{ID: 44444})
	//db.Find(&moments, map[string]interface{}{"ID": 44444})
	//db.Where("ID = ?", 44444).Or("content like ?","%哈哈%").Find(&moments)
	//db.Where(&Moment{ID: 44444}).Or(&Moment{Content: "哈哈"}).Find(&moments)
	//db.Where(&Moment{ID: 44444}).Or(map[string]interface{}{"Content": "哈哈"}).Find(&moments)

	//db.FirstOrInit(&moments, Moment{ID: 44445})
	//db.Where(map[string]interface{}{"ID": 44444}).FirstOrInit(&moments)
	//db.Where(Moment{ID: 44445}).Attrs(Moment{Content: "dubug"}).FirstOrCreate(&moments)

	//db.Select("id, content").Where(Moment{ID: 44444}).Find(&moments)

	//db.Where("content like ?", "%哈哈%").Order("id").Limit(3).Find(&moments)
	//var count int
	//db.Where("content like ? ", "%哈哈%").Limit(-1).Find(&moments).Count(&count)
	//db.Where("content like ?", "%哈哈%").Limit(1).Offset(-1).Find(&moments).Count(&count)

	//rows, _ := db.Table("moments").Select("date(created_at) as date, sum(id) as total").Group("date(created_at)").Rows()
	//for rows.Next(){
	//	fmt.Println(rows)
	//}

	//rows, _ := db.Table("moments").Select("moments.id, members.id, content").Joins("left join members on moments.member_id = members.id").Rows()
	//for rows.Next(){
	//	fmt.Println(rows.Columns())
	//}
	//var like_counts []int
	//db.Model(&Moment{}).Find(&moments).Pluck("like_count", &like_counts)
	//var contents []string
	//db.Model(&Moment{}).Where("content like ?", "%哈哈%").Find(&moments).Pluck("content", &contents)

	//type Info struct {
	//	ID int
	//	MemberId int
	//	Content string
	//	LikeCount int
	//}
	//var infos []Info
	////db.Model(&Moment{}).Where("content like ?", "%哈哈%").Scan(&infos) // 搜索结果拷贝到另一个结构体中
	//db.Model(&Moment{}).Select("id, member_id, content, like_count").Where("content like ?", "%%哈哈").Scan(&infos)
	//fmt.Println(infos)
	//fmt.Println(moments)


	//db.Model(&Moment{}).Where("id = 44444").Update(&Moment{Content:"哈哈"})
	//db.Where("id = 44444").Find(&moments)
	//count := db.Model(&Moment{}).Where("content like ?", "%哈哈%").Updates(map[string]interface{}{"like_count": 0}).RowsAffected // 影响的行数
	//fmt.Println(count)
	fmt.Println(moments[0])
	//fmt.Printf("%s", moments[0].Content)
}