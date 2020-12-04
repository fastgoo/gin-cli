package validator

import (
	"gin-cli/modules/admin/models"
	"gin-cli/pkg/e"
	"gin-cli/pkg/util"
	"gin-cli/plugins/qiniu"
	"github.com/go-playground/validator/v10"
	"strconv"
	"strings"
)

type CompanyRegisterParams struct {
	CompanyId     int    `form:"company_id" json:"company_id"`
	Name          string `form:"name" json:"name" binding:"required,CompanyNameNotExistValid"`
	ContactName   string `form:"contact_name" json:"contact_name" binding:"required"`
	ContactMobile string `form:"contact_mobile"  json:"contact_mobile"  binding:"required,TelValid"`
	Address       string `form:"address"  json:"address"  binding:"required"`
	AddressPos    string `form:"address_pos"  json:"address_pos"  binding:"required,AddressPosValid"`
	LicensePic    string `form:"license_pic"  json:"license_pic"  binding:"required,LicensePicValid"`
	IsDraft       int    `form:"is_draft"  json:"is_draft"`
}

type CompanyVerifyParams struct {
	CompanyId    int    `form:"company_id" json:"company_id" binding:"required"`
	VerifyRemark string `form:"remark" json:"remark" binding:"required"`
	Status       int    `form:"status"  json:"status"  binding:"VerifyStatusValid"`
}

type CompanyListParams struct {
	Page          int    `form:"page" json:"page"`
	Size          int    `form:"size" json:"size"`
	CompanyId     int    `form:"company_id" json:"company_id"`
	Name          string `form:"company_name" json:"company_name"`
	ContactName   string `form:"contact_name" json:"contact_name"`
	ContactMobile string `form:"contact_mobile" json:"contact_mobile"`
	Address       string `form:"address" json:"address"`
	Status        int    `form:"status" json:"status"`
}

// 绑定模型获取验证错误的方法
func (c *CompanyRegisterParams) GetError(errs validator.ValidationErrors) string {
	for _, err := range errs {
		if err.Field() == "ContactName" {
			switch err.Tag() {
			case "required":
				return e.GetErrMsg(e.ERR_PARAMS_EMPTY_MOBILE)
			}
		}
		if err.Field() == "Address" {
			switch err.Tag() {
			case "required":
				return e.GetErrMsg(e.ERR_PARAMS_EMPTY_MOBILE)
			}
		}
		if err.Field() == "Name" {
			switch err.Tag() {
			case "required":
				return e.GetErrMsg(e.ERR_PARAMS_EMPTY_MOBILE)
			case "CompanyNameNotExistValid":
				return e.GetErrMsg(e.ERR_PARAMS_MOBILE_INVALID)
			}
		}
		if err.Field() == "ContactMobile" {
			switch err.Tag() {
			case "required":
				return e.GetErrMsg(e.ERR_PARAMS_EMPTY_MOBILE)
			case "TelValid":
				return e.GetErrMsg(e.ERR_PARAMS_MOBILE_INVALID)
			}
		}
		if err.Field() == "AddressPos" {
			switch err.Tag() {
			case "required":
				return e.GetErrMsg(e.ERR_PARAMS_EMPTY_PASSWORD)
			case "TelValid":
				return e.GetErrMsg(e.ERR_PARAMS_MOBILE_INVALID)
			}
		}
		if err.Field() == "LicensePic" {
			switch err.Tag() {
			case "required":
				return e.GetErrMsg(e.ERR_PARAMS_EMPTY_PASSWORD)
			case "LicensePicValid":
				return e.GetErrMsg(e.ERR_PARAMS_MOBILE_INVALID)
			}
		}
	}
	return e.GetErrMsg(e.INVALID_PARAMS)
}

// 绑定模型获取验证错误的方法
func (c *CompanyVerifyParams) GetError(errs validator.ValidationErrors) string {
	for _, err := range errs {
		if err.Field() == "CompanyId" {
			switch err.Tag() {
			case "required":
				return e.GetErrMsg(e.ERR_PARAMS_COMPANY_VERIFY_EMPTY_COMPANY_ID)
			}
		}
		if err.Field() == "Status" {
			switch err.Tag() {
			case "VerifyStatusValid":
				return e.GetErrMsg(e.ERR_PARAMS_COMPANY_VERIFY_STATUS_ERROR)
			}
		}
	}
	return e.GetErrMsg(e.INVALID_PARAMS)
}

// 校验公司名称是否重复
func companyNameNotExistValid(fl validator.FieldLevel) bool {
	companyModel := models.WkCompany{}
	_, has := companyModel.Get("id", "name = ?", fl.Field().String())
	return !has
}

// 校验联系方式是否是手机号码或者座机
func telValid(fl validator.FieldLevel) bool {
	return util.VerifyMobileFormat(fl.Field().String()) || util.VerifyTelFormat(fl.Field().String())
}

// 校验图片是否是自己服务器七牛云上传的
func licensePicValid(fl validator.FieldLevel) bool {
	return strings.Contains(fl.Field().String(), qiniu.GetUrl(""))
}

// 0待认证 1已认证 2认证失败
func verifyStatusValid(fl validator.FieldLevel) bool {
	status := fl.Field().Int()
	return status == 1 || status == 2 || status == 3
}

// 地址定位是否是逗号分隔的经纬度
func addressPosValid(fl validator.FieldLevel) bool {
	posArr := strings.Split(fl.Field().String(), ",")
	if len(posArr) != 2 {
		return false
	}
	lng, lat := posArr[0], posArr[1]
	_, err := strconv.ParseFloat(lng, len(lng))
	if err != nil {
		return false
	}
	_, err = strconv.ParseFloat(lng, len(lat))
	if err != nil {
		return false
	}
	return true
}
