package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/MaestroShifu/concurrent-queue-golang/queue"
)

func main() {
	//validateWithQueue()
	validateWithForLoop()

	/* 	fmt.Println("=== Inicia el pool ===")

	   	workers := queue.NewWorkerPool(2, 5)
	   	for i := 0; i < 20; i++ {
	   		workers.SubmitJob(i)
	   	}

	   	fmt.Println("=== Finaliza el pool ===") */
}

func validateWithForLoop() { // 82.491292ms
	array := []int{}
	for i := 0; i < 50000; i++ {
		array = append(array, i)
	}
	start := time.Now()
	fmt.Println("=== Inicia el array ===")
	for element := range array {
		log.Printf("Data del array %d\n", element)
	}
	fmt.Println("=== Fin el array ===")
	elapsed := time.Since(start)
	fmt.Printf("Time execution %s", elapsed)
}

func validateWithQueue() { // 74.311291ms
	queue := queue.NewQueueBasic(50000)
	for i := 0; i < 50000; i++ {
		queue.Put(i)
	}
	start := time.Now()
	fmt.Println("=== Inicia el queue ===")
	for {
		data, err := queue.Pop()
		if err != nil {
			log.Fatal(err)
			os.Exit(0)
		}
		log.Printf("Data de la cola %d\n", data)
		if queue.IsEmpty() {
			break
		}
	}
	fmt.Println("=== Fin el queue ===")
	elapsed := time.Since(start)
	fmt.Printf("Time execution %s", elapsed)
}
