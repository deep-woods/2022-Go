package Database

/* What's changed: refactoring in /Database/DatabaseConnection

The refactoring was done because:
- we want to use the same instance of the DatabaseConnection throughout this application in every service.
- This makes the application more modular.

This is a bit of abstraction, which is way better than using boiler plates and random variable decorations.
*/

import (
	"log"

	"github.com/gocql/gocql"
)

type DBconnection struct {
	cluster *gocql.ClusterConfig
	session *gocql.Session
}

var (
	IP string = "172.17.0.2"
	//	connection DBconnection
)

func SetupDBConnection() *DBconnection {
	cluster := gocql.NewCluster(IP)
	cluster.Consistency = gocql.Quorum
	cluster.Keyspace = "rest_api"         // Cassandra - declare keyspace to use.
	session, _ := cluster.CreateSession() // Cassandra - declare session where we actually connect everything.
	return &DBconnection{session: session}
}

func (db *DBconnection) ExecuteQuery(query string, value ...interface{}) {
	if err := db.session.Query(query).Bind(value...).Exec(); err != nil {
		log.Fatal(err)
	}
}

/*

{
    "FirstName": "Coding",
    "LastName": "Forest",
    "Email": "coding@forest.com"
}

*/
