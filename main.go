package main

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

import (
	"go-blog-build/config"
	"go-blog-build/models"
)

func main() {
	db, err := config.ConnectDB()
	if err != nil {
		panic("failed to connect database")
	}

	// 自动迁移表
	_ = db.AutoMigrate(&models.User{}, &models.Post{}, &models.Comment{})
}
