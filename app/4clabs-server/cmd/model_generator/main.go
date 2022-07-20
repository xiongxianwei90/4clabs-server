package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gen/field"
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

	db, _ := gorm.Open(mysql.Open("dev:Zhengli_0220@tcp(rm-2ze69093m45mv5730ao.mysql.rds.aliyuncs.com:3306)/forclabs?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai"))
	g.UseDB(db)

	user := g.GenerateModel("users")
	ticketWls := g.GenerateModel("ticket_wls")
	registerNfts := g.GenerateModel("register_nfts")
	comics := g.GenerateModel("comics")
	comics_nfts := g.GenerateModel("comics_nfts", gen.FieldRelate(field.BelongsTo, "Comic", comics, &field.RelateConfig{
		GORMTag: "foreignKey:comics_id",
	}))
	nfts := g.GenerateModel("nfts")
	g.ApplyBasic(user, ticketWls, registerNfts, comics, nfts, comics_nfts)
	g.Execute()
}
