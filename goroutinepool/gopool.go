package main

import (
	"fmt"
	"time"
)

type GoPool struct {
	Queue 		chan func() error
	RuntimeNo 	int

	Total 		int

	Result 		chan error
	FinishCallback func()
}

func (gp *GoPool) Init(runtimeNo, total int) {
	gp.RuntimeNo = runtimeNo
	gp.Total = total
	gp.Queue = make(chan func() error, total)
	gp.Result = make(chan error, total)
}

func (gp *GoPool) Start() {

	// 开启RuntimeNo个goroutine
	for i := 0; i < gp.RuntimeNo; i++ {
		go func() {
			for {
				task, ok := <- gp.Queue
				if !ok {
					//
					break
				}

				// 执行目标任务
				err := task()
				gp.Result <- err
			}
		}()
	}

	// 获取每个任务的处理结果
	for j := 0; j < gp.RuntimeNo; j++ {
		res, ok := <- gp.Result
		if !ok {
			break
		}
		if res != nil {
			fmt.Println(res)
		}
	}

	// 结束后回调
	if gp.FinishCallback != nil {
		gp.FinishCallback()
	}
}

func (gp *GoPool) Stop() {
	close(gp.Queue)
	close(gp.Result)
}

func (gp *GoPool) AddTask(task func() error)  {
	gp.Queue <- task
}

func (gp *GoPool) SetFinishCallback(fun func()) {
	gp.FinishCallback = fun
}

func main() {
	var gp GoPool
	urls := []string{"192.168.168.154", "192.168.152.174", "192.168.152.183","192.168.168.151", "192.168.152.172", "192.168.152.184","192.168.168.158", "192.168.152.170", "192.168.152.181"}

	gp.Init(3, len(urls))
	for i := range urls {
		url := urls[i]
		gp.AddTask(func() error {
			return download(url)
		})
	}

	gp.SetFinishCallback(downloadFin)
	gp.Start()
	gp.Stop()
	//
	//gp.AddTask(func() error {
	//	return download("1234565")
	//})
}

func download(url string) error {
	fmt.Printf("Downloading from %s\n", url)
	time.Sleep(3 * time.Second)
	return nil
}

func downloadFin() {
	fmt.Println("Download Finish.")
}