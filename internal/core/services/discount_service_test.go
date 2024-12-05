package services

import (
	"testing"
	"time"

	"github.com/iexcalibur/backend/internal/core/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockDiscountRepository struct {
	mock.Mock
}

func (m *MockDiscountRepository) Create(discount domain.Discount) error {
	args := m.Called(discount)
	return args.Error(0)
}

func (m *MockDiscountRepository) GetByCode(code string) (*domain.Discount, error) {
	args := m.Called(code)
	return args.Get(0).(*domain.Discount), args.Error(1)
}

func (m *MockDiscountRepository) MarkAsUsed(code string) error {
	args := m.Called(code)
	return args.Error(0)
}

func (m *MockDiscountRepository) GetAll() []domain.Discount {
	args := m.Called()
	return args.Get(0).([]domain.Discount)
}

type MockOrderRepository struct {
	mock.Mock
}

func (m *MockOrderRepository) Create(order domain.Order) error {
	args := m.Called(order)
	return args.Error(0)
}

func (m *MockOrderRepository) GetOrderCount() int {
	args := m.Called()
	return args.Int(0)
}

func (m *MockOrderRepository) GetOrderStats() domain.OrderStats {
	args := m.Called()
	return args.Get(0).(domain.OrderStats)
}

func TestDiscountService_GenerateDiscountCode(t *testing.T) {
	discountRepo := new(MockDiscountRepository)
	orderRepo := new(MockOrderRepository)
	service := NewDiscountService(discountRepo, orderRepo, 10)

	t.Run("Generate new discount code when none exists", func(t *testing.T) {
		discountRepo.On("GetAll").Return([]domain.Discount{})
		discountRepo.On("Create", mock.AnythingOfType("domain.Discount")).Return(nil)

		discount, err := service.GenerateDiscountCode()

		assert.NoError(t, err)
		assert.NotNil(t, discount)
		assert.False(t, discount.Used)
		discountRepo.AssertExpectations(t)
	})

	t.Run("Cannot generate when active code exists", func(t *testing.T) {
		activeDiscount := domain.Discount{
			Code: "ACTIVE",
			Used: false,
		}
		discountRepo.On("GetAll").Return([]domain.Discount{activeDiscount})

		_, err := service.GenerateDiscountCode()

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "already an active discount code")
		discountRepo.AssertExpectations(t)
	})
}

func TestDiscountService_ValidateDiscountCode(t *testing.T) {
	discountRepo := new(MockDiscountRepository)
	orderRepo := new(MockOrderRepository)
	service := NewDiscountService(discountRepo, orderRepo, 10)

	validDiscount := &domain.Discount{
		Code:      "VALID10",
		Used:      false,
		CreatedAt: time.Now(),
	}

	t.Run("Valid discount code", func(t *testing.T) {
		discountRepo.On("GetByCode", "VALID10").Return(validDiscount, nil)
		discountRepo.On("GetAll").Return([]domain.Discount{*validDiscount})
		discountRepo.On("MarkAsUsed", "VALID10").Return(nil)

		discount, err := service.ValidateDiscountCode("VALID10")

		assert.NoError(t, err)
		assert.Equal(t, validDiscount.Code, discount.Code)
		discountRepo.AssertExpectations(t)
	})

	t.Run("Used discount code", func(t *testing.T) {
		usedDiscount := &domain.Discount{
			Code: "USED10",
			Used: true,
		}
		discountRepo.On("GetByCode", "USED10").Return(usedDiscount, nil)

		_, err := service.ValidateDiscountCode("USED10")

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "already used")
		discountRepo.AssertExpectations(t)
	})
}
