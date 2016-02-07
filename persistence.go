package engine

import (
    "fmt"
    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
)

type Subject struct {
    id bson.ObjectId `bson:"_id,omitempty"`
    username string
    email string
}

type Project struct {
    id bson.ObjectId `bson:"_id,omitempty"`
    subject Subject `bson:",inline"`
    ownerId string
    name string
    description string
}

type Document struct {
    id bson.ObjectId `bson:"_id,omitempty"`
    project Project `bson:",inline"`
    path string
}

type Persistence struct {
    session *mgo.Session
}

func (persistence Persistence) Init() (*Persistence, error) {
    session, err := mgo.Dial("192.168.99.100")
    if err != nil {
        panic(err)
    }

    fmt.Println("Connected to mongo")
    return &Persistence{session}, nil
}
