package cassandra

import (
	"fmt"

	"github.com/gocql/gocql"
)

var (
	session *gocql.Session
)

func init() {
	// connec to cassandra cluster:
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "oauth"
	cluster.Consistency = gocql.Quorum
	cluster.ProtoVersion = 4

	var err error
	if session, err = cluster.CreateSession(); err != nil {
		panic(err)
	}
	fmt.Println("connection successfully")
	defer session.Close()
}

func GetSession() *gocql.Session {
	return session
}
