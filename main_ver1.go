package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gen"
	"gorm.io/gorm"
)

//
//var user User
//
//func GetAll(db *gorm.DB) ([]User, error) {
//	var users []User
//	err := db.Model(&User{}).Preload("Ordertabs").Find(&users).Error
//	return users, err
//}

func main1() {

	// https://github.com/go-gorm/postgres
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  "user=postgres password=123456 dbname=gorm port=5432 sslmode=disable",
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})

	if err != nil {

	}

	//fmt.Println(db.Config)

	g := gen.NewGenerator(gen.Config{
		OutPath: "query",
	})

	g.UseDB(db)
	g.ApplyBasic(g.GenerateAllTable()...)

	//u := query.Use(db).User
	//u.WithContext(context.Background()).Where()
	//users, err := u.WithContext(context.Background()).First()
	//
	//if err != nil {
	//
	//}
	//
	//fmt.Println(users)
	g.Execute()
	//
	//db.Where()

	//users, _ := GetAll(db)
	//
	//for _, u := range users {
	//	fmt.Println(u.Ordertabs)
	//}

	//db.AutoMigrate(&User{})
	//
	//db.First(&user)
	//
	////fmt.Println(db.)

	//var users = []User{{Name: "jinzhu1"}, {Name: "jinzhu2"}, {Name: "jinzhu3"}}
	//db.Create(&users)

	//db.Model(&User{}).Create(map[string]interface{}{
	//	"Name": "Jinju",
	//})

	//result := map[string]interface{}{}
	//db.Find(&users)
	//for _, u := range users {
	//	fmt.Println(u)
	//}
	//	tx := db.Where(&User{Name:"da" }
	////db.Where(&{User{Name: }})
	//	if true {
	//		tx = tx.Where("email = ?", 4)
	//	} else {
	//		tx = tx.Where("id = ?", 8)
	//	}
	//
	//	tx.Find(&user)
	//
	//	fmt.Println(user)
}
