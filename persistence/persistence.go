package persistence

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Subject struct {
	Id       bson.ObjectId `bson:"_id,omitempty"`
	Username string
	Email    string
}

type Project struct {
	Id          bson.ObjectId `bson:"_id,omitempty"`
	Owner       mgo.DBRef
	Name        string
	Description string
}

type Document struct {
	Id      bson.ObjectId `bson:"_id,omitempty"`
	Project mgo.DBRef
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
	subject.Id = bson.NewObjectId()
	c := persistence.session.DB("huskydocs").C("subject")
	err := c.Insert(subject)
	return err
}

func (persistence PersistenceSession) Subject(username string) (Subject, error) {
	c := persistence.session.DB("huskydocs").C("subject")
	var subject Subject
	err := c.Find(bson.M{"username": username}).One(&subject)
	return subject, err
}

func (persistence PersistenceSession) DeleteSubject(subject *Subject) error {
	c := persistence.session.DB("huskydocs").C("subject")
	projects, err := persistence.Projects(subject)

	if err != nil {
		return err
	}

	for i := range projects {
		var project = projects[i]
		err = persistence.DeleteProject(&project)
		if err != nil {
			return err
		}
	}
	err = c.RemoveId(subject.Id)
	return err
}

func (persistence PersistenceSession) Project(owner *Subject, project string) (Project, error) {
	c := persistence.session.DB("huskydocs").C("project")
	ownerRef := mgo.DBRef{Collection: "subject", Id: owner.Id}
	var result Project
	err := c.Find(bson.M{"owner": ownerRef, "name": project}).One(result)
	return result, err
}

func (persistence PersistenceSession) Projects(owner *Subject) ([]Project, error) {
	c := persistence.session.DB("huskydocs").C("project")
	ownerRef := mgo.DBRef{Collection: "subject", Id: owner.Id}
	var results []Project
	err := c.Find(bson.M{"owner": ownerRef}).All(&results)
	return results, err
}

func (persistence PersistenceSession) CreateProject(project *Project) error {
	project.Id = bson.NewObjectId()
	c := persistence.session.DB("huskydocs").C("project")
	err := c.Insert(project)
	return err
}

func (persistence PersistenceSession) DeleteProject(project *Project) error {
	c := persistence.session.DB("huskydocs").C("project")
	documents, err := persistence.Documents(project)
	if err != nil {
		return err
	}

	for j := range documents {
		err = persistence.DeleteDocument(&documents[j])
		if err != nil {
			return err
		}
	}
	err = c.RemoveId(project.Id)
	return err
}

func (persistence PersistenceSession) Documents(project *Project) ([]Document, error) {
	c := persistence.session.DB("huskydocs").C("document")
	projectRef := mgo.DBRef{Collection: "project", Id: project.Id}
	var results []Document
	err := c.Find(bson.M{"project": projectRef}).All(&results)
	return results, err
}

func (persistence PersistenceSession) CreateDocument(document *Document) error {
	document.Id = bson.NewObjectId()
	c := persistence.session.DB("huskydocs").C("document")
	err := c.Insert(document)
	return err
}

func (persistence PersistenceSession) DeleteDocument(document *Document) error {
	c := persistence.session.DB("huskydocs").C("document")
	err := c.RemoveId(document.Id)
	return err
}
