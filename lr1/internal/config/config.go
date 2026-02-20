package config

import (
	"fmt"
	"os"
	"strconv"
	"sync" // Для thread-safe инициализации
)

type Config struct {
	ServerHost string
	ServerPort string
	StudentID  string
	WeekDay    string
	Timeout    int
}

var (
	cfg  *Config
	once sync.Once // Защита от race condition
)

func Load() *Config {
	once.Do(func() { // Thread-safe инициализация
		cfg = &Config{
			ServerHost: getEnv("SERVER_HOST_LR1", "109.167.241.225"),
			ServerPort: getEnv("SERVER_PORT_LR1", "8001"),
			StudentID:  getEnv("STUDENT_ID_LR1", "30"),
			WeekDay:    getEnv("WEEK_DAY_LR1", "5"),
			Timeout:    getEnvInt("TIMEOUT_LR1", 10),
		}
	})
	return cfg
}

func (c *Config) Validate() error {
	if c.StudentID == "" {
		return fmt.Errorf("student ID cannot be empty")
	}
	if c.WeekDay == "" {
		return fmt.Errorf("week day cannot be empty")
	}
	if _, err := strconv.Atoi(c.StudentID); err != nil {
		return fmt.Errorf("student ID must be a number")
	}
	if _, err := strconv.Atoi(c.WeekDay); err != nil {
		return fmt.Errorf("week day must be a number")
	}
	return nil
}

func (c *Config) ServerAddress() string {
	return fmt.Sprintf("%s:%s", c.ServerHost, c.ServerPort)
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intVal, err := strconv.Atoi(value); err == nil {
			return intVal
		}
	}
	return defaultValue
}
