package main

import (
	"crypto/tls"
	"fmt"
	"os"
	"log"
	"net"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

// Model
type Package struct {
	Id bson.ObjectId  `bson:"_id,omitempty"`
	FullName      string
	Description   string
	StarsCount    int
	ForksCount    int
	LastUpdatedBy string
}

func main() {

	tlsConfig := &tls.Config{}

	// InsecureSkipVerify controls whether a client verifies the
	// server's certificate chain and host name.
	// If InsecureSkipVerify is true, TLS accepts any certificate
	// presented by the server and any host name in that certificate.
	// In this mode, TLS is susceptible to man-in-the-middle attacks.
	// This should be used only for testing.
	tlsConfig.InsecureSkipVerify = true

	// DialInfo holds options for establishing a session with a MongoDB cluster.
	dialInfo := &mgo.DialInfo{
		Addrs:    []string{"golang-couch.documents.azure.com:10255"}, // Get HOST + PORT
		Timeout:  60 * time.Second,
		Database: "golang-couch",                                                                             // It can be anything
		Username: "golang-couch",                                                                             // Username
		Password: "02lgiHBWQpRndXAp8SfYhOfe25kVO5DgD57ONNEE5C2gtT3C8Wp3CgxGblyxl9dfuANxFwYSgTiRbob5Iar3Ow==", // PASSWORD
	}

	dialInfo.DialServer = func(serverAddress *mgo.ServerAddr) (net.Conn, error) {
		fmt.Println(serverAddress.String());
		connection, err := tls.Dial("tcp", serverAddress.String(), tlsConfig)
		return connection, err

	}

	// Create a session which maintains a pool of socket connections
	// to our MongoDB.
	session, err := mgo.DialWithInfo(dialInfo)

	if err != nil {
		fmt.Printf("Can't connect to mongo, go error %v\n", err)
		os.Exit(1)
	}

	defer session.Close()

	// SetSafe changes the session safety mode.
	// If the safe parameter is nil, the session is put in unsafe mode, and writes become fire-and-forget,
	// without error checking. The unsafe mode is faster since operations won't hold on waiting for a confirmation.
	// http://godoc.org/labix.org/v2/mgo#Session.SetMode.
	session.SetSafe(&mgo.Safe{})

	// get collection
	collection := session.DB("golang-couch").C("package")

	// insert Document in collection
	err = collection.Insert(&Package{
		FullName:"react",
		Description:"A framework for building native apps with React.",
		ForksCount: 11392,
		StarsCount:48794,
		LastUpdatedBy:"shergin",

	})

	if err != nil {
		log.Fatal("Problem inserting data: ", err)
		return
	}

	// Get Document from collection
	result := Package{}
	err = collection.Find(bson.M{"fullname": "react"}).One(&result)
	if err != nil {
		log.Fatal("Error finding record: ", err)
		return
	}

	fmt.Println("Description:", result.Description)

	// update document
	updateQuery := bson.M{"_id": result.Id}
	change := bson.M{"$set": bson.M{"fullname": "react-native"}}
	err = collection.Update(updateQuery, change)
	if err != nil {
		log.Fatal("Error updating record: ", err)
		return
	}

	// delete document
	err = collection.Remove(updateQuery)
	if err != nil {
		log.Fatal("Error deleting record: ", err)
		return
	}
}
