package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _CommandsMgr struct {
	*_BaseMgr
}

// CommandsMgr open func
func CommandsMgr(db *gorm.DB) *_CommandsMgr {
	if db == nil {
		panic(fmt.Errorf("CommandsMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_CommandsMgr{_BaseMgr: &_BaseMgr{DB: db.Table("commands"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_CommandsMgr) GetTableName() string {
	return "commands"
}

// Reset 重置gorm会话
func (obj *_CommandsMgr) Reset() *_CommandsMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_CommandsMgr) Get() (result Commands, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Commands{}).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_CommandsMgr) Gets() (results []*Commands, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Commands{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_CommandsMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(Commands{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithRowid rowid获取
func (obj *_CommandsMgr) WithRowid(rowid int) Option {
	return optionFunc(func(o *options) { o.query["rowid"] = rowid })
}

// WithCommand command获取
func (obj *_CommandsMgr) WithCommand(command string) Option {
	return optionFunc(func(o *options) { o.query["command"] = command })
}

// GetByOption 功能选项模式获取
func (obj *_CommandsMgr) GetByOption(opts ...Option) (result Commands, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Commands{}).Where(options.query).Find(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_CommandsMgr) GetByOptions(opts ...Option) (results []*Commands, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Commands{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromRowid 通过rowid获取内容
func (obj *_CommandsMgr) GetFromRowid(rowid int) (result Commands, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Commands{}).Where("`rowid` = ?", rowid).Find(&result).Error

	return
}

// GetBatchFromRowid 批量查找
func (obj *_CommandsMgr) GetBatchFromRowid(rowids []int) (results []*Commands, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Commands{}).Where("`rowid` IN (?)", rowids).Find(&results).Error

	return
}

// GetFromCommand 通过command获取内容
func (obj *_CommandsMgr) GetFromCommand(command string) (results []*Commands, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Commands{}).Where("`command` = ?", command).Find(&results).Error

	return
}

// GetBatchFromCommand 批量查找
func (obj *_CommandsMgr) GetBatchFromCommand(commands []string) (results []*Commands, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Commands{}).Where("`command` IN (?)", commands).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_CommandsMgr) FetchByPrimaryKey(rowid int) (result Commands, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Commands{}).Where("`rowid` = ?", rowid).Find(&result).Error

	return
}
