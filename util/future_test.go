package util

import (
	"log"
	"testing"
	"time"
)

func TestNewFuture(t *testing.T) {
	NewFuture().ThenAll(func() {
		log.Println("all 1-1")
		time.Sleep(time.Millisecond * 100)
		log.Println("all 1-2")
	}, func() {
		log.Println("all 2-1")
		time.Sleep(time.Millisecond * 100)
		log.Println("all 2-2")
	}).ThenAll(func() {
		log.Println("all 3-1")
		time.Sleep(time.Millisecond * 1000)
		log.Println("all 3-2")
	}).Then(func() {
		time.Sleep(time.Millisecond * 1000)
		log.Println("then 1")
	}).Then(func() {
		time.Sleep(time.Millisecond * 100)
		log.Println("then 2")
	}).Then(func() {
		time.Sleep(time.Millisecond * 10)
		log.Println("then 3")
	}).Do()
}
