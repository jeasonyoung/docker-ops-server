// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
)

// OpsGroupDao is the manager for logic model data accessing and custom defined data operations functions management.
type OpsGroupDao struct {
	Table   string          // Table is the underlying table name of the DAO.
	Group   string          // Group is the database configuration group name of current DAO.
	Columns OpsGroupColumns // Columns is the short type for Columns, which contains all the column names of Table for convenient usage.
}

// OpsGroupColumns defines and stores column names for table tbl_ops_group.
type OpsGroupColumns struct {
	Id         string // 分组ID
	Name       string // 分组名称
	Remark     string // 分组描述
	Status     string // 状态(-1:删除,0:停用,1:启用)
	CreateTime string // 创建时间
	UpdateTime string // 更新时间
}

//  opsGroupColumns holds the columns for table tbl_ops_group.
var opsGroupColumns = OpsGroupColumns{
	Id:         "id",
	Name:       "name",
	Remark:     "remark",
	Status:     "status",
	CreateTime: "create_time",
	UpdateTime: "update_time",
}

// NewOpsGroupDao creates and returns a new DAO object for table data access.
func NewOpsGroupDao() *OpsGroupDao {
	return &OpsGroupDao{
		Group:   "default",
		Table:   "tbl_ops_group",
		Columns: opsGroupColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *OpsGroupDao) DB() gdb.DB {
	return g.DB(dao.Group)
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *OpsGroupDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.Table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *OpsGroupDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}