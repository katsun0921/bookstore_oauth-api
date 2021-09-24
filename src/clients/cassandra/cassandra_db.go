package cassandra

import (
	"github.com/gocql/gocql"
)

var (
	session *gocql.Session
)

func init() {
	/* Connect Cassandra Cluster:
	The example assumes the following CQL was used to setup the keyspace:
	create keyspace example with replication = { 'class' : 'SimpleStrategy', 'replication_factor' : 1 };
	create table example.tweet(timeline text, id UUID, text text, PRIMARY KEY(id));
	create index on example.tweet(timeline);
	*/
	cluster := gocql.NewCluster("localhost:9042")
	cluster.Keyspace = "oauth"
	cluster.Consistency = gocql.Quorum

  var err error
  if session, err = cluster.CreateSession(); err != nil {
    panic(err)
  }
}

func GetSession() *gocql.Session {
	return session
}
