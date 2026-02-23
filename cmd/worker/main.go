package worker

import (
	"log"
	"time"
)

func main() {
	log.Println("Worker running...")

	for {
		log.Println("Processing job...")
		time.Sleep(10 * time.Second)
	}
}
