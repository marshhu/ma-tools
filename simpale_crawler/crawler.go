package simpale_crawler

import (
	"fmt"
	"net/http"
	"regexp"
	"sync"
)

//工作信息
type JobInfo struct {
	PlatformName string
	PlatformUrl  string
	Url          string
	FindContent  string
}

//工作结果
type WorkResult struct {
	Job        JobInfo
	HttpStatus int
	Content    []byte
}

//处理结果
type HandleResult struct {
	Job        JobInfo
	HttpStatus int
	IsFind     bool
}

//处理方法
type HandleFunc func(workResult *WorkResult) *HandleResult //定义处理函数

func CrawlerPlatform(jobInfos []JobInfo) {
	jobs := make(chan JobInfo, 10)           //工作job
	workResults := make(chan WorkResult, 10) //工作结果
	go createJobs(jobInfos, jobs)            //创建
	done := make(chan bool)
	go handleResult(workResults, findContent, done) //处理结果
	numOfWorkers := 20
	createWorkerPool(numOfWorkers, jobs, workResults) //创建工作池
	<-done
}

func createJobs(jobInfos []JobInfo, jobs chan JobInfo) {
	for _, job := range jobInfos {
		jobs <- job
	}
	close(jobs)
}

func createWorkerPool(numOfWorkers int, jobs chan JobInfo, workResults chan WorkResult) {
	var wg sync.WaitGroup
	for i := 1; i <= numOfWorkers; i++ {
		wg.Add(1)
		go worker(i, jobs, workResults, &wg)
	}
	wg.Wait()
	close(workResults)
}

func worker(id int, jobs chan JobInfo, workResults chan WorkResult, wg *sync.WaitGroup) {
	fmt.Printf("worker %d starting\n", id)
	defer wg.Done()
	for job := range jobs {
		httpStatus, content, _ := Fetcher(job.Url, 5)
		result := WorkResult{Job: job, HttpStatus: httpStatus, Content: content}
		workResults <- result
	}
}

func handleResult(workResults chan WorkResult, handelFunc HandleFunc, done chan bool) {
	for workResult := range workResults {
		handleResult := handelFunc(&workResult)
		fmt.Printf("%s平台，http状态为:%d,查找结果为：%v\n", handleResult.Job.PlatformName, handleResult.HttpStatus, handleResult.IsFind)
	}
	done <- true
}

func findContent(workResult *WorkResult) *HandleResult {
	handleResult := HandleResult{Job: workResult.Job, HttpStatus: workResult.HttpStatus}
	if workResult.HttpStatus != http.StatusOK {
		handleResult.IsFind = false
		return &handleResult
	}

	//查找内容
	re := regexp.MustCompile(workResult.Job.FindContent)
	findStr := re.FindString(string(workResult.Content))
	if len(findStr) > 0 {
		handleResult.IsFind = true
		return &handleResult
	}

	handleResult.IsFind = false
	return &handleResult
}
