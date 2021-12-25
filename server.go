package main

import (
	"config"
	"fmt"
	"log"
	"time"

	"gopkg.in/mgo.v2/bson"
)

type todolist struct {
	Id        bson.ObjectId `bson:"_id,omitempty" json:"id,omitempty"`
	Content   string        `json:"content,omitempty"`
	CreatedAt time.Time     `json:"created_at, omitempty"`
	UpdatedAt time.Time     `json:"updated_at, omitempty"`
}

func main() {
	config.Init()
	s := config.NewSession()
	defer s.Close()

	c := s.Copy().DB(config.AppConfig.Database).C("todolists")
	todolist := todolist{
		Id:        bson.NewObjectId(),
		Content:   "New Content",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err := c.Insert(&todolist)
	if err != nil {
		log.Fatalf("[InsertDB]: %s\n", err)
	}
	fmt.Println("Loi")
	blogs := []todolist{}
	c.Find(nil).All(&todolist)

	// Out put
	fmt.Println(blogs)

}
