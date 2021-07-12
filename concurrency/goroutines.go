package main

//https://yourbasic.org/golang/measure-execution-time/
import (
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
	"time"
)

const counter int = 1000

func main() {

	var wg sync.WaitGroup //declare the sync.WaitGroup

	// If the file doesn't exist, create it, or append to the file
	f, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	start := time.Now()

	sequence(counter)
	//fmt.Println(time.Since(start).Nanoseconds())
	//fmt.Println("sequence duration:", time.Since(start).Microseconds())
	duration := time.Since(start).Milliseconds()
	fmt.Println("sequence duration:", duration)
	//fmt.Println(time.Since(start).Seconds())
	rec := []byte("sequence: " + strconv.Itoa(counter) + "|" + strconv.FormatInt(duration, 10) + "\n")
	if _, err := f.Write([]byte(rec)); err != nil {
		f.Close() // ignore error; Write error takes precedence
		log.Fatal(err)
	}

	start = time.Now()
	parallel(counter, &wg)
	wg.Wait()
	duration = time.Since(start).Milliseconds()
	fmt.Println("parallel duration:", duration)

	rec = []byte("goroutines: " + strconv.Itoa(counter) + "|" + strconv.FormatInt(duration, 10) + "\n")
	if _, err := f.Write([]byte(rec)); err != nil {
		f.Close() // ignore error; Write error takes precedence
		log.Fatal(err)
	}

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}

}

func sequence(n int) {

	fmt.Println("sequence start")
	for i := 1; i <= n; i++ {
		//	fmt.Printf("sleep %v: started for one second\n", i)
		time.Sleep(time.Second * 1)
		//	fmt.Printf("sleep %v: ended\n", i)
	}

}

func parallel(n int, wg *sync.WaitGroup) {

	fmt.Println("parallel start")

	for i := 1; i <= n; i++ {
		wg.Add(1)
		go sleep(wg, i)
		//fmt.Println(i)
	}

	wg.Wait()
	fmt.Println("parallel end")
}

func sleep(wg *sync.WaitGroup, id int) {
	defer wg.Done()
	//	fmt.Printf("sleep %v: started for one second\n", id)
	time.Sleep(time.Second * 1)
	//	fmt.Printf("sleep %v: ended\n", id)
}
