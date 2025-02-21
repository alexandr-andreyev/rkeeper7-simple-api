package rest

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

// Генерация временного кода для номера карты
func (h Handler) generateTempCodeFromCard(c *fiber.Ctx) error {
	tempCode, err := h.services.CRMService.GenerateTempCodeFromCard("4444", 5*time.Minute)
	if err != nil {
		return err
	}
	return c.JSON(fiber.Map{"tempCode": tempCode})
}

// Получить номер карты по временному коду
func (h Handler) getCardByTempCode(c *fiber.Ctx) error {
	card, err := h.services.CRMService.GetCardByTempCode("4444")
	if err != nil {
		return err
	}
	return c.JSON(fiber.Map{"card": card})
}
