package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _MonikersMgr struct {
	*_BaseMgr
}

// MonikersMgr open func
func MonikersMgr(db *gorm.DB) *_MonikersMgr {
	if db == nil {
		panic(fmt.Errorf("MonikersMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_MonikersMgr{_BaseMgr: &_BaseMgr{DB: db.Table("monikers"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_MonikersMgr) GetTableName() string {
	return "monikers"
}

// Reset 重置gorm会话
func (obj *_MonikersMgr) Reset() *_MonikersMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_MonikersMgr) Get() (result Monikers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Monikers{}).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_MonikersMgr) Gets() (results []*Monikers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Monikers{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_MonikersMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(Monikers{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithRowid rowid获取
func (obj *_MonikersMgr) WithRowid(rowid int) Option {
	return optionFunc(func(o *options) { o.query["rowid"] = rowid })
}

// WithMoniker moniker获取
func (obj *_MonikersMgr) WithMoniker(moniker string) Option {
	return optionFunc(func(o *options) { o.query["moniker"] = moniker })
}

// GetByOption 功能选项模式获取
func (obj *_MonikersMgr) GetByOption(opts ...Option) (result Monikers, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Monikers{}).Where(options.query).Find(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_MonikersMgr) GetByOptions(opts ...Option) (results []*Monikers, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Monikers{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromRowid 通过rowid获取内容
func (obj *_MonikersMgr) GetFromRowid(rowid int) (result Monikers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Monikers{}).Where("`rowid` = ?", rowid).Find(&result).Error

	return
}

// GetBatchFromRowid 批量查找
func (obj *_MonikersMgr) GetBatchFromRowid(rowids []int) (results []*Monikers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Monikers{}).Where("`rowid` IN (?)", rowids).Find(&results).Error

	return
}

// GetFromMoniker 通过moniker获取内容
func (obj *_MonikersMgr) GetFromMoniker(moniker string) (results []*Monikers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Monikers{}).Where("`moniker` = ?", moniker).Find(&results).Error

	return
}

// GetBatchFromMoniker 批量查找
func (obj *_MonikersMgr) GetBatchFromMoniker(monikers []string) (results []*Monikers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Monikers{}).Where("`moniker` IN (?)", monikers).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_MonikersMgr) FetchByPrimaryKey(rowid int) (result Monikers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Monikers{}).Where("`rowid` = ?", rowid).Find(&result).Error

	return
}
