package dbmanegment

import (
	"fmt"
	"gorm.io/gorm"
	"os"
	"gorm.io/driver/mysql"
	"log"
)

//table latout
type Inputlog struct {
	ID uint64 `gorm:"primaryKey;autoIncrement:false;unique"`
	Log  string
}

type MysqlManegment struct {
	DB *gorm.DB
}


func (c *MysqlManegment) Connect ()error{
	pass := os.Getenv("MYSQL_PASSWORD")
	dsn := "root:" + pass + "@tcp(localhost:3306)/todo_db"
	var err error
	c.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	if err := c.DB.AutoMigrate(&Inputlog{}); err != nil {
		log.Fatalln("failed to create entity: ", err)
	}
	return nil
}



//insert a value
func (c *MysqlManegment) Insert(id uint64 , l string){
	c.DB.Create(&Inputlog{ID: id , Log: l})
}


//delete a value
func (c *MysqlManegment) Delete(id uint64){
	c.DB.Delete(&Inputlog{}, id)
}


//read all of the values
func (c *MysqlManegment) Read_all(){
	var logs []Inputlog
	c.DB.Find(&logs)
	for _, log := range logs {
		if log.ID < 10{
			fmt.Printf("ID: %d        Log: %s\n", log.ID, log.Log)
		}else if log.ID < 100{
			fmt.Printf("ID: %d       Log: %s\n", log.ID, log.Log)
		}else{
			fmt.Printf("ID: %d      Log: %s\n", log.ID, log.Log)
		}
	}
}


//read a value
func (c *MysqlManegment) Read(id uint64) string{
	var rlog string
	var logs []Inputlog
	//grabing a value with the desired id
	c.DB.First(&logs, id)
	for _, log := range logs {
		fmt.Printf("ID: %d        Log: %s\n", log.ID, log.Log)
		rlog = log.Log
	}
	return rlog
}

//update a value
func (c *MysqlManegment) Update(id uint64 , l string){
	var logs Inputlog
	//grab an existing value and changing it
	c.DB.First(&logs, id)
	logs.Log  = l
	c.DB.Save(&logs)
}


