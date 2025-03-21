package utils

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

// TranslateValidationError 接收校验错误 + 字段映射，返回首条错误信息（中文）
func TranslateValidationError(errs validator.ValidationErrors, fieldMap map[string]string) string {
	if len(errs) == 0 {
		return "参数验证失败"
	}

	fe := errs[0] // 这里只取第一条（也可以遍历所有）
	name, ok := fieldMap[fe.Field()]
	if !ok {
		name = fe.Field()
	}

	switch fe.Tag() {
	case "required":
		return fmt.Sprintf("%s是必填字段", name)
	case "gte":
		return fmt.Sprintf("%s不能小于 %s", name, fe.Param())
	case "lte":
		return fmt.Sprintf("%s不能大于 %s", name, fe.Param())
	default:
		return fmt.Sprintf("%s格式不正确", name)
	}
}
