package gocore

import (
	"log"
	"strconv"

	"time"

	"github.com/gocql/gocql"
)

func runParallelUpdate() {
	for i := 0; i < 10; i++ {
		go runCQLUpdate(i)
	}

	time.Sleep(time.Second * 5)
}

func runCQLUpdate(key int) {
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

	var ts int64
	now := time.Now()
	millis := now.UnixNano() / 1000000

	ts = millis
	log.Println("Current milliseconds", ts)

	qry := "UPDATE table_ts_changelog SET value = value + {'change log cn="
	qry = qry + strconv.FormatInt(ts, 10) + "-" + strconv.Itoa(key) + "'} WHERE pid = ? AND key_ts = ?"
	if err := session.Query(qry, 1, ts).Exec(); err != nil {
		//"UPDATE table_ts_changelog SET value = value + {'change entry'} WHERE pid = ? AND key_ts = ?", 1, ts).Exec(); err != nil {
		log.Fatal(err)
	}

}
