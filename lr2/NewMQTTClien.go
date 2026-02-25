package lr2

import (
	"fmt"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type Config struct {
	Broker     string
	ClientID   string
	StudentNum int
	Topics     []string
}

// MQTTClient — обёртка над MQTT клиентом
type MQTTClient struct {
	client mqtt.Client
	config Config
}

// NewMQTTClient создаёт и подключает клиент
func NewMQTTClient(cfg Config) (*MQTTClient, error) {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s", cfg.Broker))
	opts.SetClientID(cfg.ClientID)
	opts.SetAutoReconnect(true)
	opts.SetConnectTimeout(10 * time.Second)
	opts.SetKeepAlive(30 * time.Second)

	// Обработчик при успешном подключении
	opts.SetOnConnectHandler(func(c mqtt.Client) {
		fmt.Println("✓ Подключено к брокеру")

		// Подписка на топики после подключения
		for _, topic := range cfg.Topics {
			token := c.Subscribe(topic, 1, messageHandler)
			token.Wait()
			if token.Error() != nil {
				fmt.Printf("✗ Ошибка подписки на %s: %v\n", topic, token.Error())
			} else {
				fmt.Printf("✓ Подписан на: %s\n", topic)
			}
		}
	})

	// Обработчик потери соединения
	opts.SetConnectionLostHandler(func(c mqtt.Client, err error) {
		fmt.Printf("✗ Потеряно соединение: %v\n", err)
	})

	// Создание клиента
	client := mqtt.NewClient(opts)

	// Подключение
	token := client.Connect()
	token.Wait()
	if token.Error() != nil {
		return nil, fmt.Errorf("Ошибка подключения: %w", token.Error())
	}

	return &MQTTClient{client: client, config: cfg}, nil
}

// messageHandler — обработка входящих сообщений
func messageHandler(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("\n📨 Получено:\n")
	fmt.Printf("   Топик: %s\n", msg.Topic())
	fmt.Printf("   Payload: %s\n", string(msg.Payload()))
	fmt.Printf("   QoS: %d, Retained: %v\n", msg.Qos(), msg.Retained())
}

// Publish отправляет сообщение в топик
func (m *MQTTClient) Publish(topic string, payload interface{}) error {
	token := m.client.Publish(topic, 1, false, payload)
	token.Wait()
	return token.Error()
}

// Disconnect закрывает соединение
func (m *MQTTClient) Disconnect() {
	m.client.Disconnect(250)
	fmt.Println("✓ Отключено от брокера")
}

// IsConnected проверяет статус соединения
func (m *MQTTClient) IsConnected() bool {
	return m.client.IsConnected()
}
