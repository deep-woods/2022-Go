package Database

import (
	"log"

	"github.com/gocql/gocql"
)

type DBconnection struct {
	cluster *gocql.ClusterConfig
	session *gocql.Session
}

var (
	connection DBconnection
	IP         string = "172.17.0.2"
	/* tried:
	"172.0.0.1:7199"  fail
	"172.17.0.2:7199" fail

	"172.17.0.2:7199", localhost:7199/api/users/getUser  - this works on thunder client!!
	*/
)

func SetupDBConnection() {
	connection.cluster = gocql.NewCluster(IP) // 👈 Docker container port IP.
	connection.cluster.Consistency = gocql.Quorum
	connection.cluster.Keyspace = "rest_api"                   // declare keyspace to use.
	connection.session, _ = connection.cluster.CreateSession() // declare session where we actually connect everything.
}

func (db *DBconnection) ExecuteQuery(query string, value ...interface{}) {
	if err := connection.session.Query(query).Bind(value...).Exec(); err != nil {
		log.Fatal(err)
	}
}

/*

Thunder Client test - Connection was forcibly closed by a peer.

“Connection reset by peer” is the TCP/IP equivalent of slamming the phone back on the hook. It’s more polite than merely not replying, leaving one hanging. But it’s not the FIN-ACK expected of the truly polite TCP/IP.

https://www.howtouselinux.com/post/check-connection-reset-by-peer
*/
