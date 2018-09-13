package main

import (
	"time"

	workerPool "github.com/qianguozheng/go-workerpool"
)

func main() {
	// generate worker to do job
	dispatcher := workerPool.NewDispatcher(3)
	dispatcher.Run()

	// produce job to be done
	producer := workerPool.NewProducer(40)
	producer.Run()

	time.Sleep(time.Second * 100)

}
