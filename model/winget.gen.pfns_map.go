package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _PfnsMapMgr struct {
	*_BaseMgr
}

// PfnsMapMgr open func
func PfnsMapMgr(db *gorm.DB) *_PfnsMapMgr {
	if db == nil {
		panic(fmt.Errorf("PfnsMapMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_PfnsMapMgr{_BaseMgr: &_BaseMgr{DB: db.Table("pfns_map"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_PfnsMapMgr) GetTableName() string {
	return "pfns_map"
}

// Reset 重置gorm会话
func (obj *_PfnsMapMgr) Reset() *_PfnsMapMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_PfnsMapMgr) Get() (result PfnsMap, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PfnsMap{}).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_PfnsMapMgr) Gets() (results []*PfnsMap, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PfnsMap{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_PfnsMapMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(PfnsMap{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithManifest manifest获取
func (obj *_PfnsMapMgr) WithManifest(manifest int) Option {
	return optionFunc(func(o *options) { o.query["manifest"] = manifest })
}

// WithPfn pfn获取
func (obj *_PfnsMapMgr) WithPfn(pfn int) Option {
	return optionFunc(func(o *options) { o.query["pfn"] = pfn })
}

// GetByOption 功能选项模式获取
func (obj *_PfnsMapMgr) GetByOption(opts ...Option) (result PfnsMap, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(PfnsMap{}).Where(options.query).Find(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_PfnsMapMgr) GetByOptions(opts ...Option) (results []*PfnsMap, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(PfnsMap{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromManifest 通过manifest获取内容
func (obj *_PfnsMapMgr) GetFromManifest(manifest int) (results []*PfnsMap, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PfnsMap{}).Where("`manifest` = ?", manifest).Find(&results).Error

	return
}

// GetBatchFromManifest 批量查找
func (obj *_PfnsMapMgr) GetBatchFromManifest(manifests []int) (results []*PfnsMap, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PfnsMap{}).Where("`manifest` IN (?)", manifests).Find(&results).Error

	return
}

// GetFromPfn 通过pfn获取内容
func (obj *_PfnsMapMgr) GetFromPfn(pfn int) (results []*PfnsMap, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PfnsMap{}).Where("`pfn` = ?", pfn).Find(&results).Error

	return
}

// GetBatchFromPfn 批量查找
func (obj *_PfnsMapMgr) GetBatchFromPfn(pfns []int) (results []*PfnsMap, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PfnsMap{}).Where("`pfn` IN (?)", pfns).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchUniqueIndexByPfnsMapPkindex primary or index 获取唯一内容
func (obj *_PfnsMapMgr) FetchUniqueIndexByPfnsMapPkindex(manifest int, pfn int) (result PfnsMap, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PfnsMap{}).Where("`manifest` = ? AND `pfn` = ?", manifest, pfn).Find(&result).Error

	return
}

// FetchIndexByPfnsMapIndex  获取多个内容
func (obj *_PfnsMapMgr) FetchIndexByPfnsMapIndex(manifest int) (results []*PfnsMap, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PfnsMap{}).Where("`manifest` = ?", manifest).Find(&results).Error

	return
}
