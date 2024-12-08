package services

import (
	"math/rand"
	"time"

	"github.com/iexcalibur/backend/internal/core/domain"
)

type promoService struct {
	promoRepo domain.PromoRepository
}

func NewPromoService(promoRepo domain.PromoRepository) domain.PromoService {
	return &promoService{
		promoRepo: promoRepo,
	}
}

func (s *promoService) GenerateCode(rate float64) (*domain.PromoCode, error) {
	// Generate random code: PROMO followed by 6 random characters
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	code := "PROMO"

	// Add 6 random characters
	b := make([]byte, 6)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	code += string(b)

	promo := &domain.PromoCode{
		Code:         code,
		DiscountRate: rate,
		GeneratedAt:  time.Now(),
	}

	err := s.promoRepo.Save(promo)
	if err != nil {
		return nil, err
	}

	return promo, nil
}

func (s *promoService) ValidateCode(code string) (*domain.PromoCode, error) {
	return s.promoRepo.FindByCode(code)
}

func (s *promoService) GetActiveCode() (*domain.PromoCode, error) {
	return s.promoRepo.FindActive()
}
