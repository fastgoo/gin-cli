package validator

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func init() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("UsernameValid", usernameValid)
		v.RegisterValidation("UsernameExistValid", usernameExistValid)
		v.RegisterValidation("UsernameNotExistValid", usernameNotExistValid)
		v.RegisterValidation("PassSecurityValid", passSecurityValid)

		v.RegisterValidation("CompanyNameNotExistValid", companyNameNotExistValid)
		v.RegisterValidation("TelValid", telValid)
		v.RegisterValidation("LicensePicValid", licensePicValid)
		v.RegisterValidation("AddressPosValid", addressPosValid)
		v.RegisterValidation("VerifyStatusValid", verifyStatusValid)
	}
}
