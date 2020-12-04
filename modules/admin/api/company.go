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

// 获取公司的基础信息
func CompanyList(c *gin.Context) {
	var params myValidator.CompanyListParams
	if err := c.ShouldBind(&params); err != nil {
		response.Fail(c, 200, e.INVALID_PARAMS)
		return
	}
	if params.Page == 0 {
		params.Page = 1
	}
	if params.Size == 0 {
		params.Size = 15
	}
	conditions, conditionsParams := "status > ?", []interface{}{-1}
	if params.Status > 0 {
		conditions += " AND status = ?"
		conditionsParams = append(conditionsParams, params.Status-1)
	}
	if params.CompanyId != 0 {
		conditions += " AND id = ?"
		conditionsParams = append(conditionsParams, params.CompanyId)
	}
	if params.Name != "" {
		conditions += " AND name LIKE ?"
		conditionsParams = append(conditionsParams, "%"+params.Name+"%")
	}
	if params.ContactName != "" {
		conditions += " AND contact_name LIKE ?"
		conditionsParams = append(conditionsParams, "%"+params.ContactName+"%")
	}
	if params.ContactMobile != "" {
		conditions += " AND contact_mobile LIKE ?"
		conditionsParams = append(conditionsParams, "%"+params.ContactMobile+"%")
	}
	if params.Address != "" {
		conditions += " AND address LIKE ?"
		conditionsParams = append(conditionsParams, "%"+params.Address+"%")
	}
	companyModel := models.WkCompany{}
	pages, has := companyModel.Pages("*", "`status` ASC,id desc", params.Page, params.Size, conditions, conditionsParams...)
	var data interface{}
	if has {
		data = pages
	}
	response.Success(c, e.SUCCESS, data)
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
	//如果是保存到草稿，那么后台就不显示这条记录以及审核操作
	if params.IsDraft == 1 {
		companyModel.Status = -1
	}
	if companyModel.Insert(companyModel) == 0 {
		response.Fail(c, 200, e.ERR_COMPANY_CREATE_FAIL)
		return
	}
	response.Success(c, 200)
}

// 公司信息变更
func CompanyChange(c *gin.Context) {
	var params myValidator.CompanyRegisterParams
	if err := c.ShouldBind(&params); err != nil {
		response.Fail(c, 200, e.INVALID_PARAMS, params.GetError(err.(validator.ValidationErrors)))
		return
	}
	if params.CompanyId == 0 {
		response.Fail(c, 200, e.INVALID_PARAMS)
		return
	}
	companyInfo := map[string]interface{}{
		"name":           params.Name,
		"contact_name":   params.ContactName,
		"contact_mobile": params.ContactMobile,
		"address":        params.Address,
		"address_pos":    params.AddressPos,
		"license_pic":    params.LicensePic,
		"status":         0,
	}

	if !new(models.WkCompany).Update(companyInfo, "id = ?", params.CompanyId) {
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
