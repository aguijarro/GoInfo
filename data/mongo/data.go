package mongo

import (
        "gopkg.in/mgo.v2"
)

func db()(*mgo.Session, error) {
        session, err := mgo.Dial("localhost:27017")
        if err != nil {
                panic(err)
        }
		return session, err
}		
