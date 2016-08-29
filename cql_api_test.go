package gocore

import (
	"log"
	"testing"
)

func TestCQLApi(t *testing.T) {
	log.Println("Start update set test...")
	runParallelUpdate()
	log.Println("Start query test...")
	runCQLQuery()
}
