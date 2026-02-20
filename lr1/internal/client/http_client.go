package client

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"lr1/lr1/internal/config"
	"lr1/lr1/internal/model"
	"net"
	"net/url"
	"time"
)

var (
	ErrConnectionFailed = errors.New("connection failed")
	ErrRequestFailed    = errors.New("request failed")
	ErrParseFailed      = errors.New("failed to parse response")
)

func Get() (*model.ServerResponse, error) {
	cfg := config.Load()

	if err := cfg.Validate(); err != nil {
		return nil, fmt.Errorf("invalid config: %w", err)
	}

	params := url.Values{}
	params.Add("student", cfg.StudentID)
	params.Add("wday", cfg.WeekDay)

	request := fmt.Sprintf(
		"GET /http_example/give_me_five?%s HTTP/1.0\r\n"+
			"REQUEST_AGENT: ITMO student\r\n"+
			"COURSE: Net Protocols\r\n"+
			"\r\n",
		params.Encode(),
	)

	conn, err := net.DialTimeout("tcp", cfg.ServerAddress(), time.Duration(cfg.Timeout)*time.Second)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrConnectionFailed, err)
	}
	defer conn.Close()

	_, err = conn.Write([]byte(request))
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrRequestFailed, err)
	}

	reader := bufio.NewReader(conn)

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			return nil, fmt.Errorf("%w: %v", ErrRequestFailed, err)
		}
		if line == "\r\n" {
			break
		}
	}

	var response model.ServerResponse
	if err := json.NewDecoder(reader).Decode(&response); err != nil {
		return nil, fmt.Errorf("%w: %v", ErrParseFailed, err)
	}

	if !response.IsValid() {
		return nil, ErrParseFailed
	}

	return &response, nil
}
