package services

import (
	"crypto/rand"
	"errors"
	"fmt"
	"sync"
	"time"
)

type CRMService struct {
}

func NewCrmService() *CRMService {
	return &CRMService{}
}

type tempCode struct {
	cardNumber string
	expiry     time.Time
}

var (
	codeStorage = make(map[string]tempCode)
	mu          sync.Mutex
)

// Генерация случайного 3-значного числа
func generateRandomPart() string {
	b := make([]byte, 1)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("%03d", int(b[0])%900+100) // Диапазон 100-999
}

func (s CRMService) GenerateTempCodeFromCard(cardNumber string, ttl time.Duration) (string, error) {
	mu.Lock()
	defer mu.Unlock()
	part1 := generateRandomPart()
	part2 := generateRandomPart()
	code := fmt.Sprintf("%s-%s", part1, part2)

	// Сохраняем код с временем истечения
	codeStorage[code] = tempCode{
		cardNumber: cardNumber,
		expiry:     time.Now().Add(ttl),
	}

	// Удаление кода через TTL
	time.AfterFunc(ttl, func() {
		mu.Lock()
		delete(codeStorage, code)
		mu.Unlock()
	})

	return code, nil
}

func (s CRMService) GetCardByTempCode(tempCode string) (string, error) {
	mu.Lock()
	defer mu.Unlock()

	data, exists := codeStorage[tempCode]
	if !exists || time.Now().After(data.expiry) {
		return "", errors.New("temp code not found or expired")
	}
	return data.cardNumber, nil
}
