package tasks

import (
  "log"

  mgo "gopkg.in/mgo.v2"
  "gopkg.in/mgo.v2/bson"
  "github.com/BurntSushi/toml"
)

type DBConfig struct {
  Server   string
  Database string
  Collection string
}

var dbConfig = DBConfig{}
var db *mgo.Database

func init() {
  dbConfig.Read()
  dbConfig.Connect()
}

func (c *DBConfig) Read() {
  if _, err := toml.DecodeFile("database.toml", &c); err != nil {
    log.Fatal(err)
  }
}

// Establish a connection to database
func (c *DBConfig) Connect() {
  session, err := mgo.Dial(c.Server)
  if err != nil {
    log.Fatal(err)
  }
  db = session.DB(c.Database)
}

// Find list of tasks
func (c *DBConfig) FindAll() ([]Task, error) {
  var tasks []Task
  err := db.C(c.Collection).Find(bson.M{}).Sort("-_id").All(&tasks)
  return tasks, err
}

// Find a task by its id
func (c *DBConfig) FindById(id string) (Task, error) {
  var task Task
  err := db.C(c.Collection).FindId(bson.ObjectIdHex(id)).One(&task)
  return task, err
}

// Insert a task into database
func (c *DBConfig) Insert(task Task) error {
  err := db.C(c.Collection).Insert(&task)
  return err
}

// Delete a task
func (c *DBConfig) Delete(task Task) error {
  err := db.C(c.Collection).Remove(&task)
  return err
}

// Update a task
func (m *DBConfig) Update(task Task) error {
  err := db.C(m.Collection).UpdateId(task.ID, &task)
  return err
}
