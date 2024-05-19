package main

import (
	"log"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"
)

func main() {
	// 指定单线程运行
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()
	log.Println("current pid ", os.Getpid())

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT)

	file, err := os.Create("/tmp/tracee.txt")
	if err != nil {
		log.Fatalf("failed in open: %v", err)
	}
	defer file.Close()
Loop:
	for {
		select {
		case <-sig:
			log.Println("Interrupt")
			break Loop
		default:
			if err = file.Truncate(0); err != nil {
				log.Fatalf("failed in truncate: %v", err)
			}
			tick := time.Now()
			if _, err := file.Write([]byte(tick.Format("2006/01/02T15:04:05\n"))); err != nil {
				log.Fatalf("failed in write: %v", err)
			}
			log.Printf("write cost: %s", time.Since(tick))
			// n, err := file.Read(readBuf)
			reads, err := os.ReadFile("/tmp/tracee.txt")
			if err != nil {
				log.Fatalf("failed in read: %v", err)
			}
			log.Printf("file content: %s", reads)
		}
		time.Sleep(2 * time.Second)
	}
}
