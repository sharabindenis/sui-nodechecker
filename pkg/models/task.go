package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

type Task struct {
	//gorm.Model
	Ip       string `json:"ip"`
	Owner    string `json:"owner"`
	IsActive bool   `json:"isactive"`
}

//var Tasks map[string]string {
//	"ip": Ip,
//	"owner": Owner,
//	"isactive": IsActive
//}
var alltasks = []Task{}

func (b *Task) CreateTask() *Task {

	return b
}

func GetTasks() []Task {
	var Tasks []Task

	return Tasks
}

//func init() {
//	config.Connect()
//	db = config.GetDB()
//	db.AutoMigrate(&Node{})
//}

func GetAllTasks() []Task {
	var Nodes []Task
	db.Find(&Nodes)
	return Nodes
}
