package request

import (
	"ginshop.com/admin/model"
)

type SysOperationRecordSearch struct {
	model.SysOperationRecord
	PageInfo
}
