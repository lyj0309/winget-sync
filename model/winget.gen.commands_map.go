package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _CommandsMapMgr struct {
	*_BaseMgr
}

// CommandsMapMgr open func
func CommandsMapMgr(db *gorm.DB) *_CommandsMapMgr {
	if db == nil {
		panic(fmt.Errorf("CommandsMapMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_CommandsMapMgr{_BaseMgr: &_BaseMgr{DB: db.Table("commands_map"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_CommandsMapMgr) GetTableName() string {
	return "commands_map"
}

// Reset 重置gorm会话
func (obj *_CommandsMapMgr) Reset() *_CommandsMapMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_CommandsMapMgr) Get() (result CommandsMap, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CommandsMap{}).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_CommandsMapMgr) Gets() (results []*CommandsMap, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CommandsMap{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_CommandsMapMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(CommandsMap{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithManifest manifest获取
func (obj *_CommandsMapMgr) WithManifest(manifest int) Option {
	return optionFunc(func(o *options) { o.query["manifest"] = manifest })
}

// WithCommand command获取
func (obj *_CommandsMapMgr) WithCommand(command int) Option {
	return optionFunc(func(o *options) { o.query["command"] = command })
}

// GetByOption 功能选项模式获取
func (obj *_CommandsMapMgr) GetByOption(opts ...Option) (result CommandsMap, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(CommandsMap{}).Where(options.query).Find(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_CommandsMapMgr) GetByOptions(opts ...Option) (results []*CommandsMap, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(CommandsMap{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromManifest 通过manifest获取内容
func (obj *_CommandsMapMgr) GetFromManifest(manifest int) (results []*CommandsMap, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CommandsMap{}).Where("`manifest` = ?", manifest).Find(&results).Error

	return
}

// GetBatchFromManifest 批量查找
func (obj *_CommandsMapMgr) GetBatchFromManifest(manifests []int) (results []*CommandsMap, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CommandsMap{}).Where("`manifest` IN (?)", manifests).Find(&results).Error

	return
}

// GetFromCommand 通过command获取内容
func (obj *_CommandsMapMgr) GetFromCommand(command int) (results []*CommandsMap, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CommandsMap{}).Where("`command` = ?", command).Find(&results).Error

	return
}

// GetBatchFromCommand 批量查找
func (obj *_CommandsMapMgr) GetBatchFromCommand(commands []int) (results []*CommandsMap, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CommandsMap{}).Where("`command` IN (?)", commands).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchUniqueIndexByCommandsMapPkindex primary or index 获取唯一内容
func (obj *_CommandsMapMgr) FetchUniqueIndexByCommandsMapPkindex(manifest int, command int) (result CommandsMap, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CommandsMap{}).Where("`manifest` = ? AND `command` = ?", manifest, command).Find(&result).Error

	return
}
