package main

import (
	"fmt"
	"lr2/lr2"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	cfg := lr2.Config{
		Broker:     "broker.hivemq.com:1883",
		ClientID:   fmt.Sprintf("go-client-%d", time.Now().UnixNano()),
		StudentNum: 507057, // Ваш номер!
	}

	cfg.Topics = []string{
		fmt.Sprintf("ITMO/Student%d/Value1", cfg.StudentNum),
		fmt.Sprintf("ITMO/Student%d/Value2", cfg.StudentNum),
		fmt.Sprintf("ITMO/Student%d/Value3", cfg.StudentNum),
	}

	// Подключение
	client, err := lr2.NewMQTTClient(cfg)
	if err != nil {
		fmt.Printf("✗ Ошибка: %v\n", err)
		return
	}
	defer client.Disconnect()

	// Пример публикации (если нужно отправить данные)
	// err = client.Publish(cfg.Topics[0], "Hello from Go!")
	// if err != nil {
	//     fmt.Printf("✗ Ошибка публикации: %v\n", err)
	// }

	fmt.Println("\n🔄 Ожидание сообщений... (Ctrl+C для выхода)")

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	fmt.Println("\n👋 Завершение работы...")
}
