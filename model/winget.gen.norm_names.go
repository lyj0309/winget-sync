package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _NormNamesMgr struct {
	*_BaseMgr
}

// NormNamesMgr open func
func NormNamesMgr(db *gorm.DB) *_NormNamesMgr {
	if db == nil {
		panic(fmt.Errorf("NormNamesMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_NormNamesMgr{_BaseMgr: &_BaseMgr{DB: db.Table("norm_names"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_NormNamesMgr) GetTableName() string {
	return "norm_names"
}

// Reset 重置gorm会话
func (obj *_NormNamesMgr) Reset() *_NormNamesMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_NormNamesMgr) Get() (result NormNames, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(NormNames{}).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_NormNamesMgr) Gets() (results []*NormNames, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(NormNames{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_NormNamesMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(NormNames{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithRowid rowid获取
func (obj *_NormNamesMgr) WithRowid(rowid int) Option {
	return optionFunc(func(o *options) { o.query["rowid"] = rowid })
}

// WithNormName norm_name获取
func (obj *_NormNamesMgr) WithNormName(normName string) Option {
	return optionFunc(func(o *options) { o.query["norm_name"] = normName })
}

// GetByOption 功能选项模式获取
func (obj *_NormNamesMgr) GetByOption(opts ...Option) (result NormNames, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(NormNames{}).Where(options.query).Find(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_NormNamesMgr) GetByOptions(opts ...Option) (results []*NormNames, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(NormNames{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromRowid 通过rowid获取内容
func (obj *_NormNamesMgr) GetFromRowid(rowid int) (result NormNames, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(NormNames{}).Where("`rowid` = ?", rowid).Find(&result).Error

	return
}

// GetBatchFromRowid 批量查找
func (obj *_NormNamesMgr) GetBatchFromRowid(rowids []int) (results []*NormNames, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(NormNames{}).Where("`rowid` IN (?)", rowids).Find(&results).Error

	return
}

// GetFromNormName 通过norm_name获取内容
func (obj *_NormNamesMgr) GetFromNormName(normName string) (results []*NormNames, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(NormNames{}).Where("`norm_name` = ?", normName).Find(&results).Error

	return
}

// GetBatchFromNormName 批量查找
func (obj *_NormNamesMgr) GetBatchFromNormName(normNames []string) (results []*NormNames, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(NormNames{}).Where("`norm_name` IN (?)", normNames).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_NormNamesMgr) FetchByPrimaryKey(rowid int) (result NormNames, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(NormNames{}).Where("`rowid` = ?", rowid).Find(&result).Error

	return
}
