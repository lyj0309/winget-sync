package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _NamesMgr struct {
	*_BaseMgr
}

// NamesMgr open func
func NamesMgr(db *gorm.DB) *_NamesMgr {
	if db == nil {
		panic(fmt.Errorf("NamesMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_NamesMgr{_BaseMgr: &_BaseMgr{DB: db.Table("names"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_NamesMgr) GetTableName() string {
	return "names"
}

// Reset 重置gorm会话
func (obj *_NamesMgr) Reset() *_NamesMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_NamesMgr) Get() (result Names, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Names{}).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_NamesMgr) Gets() (results []*Names, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Names{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_NamesMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(Names{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithRowid rowid获取
func (obj *_NamesMgr) WithRowid(rowid int) Option {
	return optionFunc(func(o *options) { o.query["rowid"] = rowid })
}

// WithName name获取
func (obj *_NamesMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// GetByOption 功能选项模式获取
func (obj *_NamesMgr) GetByOption(opts ...Option) (result Names, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Names{}).Where(options.query).Find(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_NamesMgr) GetByOptions(opts ...Option) (results []*Names, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Names{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromRowid 通过rowid获取内容
func (obj *_NamesMgr) GetFromRowid(rowid int) (result Names, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Names{}).Where("`rowid` = ?", rowid).Find(&result).Error

	return
}

// GetBatchFromRowid 批量查找
func (obj *_NamesMgr) GetBatchFromRowid(rowids []int) (results []*Names, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Names{}).Where("`rowid` IN (?)", rowids).Find(&results).Error

	return
}

// GetFromName 通过name获取内容
func (obj *_NamesMgr) GetFromName(name string) (results []*Names, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Names{}).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找
func (obj *_NamesMgr) GetBatchFromName(names []string) (results []*Names, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Names{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_NamesMgr) FetchByPrimaryKey(rowid int) (result Names, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Names{}).Where("`rowid` = ?", rowid).Find(&result).Error

	return
}
