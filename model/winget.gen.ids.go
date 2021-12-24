package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _IDsMgr struct {
	*_BaseMgr
}

// IDsMgr open func
func IDsMgr(db *gorm.DB) *_IDsMgr {
	if db == nil {
		panic(fmt.Errorf("IDsMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_IDsMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ids"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_IDsMgr) GetTableName() string {
	return "ids"
}

// Reset 重置gorm会话
func (obj *_IDsMgr) Reset() *_IDsMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_IDsMgr) Get() (result IDs, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(IDs{}).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_IDsMgr) Gets() (results []*IDs, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(IDs{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_IDsMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(IDs{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithRowid rowid获取
func (obj *_IDsMgr) WithRowid(rowid int) Option {
	return optionFunc(func(o *options) { o.query["rowid"] = rowid })
}

// WithID id获取
func (obj *_IDsMgr) WithID(id string) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// GetByOption 功能选项模式获取
func (obj *_IDsMgr) GetByOption(opts ...Option) (result IDs, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(IDs{}).Where(options.query).Find(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_IDsMgr) GetByOptions(opts ...Option) (results []*IDs, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(IDs{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromRowid 通过rowid获取内容
func (obj *_IDsMgr) GetFromRowid(rowid int) (result IDs, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(IDs{}).Where("`rowid` = ?", rowid).Find(&result).Error

	return
}

// GetBatchFromRowid 批量查找
func (obj *_IDsMgr) GetBatchFromRowid(rowids []int) (results []*IDs, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(IDs{}).Where("`rowid` IN (?)", rowids).Find(&results).Error

	return
}

// GetFromID 通过id获取内容
func (obj *_IDsMgr) GetFromID(id string) (results []*IDs, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(IDs{}).Where("`id` = ?", id).Find(&results).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_IDsMgr) GetBatchFromID(ids []string) (results []*IDs, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(IDs{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_IDsMgr) FetchByPrimaryKey(rowid int) (result IDs, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(IDs{}).Where("`rowid` = ?", rowid).Find(&result).Error

	return
}
