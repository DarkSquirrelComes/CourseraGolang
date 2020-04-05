package main

import (
	"fmt"
	"sort"
	"strconv"
	"sync"
	"time"
)

func ExecutePipeline(jobs ...job) {
	wg := &sync.WaitGroup{}
	var chans = make([]chan interface{}, len(jobs)+1)
	for i := range chans {
		chans[i] = make(chan interface{})
	}
	//mu := &sync.Mutex{} //may be it fixes RACE
	for i := range jobs {
		//fmt.Printf("function %d go!\n", i)
		//mu.Lock()
		chIn, chOut, f := chans[i], chans[i+1], jobs[i]
		//mu.Unlock()
		wg.Add(1)
		go func(wg *sync.WaitGroup, in, out chan interface{}, f func(chan interface{}, chan interface{})) {
			defer wg.Done()
			defer close(out)
			f(in, out)
		}(wg, chIn, chOut, f)
	}
	//time.Sleep(time.Millisecond)

	// time.Sleep(time.Second)
	wg.Wait()
	// return
}

func SingleHash(in, out chan interface{}) {
	wgBig := &sync.WaitGroup{}
	quotaCh := make(chan struct{}, 1)
	for data := range in {
		wgBig.Add(1)
		go func(in, out chan interface{}, d int, wgBig *sync.WaitGroup, quotaCh chan struct{}) {
			defer wgBig.Done()
			start := time.Now()
			wg := &sync.WaitGroup{}
			wg.Add(2)
			fmt.Println("SingleHash got", d)
			var resCM string
			var resC string
			go func(resCM *string, wg *sync.WaitGroup, quotaCh chan struct{}) {
				defer wg.Done()
				quotaCh <- struct{}{}
				dSMD5 := DataSignerMd5(strconv.Itoa(d))
				<-quotaCh
				*resCM = DataSignerCrc32(dSMD5)
			}(&resCM, wg, quotaCh)
			go func(resC *string, wg *sync.WaitGroup) {
				defer wg.Done()
				*resC = DataSignerCrc32(strconv.Itoa(d))
			}(&resC, wg)
			wg.Wait()
			fmt.Println("SingleHash send", d, ":", resC+"~"+resCM, time.Since(start))
			out <- resC + "~" + resCM
		}(in, out, data.(int), wgBig, quotaCh)
	}
	wgBig.Wait()
}

func MultiHash(in, out chan interface{}) {
	wgBig := &sync.WaitGroup{}
	for data := range in {
		wgBig.Add(1)
		go func(in, out chan interface{}, d string, wgBig *sync.WaitGroup) {
			defer wgBig.Done()
			fmt.Println("MultiHash got", d)
			hashes := make([]string, 6)
			wg := &sync.WaitGroup{}
			for i := 0; i < 6; i++ {
				wg.Add(1)
				go func(res *string, d string, i int, wg *sync.WaitGroup) {
					defer wg.Done()
					preRes := DataSignerCrc32(strconv.Itoa(i) + d)
					*res = preRes
				}(&hashes[i], d, i, wg)
			}
			wg.Wait()
			var res string
			for _, h := range hashes {
				res += h
			}
			fmt.Println("MultiHash send", res)
			out <- res
		}(in, out, data.(string), wgBig)
	}
	wgBig.Wait()
}

func CombineResults(in, out chan interface{}) {
	var hashes []string
	for data := range in {
		fmt.Println("CombineResults got", data.(string))
		hashes = append(hashes, data.(string))
	}
	fmt.Println("CombineResults got all:", hashes)
	sort.Slice(hashes, func(i, j int) bool {
		return hashes[i] < hashes[j] //check in case of WA
	})
	var res string
	for i, h := range hashes {
		if i != 0 {
			res += "_"
		}
		res += h
	}
	fmt.Println("CombineResults send:", res)
	out <- res
}

func main() {
	return
}
