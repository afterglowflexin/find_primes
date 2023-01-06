package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

type ranges []string

func (i *ranges) String() string {
	return fmt.Sprint(*i)
}

func (i *ranges) Set(value string) error {
	*i = append(*i, value)
	return nil
}

var (
	fileName   string
	timeout    string
	rangesFlag ranges
	start      time.Time
)

func init() {
	flag.StringVar(&fileName, "file", "testfile.txt", "file for output")
	flag.StringVar(&timeout, "timeout", "10", "timeout for execution")
	flag.Var(&rangesFlag, "range", "range of numbers to find prime number")
	start = time.Now()
}

func main() {
	flag.Parse()

	file, err := os.Create(fileName)
	if err != nil {
		log.Println("Unable to create file:", err)
		os.Exit(1)
	}
	defer file.Close()

	duration, err := time.ParseDuration(timeout + "s")
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()

	ch := make(chan string)
	ch2 := make(chan string)

	for _, x := range rangesFlag {
		go func(x string) {
			num1, num2 := getRange(x)
			res := printPrimeNumbers(num1, num2)
			ch <- res
		}(x)
	}

	go func() {
		for i := 0; i < len(rangesFlag); i++ {
			res := <-ch
			file.WriteString(res)
		}
		ch2 <- ""
	}()

	select {
	case <-ctx.Done():
		log.Println(ctx.Err())
	case <-ch2:
		log.Println("Done")
	}

	fmt.Println(time.Since(start))
}

func getRange(s string) (int, int) {
	rng := strings.Split(s, ":")
	num1, err := strconv.Atoi(rng[0])

	if err != nil {
		log.Fatal("Wrong range format")
	}

	num2, err := strconv.Atoi(rng[1])
	if err != nil {
		log.Fatal("Wrong range format")
	}
	return num1, num2
}

func printPrimeNumbers(num1, num2 int) string {
	res := ""
	if num1 < 2 || num2 < 2 {
		log.Fatal("Numbers must be greater than 2.")
		return res
	}

	if num1 > num2 {
		num1, num2 = num2, num1
	}

	for num1 <= num2 {
		isPrime := true
		for i := 2; i <= int(math.Sqrt(float64(num1))); i++ {
			if num1%i == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			res = res + " " + strconv.Itoa(num1)
		}
		num1++
	}
	return res
}
