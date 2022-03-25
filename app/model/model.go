// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package model

import (
	"github.com/gogf/gf/os/gtime"
)

// DictData is the golang structure for table dict_data.
type DictData struct {
	Id         uint64      `orm:"id,primary"  json:"id"`         // 字典数据ID
	Code       uint        `orm:"code"        json:"code"`       // 字典代码(排序)
	Label      string      `orm:"label"       json:"label"`      // 字典标签
	Value      string      `orm:"value"       json:"value"`      // 字典键值
	IsDefault  uint        `orm:"is_default"  json:"isDefault"`  // 是否默认(0:否,1:是)
	Type       string      `orm:"type"        json:"type"`       // 字典类型
	CssClass   string      `orm:"css_class"   json:"cssClass"`   // 样式属性
	ListClass  string      `orm:"list_class"  json:"listClass"`  // 表格回显样式
	Remark     string      `orm:"remark"      json:"remark"`     // 字典数据备注
	Status     int         `orm:"status"      json:"status"`     // 状态(-1:删除,0:停用,1:启用)
	CreateTime *gtime.Time `orm:"create_time" json:"createTime"` // 创建时间
	UpdateTime *gtime.Time `orm:"update_time" json:"updateTime"` // 更新时间
}

// DictType is the golang structure for table dict_type.
type DictType struct {
	Id         uint64      `orm:"id,primary"  json:"id"`         // 字典ID
	Name       string      `orm:"name"        json:"name"`       // 字典名称
	Type       string      `orm:"type,unique" json:"type"`       // 字典类型
	Remark     string      `orm:"remark"      json:"remark"`     // 字典备注
	Status     int         `orm:"status"      json:"status"`     // 状态(-1:删除,0:停用,1:启用)
	CreateTime *gtime.Time `orm:"create_time" json:"createTime"` // 创建时间
	UpdateTime *gtime.Time `orm:"update_time" json:"updateTime"` // 更新时间
}

// OpsDeploy is the golang structure for table ops_deploy.
type OpsDeploy struct {
	Id           uint64      `orm:"id,primary"    json:"id"`           // 部署ID
	Name         string      `orm:"name,unique"   json:"name"`         // 部署名称
	Remark       string      `orm:"remark"        json:"remark"`       // 部署描述
	GroupId      uint64      `orm:"group_id"      json:"groupId"`      // 服务器分组ID
	RepositoryId uint64      `orm:"repository_id" json:"repositoryId"` // 所属镜像仓库ID
	Status       int         `orm:"status"        json:"status"`       // 状态(-1:删除,0:停用,1:启用)
	CreateTime   *gtime.Time `orm:"create_time"   json:"createTime"`   // 创建时间
	UpdateTime   *gtime.Time `orm:"update_time"   json:"updateTime"`   // 更新时间
}

// OpsDeployTask is the golang structure for table ops_deploy_task.
type OpsDeployTask struct {
	Id            uint64      `orm:"id,primary"      json:"id"`            // 部署任务ID
	Name          string      `orm:"name,unique"     json:"name"`          // 部署任务名称
	Remark        string      `orm:"remark"          json:"remark"`        // 部署任务描述
	DeployImageId uint64      `orm:"deploy_image_id" json:"deployImageId"` // 部署镜像ID
	Progress      int         `orm:"progress"        json:"progress"`      // 进度(-1:部署失败,0:未部署,1:部署中,2:部署完成)
	FailMsg       string      `orm:"fail_msg"        json:"failMsg"`       // 失败消息
	Status        int         `orm:"status"          json:"status"`        // 状态(-1:删除,0:停用,1:启用)
	CreateTime    *gtime.Time `orm:"create_time"     json:"createTime"`    // 创建时间
	UpdateTime    *gtime.Time `orm:"update_time"     json:"updateTime"`    // 更新时间
}

// OpsDeployTaskServer is the golang structure for table ops_deploy_task_server.
type OpsDeployTaskServer struct {
	Id         uint64      `orm:"id,primary"  json:"id"`         // 部署任务服务器ID
	TaskId     uint64      `orm:"task_id"     json:"taskId"`     // 所属部署任务ID
	ServerId   uint64      `orm:"server_id"   json:"serverId"`   // 所属服务器ID
	Progress   int         `orm:"progress"    json:"progress"`   // 进度(-1:部署失败,0:未部署,1:部署中,2:部署完成)
	FailMsg    string      `orm:"fail_msg"    json:"failMsg"`    // 失败消息
	CreateTime *gtime.Time `orm:"create_time" json:"createTime"` // 创建时间
	UpdateTime *gtime.Time `orm:"update_time" json:"updateTime"` // 更新时间
}

// OpsGroup is the golang structure for table ops_group.
type OpsGroup struct {
	Id         uint64      `orm:"id,primary"  json:"id"`         // 分组ID
	Name       uint64      `orm:"name,unique" json:"name"`       // 分组名称
	Remark     string      `orm:"remark"      json:"remark"`     // 分组描述
	Status     int         `orm:"status"      json:"status"`     // 状态(-1:删除,0:停用,1:启用)
	CreateTime *gtime.Time `orm:"create_time" json:"createTime"` // 创建时间
	UpdateTime *gtime.Time `orm:"update_time" json:"updateTime"` // 更新时间
}

// OpsGroupServers is the golang structure for table ops_group_servers.
type OpsGroupServers struct {
	GroupId  uint64 `orm:"group_id,primary"  json:"groupId"`  // 分组ID
	ServerId uint64 `orm:"server_id,primary" json:"serverId"` // 服务器ID
}

// OpsRepository is the golang structure for table ops_repository.
type OpsRepository struct {
	Id            uint64      `orm:"id,primary"      json:"id"`            // 镜像仓库ID
	Name          string      `orm:"name,unique"     json:"name"`          // 镜像仓库名称
	Remark        string      `orm:"remark"          json:"remark"`        // 镜像仓库描述
	Addr          string      `orm:"addr"            json:"addr"`          // 镜像仓库地址
	TmpLoginCmd   string      `orm:"tmp_login_cmd"   json:"tmpLoginCmd"`   // 临时登录指令
	TmpExpireTime *gtime.Time `orm:"tmp_expire_time" json:"tmpExpireTime"` // 临时指令到期时间
	Status        int         `orm:"status"          json:"status"`        // 状态(-1:删除,0:停用,1:启用)
	CreateTime    *gtime.Time `orm:"create_time"     json:"createTime"`    // 创建时间
	UpdateTime    *gtime.Time `orm:"update_time"     json:"updateTime"`    // 更新时间
}

// OpsRepositoryImages is the golang structure for table ops_repository_images.
type OpsRepositoryImages struct {
	Id           uint64      `orm:"id,primary"    json:"id"`           // 镜像ID
	Title        string      `orm:"title"         json:"title"`        // 镜像标题
	ImageOrg     string      `orm:"image_org"     json:"imageOrg"`     // 镜像组织
	ImageName    string      `orm:"image_name"    json:"imageName"`    // 镜像名称
	ImageTag     string      `orm:"image_tag"     json:"imageTag"`     // 镜像标签(版本)
	Status       int         `orm:"status"        json:"status"`       // 状态(-1:删除,0:停用,1:启用)
	RepositoryId uint64      `orm:"repository_id" json:"repositoryId"` // 所属镜像仓库ID
	CreateTime   *gtime.Time `orm:"create_time"   json:"createTime"`   // 创建时间
	UpdateTime   *gtime.Time `orm:"update_time"   json:"updateTime"`   // 更新时间
}

// OpsServer is the golang structure for table ops_server.
type OpsServer struct {
	Id           uint64      `orm:"id,primary"       json:"id"`           // 服务器ID
	Name         string      `orm:"name,unique"      json:"name"`         // 服务器名称
	Title        string      `orm:"title"            json:"title"`        // 服务器标题(别名)
	AuthCode     string      `orm:"auth_code,unique" json:"authCode"`     // 服务器授权码(用于代理连接标识)
	InnerIpAddr  string      `orm:"inner_ip_addr"    json:"innerIpAddr"`  // 内网IP地址
	OuterIpAddr  string      `orm:"outer_ip_addr"    json:"outerIpAddr"`  // 外网IP地址
	Status       int         `orm:"status"           json:"status"`       // 状态(-1:删除,0:停用,1:启用)
	LastPingTime *gtime.Time `orm:"last_ping_time"   json:"lastPingTime"` // 最后心跳时间
	LastPingMsg  string      `orm:"last_ping_msg"    json:"lastPingMsg"`  // 最后心跳消息
	CreateTime   *gtime.Time `orm:"create_time"      json:"createTime"`   // 创建时间
	UpdateTime   *gtime.Time `orm:"update_time"      json:"updateTime"`   // 更新时间
}