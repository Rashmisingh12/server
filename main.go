package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	// "strings"
	"time"

	// "log"
	"os"
)

func main() {
	csvfile := flag.String("csv", "./problem.csv", "a csv file inthe format of 'ques' ans ")
	timelimit:=flag.Int("limit",10,"the time limit for quiz in second")
	flag.Parse()
	file, err := os.Open(*csvfile)
	if err != nil {
		fmt.Printf("failed to open:%s", *csvfile)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	s:=make(map[string]string)
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		s[record[0]]=record[1]
		

	}
	timer:=time.NewTimer(time.Duration(*timelimit)*time.Second)
	
	var correctans int
	var attempt int
	wrong:=attempt-correctans
	
	for k,v:=range s{
		fmt.Println(k)
		answer:=make(chan string)
		go func(){
			var ans string
		 fmt.Scanf("%s",&ans)
		 answer<-ans
		}()
		select{
		case<-timer.C:
			fmt.Printf("%d\n %d\n%d\n",correctans,attempt,wrong)
			return 
		case ans:= <-answer:
		 
		if ans==v{
			correctans++
		 }
				attempt++
		 

	}
}
         
   fmt.Printf("%d\n %d\n%d\n",correctans,attempt,wrong)
   
} 

// func a(){
// 	timeout := time.After(3 * time.Second)
// }
// // func TestWithTimeOut(t *testing.T) {
    
//     done := make(chan bool)
//     go func() {
//         // do your testing
//         time.Sleep(5 * time.Second)
//         done <- true
//     }()

//     select {
//     case <-timeout:
//         t.Fatal("Test didn't finish in time")
//     case <-done:
//     }
// }


