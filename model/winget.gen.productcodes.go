package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _ProductcodesMgr struct {
	*_BaseMgr
}

// ProductcodesMgr open func
func ProductcodesMgr(db *gorm.DB) *_ProductcodesMgr {
	if db == nil {
		panic(fmt.Errorf("ProductcodesMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ProductcodesMgr{_BaseMgr: &_BaseMgr{DB: db.Table("productcodes"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_ProductcodesMgr) GetTableName() string {
	return "productcodes"
}

// Reset 重置gorm会话
func (obj *_ProductcodesMgr) Reset() *_ProductcodesMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_ProductcodesMgr) Get() (result Productcodes, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Productcodes{}).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_ProductcodesMgr) Gets() (results []*Productcodes, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Productcodes{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_ProductcodesMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(Productcodes{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithRowid rowid获取
func (obj *_ProductcodesMgr) WithRowid(rowid int) Option {
	return optionFunc(func(o *options) { o.query["rowid"] = rowid })
}

// WithProductcode productcode获取
func (obj *_ProductcodesMgr) WithProductcode(productcode string) Option {
	return optionFunc(func(o *options) { o.query["productcode"] = productcode })
}

// GetByOption 功能选项模式获取
func (obj *_ProductcodesMgr) GetByOption(opts ...Option) (result Productcodes, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Productcodes{}).Where(options.query).Find(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_ProductcodesMgr) GetByOptions(opts ...Option) (results []*Productcodes, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Productcodes{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromRowid 通过rowid获取内容
func (obj *_ProductcodesMgr) GetFromRowid(rowid int) (result Productcodes, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Productcodes{}).Where("`rowid` = ?", rowid).Find(&result).Error

	return
}

// GetBatchFromRowid 批量查找
func (obj *_ProductcodesMgr) GetBatchFromRowid(rowids []int) (results []*Productcodes, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Productcodes{}).Where("`rowid` IN (?)", rowids).Find(&results).Error

	return
}

// GetFromProductcode 通过productcode获取内容
func (obj *_ProductcodesMgr) GetFromProductcode(productcode string) (results []*Productcodes, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Productcodes{}).Where("`productcode` = ?", productcode).Find(&results).Error

	return
}

// GetBatchFromProductcode 批量查找
func (obj *_ProductcodesMgr) GetBatchFromProductcode(productcodes []string) (results []*Productcodes, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Productcodes{}).Where("`productcode` IN (?)", productcodes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_ProductcodesMgr) FetchByPrimaryKey(rowid int) (result Productcodes, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Productcodes{}).Where("`rowid` = ?", rowid).Find(&result).Error

	return
}
