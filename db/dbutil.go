package db

import (
	"gopkg.in/mgo.v2"
	//"gopkg.in/mgo.v2/bson"
)

func Save(beam interface{}) (e error) {

	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("test").C("goods")
	err = c.Insert(&beam)
	if err != nil {
		return err
	}

	return nil
}
