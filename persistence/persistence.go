package persistence

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Subject struct {
	id       bson.ObjectId `bson:"_id,omitempty"`
	Username string
	Email    string
}

type Project struct {
	id          bson.ObjectId `bson:"_id,omitempty"`
	Subject     Subject       `bson:",inline"`
	OwnerId     string
	Name        string
	Description string
}

type Document struct {
	id      bson.ObjectId `bson:"_id,omitempty"`
	Project Project       `bson:",inline"`
	Path    string
}

type PersistenceSession struct {
	session *mgo.Session
}

func Init() *PersistenceSession {
	session, err := mgo.Dial("192.168.99.100")
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to mongo")
	return &PersistenceSession{session: session}
}

func (persistence PersistenceSession) CreateSubject(subject *Subject) error {
	subject.id = bson.NewObjectId()
	c := persistence.session.DB("huskydocs").C("subject")
	err := c.Insert(subject)
	return err
}

func (persistence PersistenceSession) DeleteSubject(subject *Subject) error {
	c := persistence.session.DB("huskydocs").C("subject")
	err := c.Remove(subject)
	return err
}

func (persistence PersistenceSession) Projects(subject *Subject) ([]Project, error) {
	c := persistence.session.DB("huskydocs").C("project")
	var results []Project
	err := c.Find(bson.M{"subject": subject}).All(&results)
	return results, err
}

func (persistence PersistenceSession) CreateProject(project *Project) error {
	project.id = bson.NewObjectId()
	c := persistence.session.DB("huskydocs").C("project")
	err := c.Insert(project)
	return err
}

func (persistence PersistenceSession) DeleteProject(project *Project) error {
	c := persistence.session.DB("huskydocs").C("project")
	err := c.Remove(project)
	return err
}

func (persistence PersistenceSession) Documents(project *Project) ([]Document, error) {
	c := persistence.session.DB("huskydocs").C("document")
	var results []Document
	err := c.Find(bson.M{"project": project}).All(&results)
	return results, err
}

func (persistence PersistenceSession) CreateDocument(document *Document) error {
	document.id = bson.NewObjectId()
	c := persistence.session.DB("huskydocs").C("document")
	err := c.Insert(document)
	return err
}

func (persistence PersistenceSession) DeleteDocument(document *Document) error {
	c := persistence.session.DB("huskydocs").C("document")
	err := c.Remove(document)
	return err
}
