package db

import (
	"errors"
	"gopkg.in/mgo.v2"
	"github.com/danitlf/project_go/music"
	"gopkg.in/mgo.v2/bson"
)

const MusicCollection = "music"

var ErrDuplicatedMusic = errors.New("Duplicated Music")


type MusicRepository struct {
	session *mgo.Session
}

func (r *MusicRepository) Create(p *music.Music) error {
	session := r.session.Clone()
	defer session.Close()

	collection := session.DB("").C(MusicCollection)
	err := collection.Insert(p)
	if mongoErr, ok := err.(*mgo.LastError); ok {
		if mongoErr.Code == 11000 {
			return ErrDuplicatedMusic
		}
	}
	return err
}

// func (r *MusicRepository) Update(p *music.Music) error {
// 	session := r.session.Clone()
// 	defer session.Close()

// 	collection := session.DB("").C(MusicCollection)
// 	return collection.Update(bson.M{"_id": p.Id}, p)
// }

func (r *MusicRepository) Remove(id string) error {
	session := r.session.Clone()
	defer session.Close()

	collection := session.DB("").C(MusicCollection)
	return collection.Remove(bson.M{"_id": id})
}


func (r *MusicRepository) FindById(id string) (*music.Music, error) {
	session := r.session.Clone()
	defer session.Close()

	collection := session.DB("").C(MusicCollection)
	query := bson.M{"_id": id}

	music := &music.Music{}

	err := collection.Find(query).One(music)
	return music, err
}

func NewMusicRepository(session *mgo.Session) *MusicRepository {
	return &MusicRepository{session}
}
