package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
	"os"
	"path/filepath"
)

func main() {
	wd, _ := os.Getwd()
	g := gen.NewGenerator(gen.Config{
		OutPath:           filepath.Join(wd, "./app/4clabs-server/internal/adapter/data/query"),
		FieldWithTypeTag:  true,
		FieldWithIndexTag: true,
	})
	db, _ := gorm.Open(mysql.Open("dev:Zhengli_0220@tcp(master-pub.mysql.polardb.rds.aliyuncs.com:3306)/4clabs?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai"))
	g.UseDB(db)

	user := g.GenerateModel("users")
	g.ApplyBasic(user)
	g.Execute()
}
