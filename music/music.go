package music


type Music struct {
	Id      string `bson:"_id" json:"id"`
	Name    string `bson:"name"`
	Artist  string `bson:artist`
}

