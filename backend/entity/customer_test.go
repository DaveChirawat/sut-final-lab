package entity

// import(
// 	"testing"
// 	"github.com/asaskevich/govalidator"
// 	. "github.com/onsi/gomega"
// 	"gorm.io/gorm"

// )
// type Customer struct {
// 	gorm.Model
// 	Name string // ต้องไม่เป็นค่าว่าง
// 	Email string
// 	CustomerID string // รหัสลูกค้าขึ8นต้นด้วย L หรือ M หรือ H แล้วตามด้วยตัวเลขจํานวน 7 ตัว
//    }
//    func TestCustomerValidator(t *testing.T) {
// 	g := NewGomegaWithT(t)

// 	t.Run("Check", func(t *testing.T)) (
// 		d := Customer{
// 			Name:	"Dave"
// 			Email: 	"defza@gmail.com"
// 			CustomerID: 	"B6308308"
// 		}
// 		ok, err := validator.ValidateStruct(d)

// 		g.Expact(ok).NotTo(BeTrue())
// 		g.Expact(err)ToNot(BeNil())
// 		g.Expact(err.Error()).To(Equal(""))

// 	)

//    }
