package pingchecker

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

type PingResult struct {
	URL        string  `json:"url"`
	PingMs     float64 `json:"ping_ms"`
	Success    bool    `json:"success"`
	StatusCode int     `json:"status_code"`
	Error      string  `json:"error,omitempty"`
}

func PingChecker(url string, timeout time.Duration) PingResult {
	result := PingResult{
		URL:     url,
		Success: false,
	}

	// Валидация URL
	if strings.TrimSpace(url) == "" {
		result.Error = "пустой URL"
		return result
	}

	client := &http.Client{Timeout: timeout}
	start := time.Now()

	response, err := client.Get(url)

	// Всегда измеряем пинг (даже при ошибке)
	elapsed := time.Since(start)
	result.PingMs = float64(elapsed.Nanoseconds()) / 1e6 // Конвертация в мс

	if err != nil {
		// Ошибка сети, пинг есть, но статуса нет
		result.Error = err.Error()
		return result
	}
	defer response.Body.Close()

	// Сервер ответил, есть статус код
	result.StatusCode = response.StatusCode
	result.Success = response.StatusCode >= 200 && response.StatusCode < 300
	// Error остаётся пустым (это не ошибка сети)

	return result
}

func PingUrlsFromFile(path string, timeout time.Duration) ([]PingResult, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	results := []PingResult{}
	var resultsch = make(chan PingResult, 10)
	scanner := bufio.NewScanner(file)

	wg := sync.WaitGroup{}
	const numWorkers = 2
	chJobs := make(chan string, 10)

	//Сначала создаем количество работабщих воркиров, в java это симафор
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		//Отдельная рутина(в цикле задаётся их количество), которая читает задачи поступающие chJobs и отправляет результаты в resultsch
		go func() {
			defer wg.Done()
			for url := range chJobs {
				fmt.Println("Worker", i, "started with url:", url)
				resultsch <- PingChecker(url, timeout)
			}
		}()
	}

	go func() {
		for scanner.Scan() {
			url := strings.TrimSpace(scanner.Text())

			if url == "" || strings.HasPrefix(url, "#") {
				continue
			}
			chJobs <- url
		}
		close(chJobs)
	}()

	go func() {
		wg.Wait()
		close(resultsch)
	}()

	for result := range resultsch {
		results = append(results, result)
	}

	if err := scanner.Err(); err != nil {
		return results, fmt.Errorf("ошибка чтения файла: %w", err)
	}

	return results, nil
}
