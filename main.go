package main

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

import (
	"go-blog-build/routes"
	"go-blog-build/utils"
)

func main() {

	//db := config.GetDB()
	//// 自动迁移表
	//_ = db.AutoMigrate(&models.User{}, &models.Post{}, &models.Comment{})

	//TODO  to be tested!

	utils.InitLogger()
	router := routes.SetupRouter()
	_ = router.Run(":8080")

}
