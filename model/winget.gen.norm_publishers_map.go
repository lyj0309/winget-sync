package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _NormPublishersMapMgr struct {
	*_BaseMgr
}

// NormPublishersMapMgr open func
func NormPublishersMapMgr(db *gorm.DB) *_NormPublishersMapMgr {
	if db == nil {
		panic(fmt.Errorf("NormPublishersMapMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_NormPublishersMapMgr{_BaseMgr: &_BaseMgr{DB: db.Table("norm_publishers_map"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_NormPublishersMapMgr) GetTableName() string {
	return "norm_publishers_map"
}

// Reset 重置gorm会话
func (obj *_NormPublishersMapMgr) Reset() *_NormPublishersMapMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_NormPublishersMapMgr) Get() (result NormPublishersMap, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(NormPublishersMap{}).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_NormPublishersMapMgr) Gets() (results []*NormPublishersMap, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(NormPublishersMap{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_NormPublishersMapMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(NormPublishersMap{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithManifest manifest获取
func (obj *_NormPublishersMapMgr) WithManifest(manifest int) Option {
	return optionFunc(func(o *options) { o.query["manifest"] = manifest })
}

// WithNormPublisher norm_publisher获取
func (obj *_NormPublishersMapMgr) WithNormPublisher(normPublisher int) Option {
	return optionFunc(func(o *options) { o.query["norm_publisher"] = normPublisher })
}

// GetByOption 功能选项模式获取
func (obj *_NormPublishersMapMgr) GetByOption(opts ...Option) (result NormPublishersMap, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(NormPublishersMap{}).Where(options.query).Find(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_NormPublishersMapMgr) GetByOptions(opts ...Option) (results []*NormPublishersMap, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(NormPublishersMap{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromManifest 通过manifest获取内容
func (obj *_NormPublishersMapMgr) GetFromManifest(manifest int) (results []*NormPublishersMap, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(NormPublishersMap{}).Where("`manifest` = ?", manifest).Find(&results).Error

	return
}

// GetBatchFromManifest 批量查找
func (obj *_NormPublishersMapMgr) GetBatchFromManifest(manifests []int) (results []*NormPublishersMap, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(NormPublishersMap{}).Where("`manifest` IN (?)", manifests).Find(&results).Error

	return
}

// GetFromNormPublisher 通过norm_publisher获取内容
func (obj *_NormPublishersMapMgr) GetFromNormPublisher(normPublisher int) (results []*NormPublishersMap, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(NormPublishersMap{}).Where("`norm_publisher` = ?", normPublisher).Find(&results).Error

	return
}

// GetBatchFromNormPublisher 批量查找
func (obj *_NormPublishersMapMgr) GetBatchFromNormPublisher(normPublishers []int) (results []*NormPublishersMap, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(NormPublishersMap{}).Where("`norm_publisher` IN (?)", normPublishers).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchUniqueIndexByNormPublishersMapPkindex primary or index 获取唯一内容
func (obj *_NormPublishersMapMgr) FetchUniqueIndexByNormPublishersMapPkindex(manifest int, normPublisher int) (result NormPublishersMap, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(NormPublishersMap{}).Where("`manifest` = ? AND `norm_publisher` = ?", manifest, normPublisher).Find(&result).Error

	return
}

// FetchIndexByNormPublishersMapIndex  获取多个内容
func (obj *_NormPublishersMapMgr) FetchIndexByNormPublishersMapIndex(manifest int) (results []*NormPublishersMap, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(NormPublishersMap{}).Where("`manifest` = ?", manifest).Find(&results).Error

	return
}
