package model

func migration() {
	// 自动迁移  即将go数据结构映射创建为mysql表结构
	DB.Set("gorm:table_options", "charset=utf8mb4").
		AutoMigrate(&User{}).
		AutoMigrate(&Task{})
	//添加外键关联
	DB.Model(&Task{}).AddForeignKey("uid", "user(id)", "CASCADE", "CASCADE")
}
