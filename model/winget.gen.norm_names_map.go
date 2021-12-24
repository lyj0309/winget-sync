package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _NormNamesMapMgr struct {
	*_BaseMgr
}

// NormNamesMapMgr open func
func NormNamesMapMgr(db *gorm.DB) *_NormNamesMapMgr {
	if db == nil {
		panic(fmt.Errorf("NormNamesMapMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_NormNamesMapMgr{_BaseMgr: &_BaseMgr{DB: db.Table("norm_names_map"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_NormNamesMapMgr) GetTableName() string {
	return "norm_names_map"
}

// Reset 重置gorm会话
func (obj *_NormNamesMapMgr) Reset() *_NormNamesMapMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_NormNamesMapMgr) Get() (result NormNamesMap, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(NormNamesMap{}).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_NormNamesMapMgr) Gets() (results []*NormNamesMap, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(NormNamesMap{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_NormNamesMapMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(NormNamesMap{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithManifest manifest获取
func (obj *_NormNamesMapMgr) WithManifest(manifest int) Option {
	return optionFunc(func(o *options) { o.query["manifest"] = manifest })
}

// WithNormName norm_name获取
func (obj *_NormNamesMapMgr) WithNormName(normName int) Option {
	return optionFunc(func(o *options) { o.query["norm_name"] = normName })
}

// GetByOption 功能选项模式获取
func (obj *_NormNamesMapMgr) GetByOption(opts ...Option) (result NormNamesMap, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(NormNamesMap{}).Where(options.query).Find(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_NormNamesMapMgr) GetByOptions(opts ...Option) (results []*NormNamesMap, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(NormNamesMap{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromManifest 通过manifest获取内容
func (obj *_NormNamesMapMgr) GetFromManifest(manifest int) (results []*NormNamesMap, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(NormNamesMap{}).Where("`manifest` = ?", manifest).Find(&results).Error

	return
}

// GetBatchFromManifest 批量查找
func (obj *_NormNamesMapMgr) GetBatchFromManifest(manifests []int) (results []*NormNamesMap, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(NormNamesMap{}).Where("`manifest` IN (?)", manifests).Find(&results).Error

	return
}

// GetFromNormName 通过norm_name获取内容
func (obj *_NormNamesMapMgr) GetFromNormName(normName int) (results []*NormNamesMap, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(NormNamesMap{}).Where("`norm_name` = ?", normName).Find(&results).Error

	return
}

// GetBatchFromNormName 批量查找
func (obj *_NormNamesMapMgr) GetBatchFromNormName(normNames []int) (results []*NormNamesMap, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(NormNamesMap{}).Where("`norm_name` IN (?)", normNames).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchUniqueIndexByNormNamesMapPkindex primary or index 获取唯一内容
func (obj *_NormNamesMapMgr) FetchUniqueIndexByNormNamesMapPkindex(manifest int, normName int) (result NormNamesMap, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(NormNamesMap{}).Where("`manifest` = ? AND `norm_name` = ?", manifest, normName).Find(&result).Error

	return
}

// FetchIndexByNormNamesMapIndex  获取多个内容
func (obj *_NormNamesMapMgr) FetchIndexByNormNamesMapIndex(manifest int) (results []*NormNamesMap, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(NormNamesMap{}).Where("`manifest` = ?", manifest).Find(&results).Error

	return
}
