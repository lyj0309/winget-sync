package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _NormPublishersMgr struct {
	*_BaseMgr
}

// NormPublishersMgr open func
func NormPublishersMgr(db *gorm.DB) *_NormPublishersMgr {
	if db == nil {
		panic(fmt.Errorf("NormPublishersMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_NormPublishersMgr{_BaseMgr: &_BaseMgr{DB: db.Table("norm_publishers"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_NormPublishersMgr) GetTableName() string {
	return "norm_publishers"
}

// Reset 重置gorm会话
func (obj *_NormPublishersMgr) Reset() *_NormPublishersMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_NormPublishersMgr) Get() (result NormPublishers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(NormPublishers{}).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_NormPublishersMgr) Gets() (results []*NormPublishers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(NormPublishers{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_NormPublishersMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(NormPublishers{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithRowid rowid获取
func (obj *_NormPublishersMgr) WithRowid(rowid int) Option {
	return optionFunc(func(o *options) { o.query["rowid"] = rowid })
}

// WithNormPublisher norm_publisher获取
func (obj *_NormPublishersMgr) WithNormPublisher(normPublisher string) Option {
	return optionFunc(func(o *options) { o.query["norm_publisher"] = normPublisher })
}

// GetByOption 功能选项模式获取
func (obj *_NormPublishersMgr) GetByOption(opts ...Option) (result NormPublishers, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(NormPublishers{}).Where(options.query).Find(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_NormPublishersMgr) GetByOptions(opts ...Option) (results []*NormPublishers, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(NormPublishers{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromRowid 通过rowid获取内容
func (obj *_NormPublishersMgr) GetFromRowid(rowid int) (result NormPublishers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(NormPublishers{}).Where("`rowid` = ?", rowid).Find(&result).Error

	return
}

// GetBatchFromRowid 批量查找
func (obj *_NormPublishersMgr) GetBatchFromRowid(rowids []int) (results []*NormPublishers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(NormPublishers{}).Where("`rowid` IN (?)", rowids).Find(&results).Error

	return
}

// GetFromNormPublisher 通过norm_publisher获取内容
func (obj *_NormPublishersMgr) GetFromNormPublisher(normPublisher string) (results []*NormPublishers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(NormPublishers{}).Where("`norm_publisher` = ?", normPublisher).Find(&results).Error

	return
}

// GetBatchFromNormPublisher 批量查找
func (obj *_NormPublishersMgr) GetBatchFromNormPublisher(normPublishers []string) (results []*NormPublishers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(NormPublishers{}).Where("`norm_publisher` IN (?)", normPublishers).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_NormPublishersMgr) FetchByPrimaryKey(rowid int) (result NormPublishers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(NormPublishers{}).Where("`rowid` = ?", rowid).Find(&result).Error

	return
}
