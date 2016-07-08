package main

import (
	"fmt"
	"sort"
	"strconv"
	"time"

	fb "github.com/huandu/facebook"
)

var (
	nameResult = map[string]int{}
)

func main() {
	timeStart := time.Now()
	results := make(chan map[string]int)
	//done := make(chan bool)
	workerNum := mapper(results, 5000000, 5500000)

	//	reducer(results, workerNum, done)
	combiner(workerNum, results)
	// for i := 0; i < workerNum; i++ {
	//	time.Sleep(500 * time.Millisecond)
	//	<-done
	// }
	// for name, count := range nameResult {
	//	fmt.Printf("There are %v people with first name %v", count, name)
	//	fmt.Println("")
	// }
	list := []string{}
	for name, _ := range nameResult {

		list = append(list, name)
	}
	sorted := sortedName{nameMap: nameResult, nameList: list}
	sort.Sort(sorted)
	for _, name := range sorted.nameList {

		fmt.Printf("There are %v people with first name %v", sorted.nameMap[name], name)

		fmt.Println("")
	}
	duration := time.Since(timeStart)
	fmt.Printf("This takes %v to run", duration)
	fmt.Println(len(sorted.nameList))
}

type sortedName struct {
	nameMap  map[string]int
	nameList []string
}

func (s sortedName) Len() int {
	return len(s.nameMap)
}
func (s sortedName) Swap(i, j int) {
	s.nameList[i], s.nameList[j] = s.nameList[j], s.nameList[i]
}

func (s sortedName) Less(i, j int) bool {
	return s.nameMap[s.nameList[i]] > s.nameMap[s.nameList[j]]
}
func mapper(results chan map[string]int, startId, endId int) int {
	var workerNum int
	for i := startId; i < endId; i = i + 35 {

		workerNum++
		go func(i int, results chan map[string]int) {
			nameMap := make(map[string]int)
			for startId := i; startId < i+35; startId++ {
				uId := "/" + strconv.Itoa(startId)
				//				fmt.Println(uId)
				res, _ := fb.Get(uId, fb.Params{
					"fields": "first_name",
					// "fields":       "bio",
					"access_token": "1156172417738306|-XmlLqK1eD7mBMNWwVxAINpFAFM",
				})
				if res != nil && res["first_name"] != nil {
					name := res["first_name"].(string)
					if val, ok := nameMap[name]; ok {
						nameMap[name] = val + 1
					} else {
						nameMap[name] = 1
					}
				}
			}
			//fmt.Println("one worker finishes")
			results <- nameMap
		}(i, results)

	}
	return workerNum
}

// func reducer(results chan map[string]int, workerNum int, done chan bool) {
// 	var resultNum int
// 	for resultNum < workerNum {
// 		select {

// 		case result := <-results:
// 			resultNum++
// 			for key, val := range result {
// 				if _, ok := nameResult[key]; ok {
// 					nameResult[key] = nameResult[key] + val
// 				} else {
// 					nameResult[key] = val

// 				}
// 			}
// 		case <-time.After(100 * time.Millisecond):
// 		}
// 	}
// }

func combiner(workerNum int, results chan map[string]int) {
	num := workerNum / 5
	processedCount := make([]int, num)
	//	for i:=0;i<num;i++
	done := make(chan bool)
	combinedMap := make(chan map[string]int, num)
	for i := 0; i < num; i++ {
		go func(results chan map[string]int, done chan bool, count *int, combinedMap chan map[string]int) {

			localMap := make(map[string]int)

			for {
				select {
				case result := <-results:
					*count++
					for key, val := range result {
						if _, ok := localMap[key]; ok {
							localMap[key] = localMap[key] + val
						} else {
							localMap[key] = val
						}

					}

				case <-time.After(1 * time.Millisecond):
				case <-done:
					combinedMap <- localMap
					return
				}
			}

		}(results, done, &processedCount[i], combinedMap)
	}

	sum := 0
	for sum < workerNum {
		sum = 0
		for _, val := range processedCount {
			sum = sum + val
		}
fmt.Println(sum,workerNum)
		time.Sleep(100 * time.Millisecond)

	}
	for i := 0; i < num; i++ {
		done <- true
	}
	for i := 0; i < num; i++ {
		reducedMap := <-combinedMap

		for key, val := range reducedMap {
			if _, ok := nameResult[key]; ok {
				nameResult[key] = nameResult[key] + val
			} else {
				nameResult[key] = val

			}
		}
	}
}

// func reducerConcurrentMap(results chan map[string]int, workerNum int, done chan bool) {
//	var resultNum int
//	for resultNum < workerNum {
//		select {

//		case result := <-results:
//			resultNum++
//			go func(result map[string]int, done chan bool) {
//				for key, val := range result {
//					if _, ok := nameResult[key]; ok {
//						nameResult[key] = nameResult[key] + val
//					} else {
//						nameResult[key] = val

//					}
//				}
//				// done <- true
//				fmt.Println("process a result")
//			}(result, done)
//		case <-time.After(100 * time.Millisecond):
//		}
//	}
// }

// res, _ := fb.Get("/538744488", fb.Params{
//	"fields":       "first_name",
//	"access_token": "1156172417738306|-XmlLqK1eD7mBMNWwVxAINpFAFM",
// })
// for key, _ := range res {
//	fmt.Println(key)
// }
// fmt.Println("here is my facebook first name:", res["first_name"])
// res, _ := fb.Get("/search", fb.Params{
//	"access_token": "626057207558688|oIvRVCyD9ZFZjRsWIBAXLsj9Jt0",
//	"type":         "page",
//	"q":            "singapore",
// })

// var items []fb.Result

// err := res.DecodeField("data", &items)

// if err != nil {
//	fmt.Printf("An error has happened %v", err)
//	return
// }

// for _, item := range items {
//	for key, val := range item {
//		fmt.Println(key, val)
//		fmt.Println("")
//	}
// }
