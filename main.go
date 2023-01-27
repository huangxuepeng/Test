// package main

// import (
// 	"fmt"
// 	"math/rand"
// )

// type Job struct {
// 	Id     int
// 	RanNum int
// }

// type Result struct {
// 	job *Job
// 	sum int
// }

// func main() {
// 	jobChan, resultChan := make(chan *Job, 128), make(chan *Result, 128)

// 	// 创建工作池
// 	createPool(10000, jobChan, resultChan)

// 	// 打印协程
// 	go func(resultChan chan *Result) {
// 		for result := range resultChan {
// 			fmt.Printf("id=%d, randNum: %d, result: %d\n", result.job.Id, result.job.RanNum, result.sum)
// 		}
// 	}(resultChan)

// 	var id int
// 	for {
// 		id++
// 		r_num := rand.Int()
// 		job := &Job{
// 			Id:     id,
// 			RanNum: r_num,
// 		}
// 		jobChan <- job
// 	}
// }

// func createPool(num int, jobChan chan *Job, resultChan chan *Result) {

// 	for i := 0; i < num; i++ {
// 		go func(jobChan chan *Job, resultChan chan *Result) {

// 			// 遍历相加
// 			for job := range jobChan {
// 				r_num := job.RanNum
// 				var sum int
// 				for r_num != 0 {
// 					tmp := r_num % 10
// 					sum += tmp
// 					r_num /= 10
// 				}
// 				r := &Result{
// 					job: job,
// 					sum: sum,
// 				}

// 				// 扔进管道
// 				resultChan <- r
// 			}

//			}(jobChan, resultChan)
//		}
//	}
package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(1 * time.Second)

	i := 0

	go func() {
		for {
			i++
			fmt.Println(<-ticker.C)
			for i == 5 {
				ticker.Stop()
				fmt.Println("结束")
			}
		}
	}()
	for {

	}
}
