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

func (persistence Persistence) CreateSubject(subject *Subject) (error) {
    subject.id = bson.NewObjectId()
    c := persistence.session.DB("huskydocs").C("subject")
    err := c.Insert(subject)
    return err
}

func (persistence Persistence) DeleteSubject(subject *Subject) (error) {
    c := persistence.session.DB("huskydocs").C("subject")
    err := c.Remove(subject)
    return err
}

func (persistence Persistence) CreateProject(project *Project) (error) {
    project.id = bson.NewObjectId()
    c := persistence.session.DB("huskydocs").C("project")
    err := c.Insert(project)
    return err
}

func (persistence Persistence) DeleteProject(project *Project) (error) {
    c := persistence.session.DB("huskydocs").C("project")
    err := c.Remove(project)
    return err
}

func (persistence Persistence) CreateDocument(document *Document) (error) {
    document.id = bson.NewObjectId()
    c := persistence.session.DB("huskydocs").C("document")
    err := c.Insert(document)
    return err
}

func (persistence Persistence) DeleteDocument(document *Document) (error) {
    c := persistence.session.DB("huskydocs").C("document")
    err := c.Remove(document)
    return err
}
