package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"sync"
	"test_go"
	ping_checker "test_go/pingChecker"
	"time"
)

func main() {

	r, err := ping_checker.PingUrlsFromFile("/home/neket/GolandProjects/protocols/test_go/pingChecker/sites.txt", time.Second*10)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(r)

	WriteFileUser("users.json", test_go.User{Name: "John Doe", Age: 30})
	fmt.Println(ReadFileUser("users.json"))

	WorkerPool()
	WorkerPool2()

	Goroutines()
	Goroutines2()
	Goroutines3()

	fmt.Println(Divide(5, 2))
	fmt.Println(Divide(5, 0))

	account0 := test_go.NewAccount("Nikita")
	fmt.Println(account0)
	account := test_go.NewAccountWithDetails(1, 100, "John Doe")
	fmt.Println(account.GetBalance())
	account.Deposit(20)
	account.Deposit(20)
	fmt.Println(account.GetBalance())
	account2 := account.Deposit2(30)
	account2 = account2.Deposit2(40)
	fmt.Println(account2.GetBalance2())

	PrintArea(test_go.NewCircle(4))
	PrintArea(test_go.NewRectangle(4, 5))

	fmt.Println(frequencyWorld("hello world hello world hello world"))

	a1 := []int{1, 2, 3, 4, 5}
	fmt.Println(a1)
	fmt.Println(slicePowTwo1(a1))
	fmt.Println(a1)
	fmt.Println("________________________")

	a2 := []int{1, 2, 3, 4, 5}
	fmt.Println(a2)
	fmt.Println(slicePowTwo2(&a2))
	fmt.Println(a2)
	fmt.Println("________________________")

	a3 := make([]int, 0)
	a3 = append(a3, 1, 2, 3, 4, 5)
	fmt.Println(a3)
	fmt.Println(slicePowTwo2(&a3))
	fmt.Println(a3)
	fmt.Println("________________________")

}

func WriteFileUser(filename string, user test_go.User) error {

	data, err := json.Marshal(user)
	if err != nil {
		return fmt.Errorf("ошибка сериализации: %w", err)
	}
	return os.WriteFile(filename, data, 0644)
}

func ReadFileUser(filename string) (test_go.User, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return test_go.User{}, fmt.Errorf("ошибка десериализации: %w", err)
	}
	var user test_go.User
	err = json.Unmarshal(data, &user)
	if err != nil {
		return test_go.User{}, fmt.Errorf("ошибка десериализации: %w", err)
	}
	return user, nil
}

func WorkerPool() {
	var chJobs = make(chan int)
	var chResults = make(chan int)

	go func() {
		for job := range chJobs {
			chResults <- job * job
		}
		close(chResults)
	}()

	go func() {
		chJobs <- 1
		chJobs <- 2
		chJobs <- 3
		chJobs <- 4
		close(chJobs)
	}()

	for result := range chResults {
		fmt.Println("Возведение в квадрат: ", result)
	}
}

func WorkerPool2() {
	const numWorkers = 3
	const numJobs = 5
	var wg sync.WaitGroup

	chJobs := make(chan int, numJobs)
	chResults := make(chan int, numJobs)

	// Запускаем несколько воркеров
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for job := range chJobs {
				fmt.Printf("Worker %d обрабатывает job %d\n", w, job)
				//time.Sleep(time.Second * 1)
				chResults <- job * job
			}
		}()
	}

	// Отправляем задания
	for j := 1; j <= numJobs; j++ {
		chJobs <- j
	}
	close(chJobs)

	go func() {
		wg.Wait()
		close(chResults)
	}()

	// Ждём завершения всех воркеров, затем закрываем chResults
	// Собираем результаты
	for r := 1; r <= numJobs; r++ {
		fmt.Println("Результат:", <-chResults)
	}
}

func Goroutines() {
	var wg sync.WaitGroup

	fmt.Println("Goroutines:")
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			time.Sleep(1 * time.Second)
			fmt.Println(i)
		}()
	}
	wg.Wait()
	fmt.Println("Goroutines finished")
}

func Goroutines2() {
	var ch = make(chan bool)

	fmt.Println("Goroutines:")
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Printf("Горутина %d работает\n", i)
			time.Sleep(time.Second)
			ch <- true
		}()
	}

	// Ждём сигналы от всех горутин
	for i := 0; i < 10; i++ {
		<-ch
	}

	fmt.Println("Goroutines finished")
}

func Goroutines3() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
	}()

	for num := range ch {
		fmt.Println("Из канала получено: ", num)
	}
	fmt.Println("Канал закрыт, все данные получены")
}

func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

func PrintArea(shape test_go.Shape) {
	fmt.Println(shape.Area())
}

func frequencyWorld(str string) map[string]int {
	countRepeated := make(map[string]int)
	words := strings.Split(str, " ")
	for _, k := range words {
		countRepeated[string(k)]++
	}
	return countRepeated
}

func slicePowTwo1(intSlice []int) []int {
	for i := range intSlice {
		intSlice[i] = intSlice[i] * 2
	}
	return intSlice
}

func slicePowTwo2(intSlice *[]int) *[]int {
	for i := range *intSlice {
		//fmt.Println((*intSlice)[i])
		(*intSlice)[i] = (*intSlice)[i] * 2
	}
	return intSlice
}
