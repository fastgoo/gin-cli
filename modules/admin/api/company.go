package api

import (
	"gin-cli/modules/admin/models"
	"gin-cli/pkg/e"
	"gin-cli/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	myValidator "gin-cli/modules/admin/validator"
)

// 获取公司的基础信息
func CompanyInfo(c *gin.Context) {
	companyId := c.Param("company_id")
	companyModel := models.WkCompany{}
	info, has := companyModel.Get("*", "id = ?", companyId)
	if !has {
		response.Fail(c, 200, e.ERR_COMPANY_GET_FAIL)
		return
	}
	response.Success(c, e.SUCCESS, info)
}

// 公司信息提交
func CompanySave(c *gin.Context) {
	var params myValidator.CompanyRegisterParams
	if err := c.ShouldBind(&params); err != nil {
		response.Fail(c, 200, e.INVALID_PARAMS, params.GetError(err.(validator.ValidationErrors)))
		return
	}
	companyModel := models.WkCompany{
		Name:          params.Name,
		ContactName:   params.ContactName,
		ContactMobile: params.ContactMobile,
		Address:       params.Address,
		AddressPos:    params.AddressPos,
		LicensePic:    params.LicensePic,
	}
	if companyModel.Insert(companyModel) == 0 {
		response.Fail(c, 200, e.ERR_COMPANY_CREATE_FAIL)
		return
	}
	response.Success(c, 200)
}

// 公司主体认证
func CompanyVerify(c *gin.Context) {
	var params myValidator.CompanyVerifyParams
	if err := c.ShouldBind(&params); err != nil {
		response.Fail(c, 200, e.INVALID_PARAMS, params.GetError(err.(validator.ValidationErrors)))
		return
	}
	companyModel := models.WkCompany{}
	if !companyModel.Update(map[string]interface{}{"status": params.Status, "verify_remark": params.VerifyRemark}, "id = ?", params.CompanyId) {
		response.Fail(c, 200, e.ERR_COMPANY_VERIFY_FAIL)
		return
	}
	response.Success(c, 200)
}
