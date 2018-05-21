package tasks

import "gopkg.in/mgo.v2/bson"

type Task struct {
  ID          bson.ObjectId   `bson:"_id" json:"id"`
  Title       string          `json:"title"`
  Completed   bool            `json:"completed"`
}

type Tasks []Task