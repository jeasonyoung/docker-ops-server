// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
)

// DictTypeDao is the manager for logic model data accessing and custom defined data operations functions management.
type DictTypeDao struct {
	Table   string          // Table is the underlying table name of the DAO.
	Group   string          // Group is the database configuration group name of current DAO.
	Columns DictTypeColumns // Columns is the short type for Columns, which contains all the column names of Table for convenient usage.
}

// DictTypeColumns defines and stores column names for table tbl_dict_type.
type DictTypeColumns struct {
	Id         string // 字典ID
	Name       string // 字典名称
	Type       string // 字典类型
	Remark     string // 字典备注
	Status     string // 状态(-1:删除,0:停用,1:启用)
	CreateTime string // 创建时间
	UpdateTime string // 更新时间
}

//  dictTypeColumns holds the columns for table tbl_dict_type.
var dictTypeColumns = DictTypeColumns{
	Id:         "id",
	Name:       "name",
	Type:       "type",
	Remark:     "remark",
	Status:     "status",
	CreateTime: "create_time",
	UpdateTime: "update_time",
}

// NewDictTypeDao creates and returns a new DAO object for table data access.
func NewDictTypeDao() *DictTypeDao {
	return &DictTypeDao{
		Group:   "default",
		Table:   "tbl_dict_type",
		Columns: dictTypeColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *DictTypeDao) DB() gdb.DB {
	return g.DB(dao.Group)
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *DictTypeDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.Table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *DictTypeDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
