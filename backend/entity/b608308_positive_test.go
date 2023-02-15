package entity

import (
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	Name       string `valid:"requrl"`
	Email      string
	CustomerID string `valid:"matches(^//d{7})"` // รหัสลูกค้าขึ8นต้นด้วย L หรือ M หรือ H แล้วตามด้วยตัวเลขจํานวน 7 ตัว
}

func TestCustomerValidator(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run("Check Name", func(t *testing.T) {
		d := Customer{
			Name:       "Dave",
			Email:      "defza@gmail.com",
			CustomerID: "16308308",
		}
		ok, err := govalidator.ValidateStruct(d)

		g.Expact(ok).To(BeTrue())
		g.Expact(err).To(BeNil())
		g.Expact(err.Error()).To(Equal(""))

	})

}
