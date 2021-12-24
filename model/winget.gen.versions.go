package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _VersionsMgr struct {
	*_BaseMgr
}

// VersionsMgr open func
func VersionsMgr(db *gorm.DB) *_VersionsMgr {
	if db == nil {
		panic(fmt.Errorf("VersionsMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_VersionsMgr{_BaseMgr: &_BaseMgr{DB: db.Table("versions"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_VersionsMgr) GetTableName() string {
	return "versions"
}

// Reset 重置gorm会话
func (obj *_VersionsMgr) Reset() *_VersionsMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_VersionsMgr) Get() (result Versions, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Versions{}).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_VersionsMgr) Gets() (results []*Versions, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Versions{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_VersionsMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(Versions{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithRowid rowid获取
func (obj *_VersionsMgr) WithRowid(rowid int) Option {
	return optionFunc(func(o *options) { o.query["rowid"] = rowid })
}

// WithVersion version获取
func (obj *_VersionsMgr) WithVersion(version string) Option {
	return optionFunc(func(o *options) { o.query["version"] = version })
}

// GetByOption 功能选项模式获取
func (obj *_VersionsMgr) GetByOption(opts ...Option) (result Versions, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Versions{}).Where(options.query).Find(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_VersionsMgr) GetByOptions(opts ...Option) (results []*Versions, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Versions{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromRowid 通过rowid获取内容
func (obj *_VersionsMgr) GetFromRowid(rowid int) (result Versions, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Versions{}).Where("`rowid` = ?", rowid).Find(&result).Error

	return
}

// GetBatchFromRowid 批量查找
func (obj *_VersionsMgr) GetBatchFromRowid(rowids []int) (results []*Versions, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Versions{}).Where("`rowid` IN (?)", rowids).Find(&results).Error

	return
}

// GetFromVersion 通过version获取内容
func (obj *_VersionsMgr) GetFromVersion(version string) (results []*Versions, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Versions{}).Where("`version` = ?", version).Find(&results).Error

	return
}

// GetBatchFromVersion 批量查找
func (obj *_VersionsMgr) GetBatchFromVersion(versions []string) (results []*Versions, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Versions{}).Where("`version` IN (?)", versions).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_VersionsMgr) FetchByPrimaryKey(rowid int) (result Versions, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Versions{}).Where("`rowid` = ?", rowid).Find(&result).Error

	return
}
