package apps

import (
	"fmt"
	"os"
	"strings"
	"sync"
)

type HashApp struct {
	filesDir string
}

func NewHashApp(filesDir string) *HashApp {
	return &HashApp{
		filesDir: filesDir,
	}
}

type resultFile struct {
	FilePath string `json:"file_path"`
	Found    bool   `json:"found"`
	Err      error  `json:"err"`
}

type resultJson struct {
	Results []resultFile `json:"results"`
}

func newResultJson() *resultJson {
	return &resultJson{
		Results: []resultFile{},
	}
}

func (rj *resultJson) add(res resultFile) {
	rj.Results = append(rj.Results, res)
}

func (h HashApp) FindHash(hashStr string) (interface{}, error) {
	var wg sync.WaitGroup
	sema := newSemaphore()
	resChan := make(chan resultFile, 10)
	resJson := newResultJson()
	for i := 1; i <= 10; i++ {
		// acquire resources
		sema.Acquire()
		wg.Add(1)

		// process
		filePath := fmt.Sprintf("%s/file_%d.txt", h.filesDir, i)
		go findInFile(filePath, hashStr, resChan, sema, &wg)
	}
	wg.Wait()
	close(resChan)
	for res := range resChan {
		resJson.add(res)
	}

	return resJson, nil
}

func findInFile(filePath, hash string, resultChan chan resultFile, sema *semaphore, wg *sync.WaitGroup) {
	// release semaphore
	defer sema.Release()

	// release waitgroup
	defer wg.Done()

	res := resultFile{
		FilePath: filePath,
	}

	dat, err := os.ReadFile(filePath)
	if err != nil {
		res.Err = err
		resultChan <- res
		return
	}
	res.Found = strings.Contains(string(dat), hash)
	resultChan <- res
	return
}
