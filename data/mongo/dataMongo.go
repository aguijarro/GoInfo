package mongo

import (

    "gopkg.in/mgo.v2"	
    "gopkg.in/mgo.v2/bson"    
)

type Result struct{
	Description string `bson:"Description"`
	ValYear1 int `bson:"ValYear1"`
	ValYear2 int `bson:"ValYear2"`
	ValYear3 int `bson:"ValYear3"`
	ValYear4 int `bson:"ValYear4"`
	ValYear5 int `bson:"ValYear5"`		 
	ValYear6 int `bson:"ValYear6"`		
}



func ShowContentMongo()(results []Result, err error){

		
		db, err := db()
		defer db.Close()
		db.SetMode(mgo.Monotonic, true)

		
		c := db.DB("goinfo").C("data")
		
		var resultsData []Result
		err = c.Find(bson.M{"Description": "de 60 años y más"}).All(&resultsData)
 
		if err != nil {
			panic(err)
		}
			
	return resultsData, err
}
