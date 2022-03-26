package Database

import (
	"log"

	"github.com/gocql/gocql"
)

type DBconnection struct {
	cluster *gocql.ClusterConfig
	session *gocql.Session
}

var connection DBconnection

func SetupDBConnection() {
	connection.cluster = gocql.NewCluster("172.17.0.2") // 👈 Docker container port IP.
	connection.cluster.Consistency = gocql.Quorum
	connection.cluster.Keyspace = "rest_api"                   // declare keyspace to use.
	connection.session, _ = connection.cluster.CreateSession() // declare session where we actually connect everything.
}

func ExecuteQuery(query string, value ...interface{}) {
	if err := connection.session.Query(query).Bind(value...).Exec(); err != nil {
		log.Fatal(err)
	}
}
