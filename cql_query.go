package gocore

import (
	"log"

	"github.com/gocql/gocql"
)

func runCQLQuery() {
	log.Println("Welcome")
	// connect to the cluster
	cluster := gocql.NewCluster("10.20.31.17:9292")
	cluster.Keyspace = "keyspace_test"
	cluster.Consistency = gocql.LocalOne
	session, err := cluster.CreateSession()

	if err != nil {
		log.Fatal("Failed to create a session with Cassandra...")
	}

	defer session.Close()

	var key int64
	var value []string
	// list all changes

	iter := session.Query(`SELECT key_ts,value FROM table_ts_changelog`).Iter()
	if iter == nil {
		log.Fatalln("Failed to execute the query")
	}

	for iter.Scan(&key, &value) {
		log.Printf("Key:%v, value:%v\n", key, value)
	}

	if err := iter.Close(); err != nil {
		log.Fatal(err)
	}

	// Update the same user with a new age
	if err := session.Query("UPDATE table_test SET value = value + {'change 8'} WHERE key = 8").Exec(); err != nil {
		log.Fatal(err)
	}
}
