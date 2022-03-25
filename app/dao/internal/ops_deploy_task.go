// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
)

// OpsDeployTaskDao is the manager for logic model data accessing and custom defined data operations functions management.
type OpsDeployTaskDao struct {
	Table   string               // Table is the underlying table name of the DAO.
	Group   string               // Group is the database configuration group name of current DAO.
	Columns OpsDeployTaskColumns // Columns is the short type for Columns, which contains all the column names of Table for convenient usage.
}

// OpsDeployTaskColumns defines and stores column names for table tbl_ops_deploy_task.
type OpsDeployTaskColumns struct {
	Id            string // 部署任务ID
	Name          string // 部署任务名称
	Remark        string // 部署任务描述
	DeployImageId string // 部署镜像ID
	Progress      string // 进度(-1:部署失败,0:未部署,1:部署中,2:部署完成)
	FailMsg       string // 失败消息
	Status        string // 状态(-1:删除,0:停用,1:启用)
	CreateTime    string // 创建时间
	UpdateTime    string // 更新时间
}

//  opsDeployTaskColumns holds the columns for table tbl_ops_deploy_task.
var opsDeployTaskColumns = OpsDeployTaskColumns{
	Id:            "id",
	Name:          "name",
	Remark:        "remark",
	DeployImageId: "deploy_image_id",
	Progress:      "progress",
	FailMsg:       "fail_msg",
	Status:        "status",
	CreateTime:    "create_time",
	UpdateTime:    "update_time",
}

// NewOpsDeployTaskDao creates and returns a new DAO object for table data access.
func NewOpsDeployTaskDao() *OpsDeployTaskDao {
	return &OpsDeployTaskDao{
		Group:   "default",
		Table:   "tbl_ops_deploy_task",
		Columns: opsDeployTaskColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *OpsDeployTaskDao) DB() gdb.DB {
	return g.DB(dao.Group)
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *OpsDeployTaskDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.Table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *OpsDeployTaskDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
