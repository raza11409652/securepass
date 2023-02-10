package validators

// import (
// 	"github.com/gin-gonic/gin"
// 	"github.com/raza11409652/securepass/common"
// )

// type LoginValidator struct {
// 	User struct {
// 		Email    string `validate:"required"`
// 		Password string `form:"password" json:"password"`
// 	} `json:"user"`
// }

// func (self *LoginValidator) Bind(c *gin.Context) error {
// 	err := common.Bind(c, self)
// 	if err != nil {
// 		return err
// 	}
// 	// self.userModel.Email = self.User.Email
// 	return nil
// }

// // You can put the default value of a Validator here
// func NewLoginValidator() LoginValidator {
// 	loginValidator := LoginValidator{}
// 	return loginValidator
// }
