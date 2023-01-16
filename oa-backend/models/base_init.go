package models

func Init() {
	//初始化数据库
	initDB()
	//初始化定时任务
	InitCronTabs()
	//初始化系统变量
	InitSystem()
}
