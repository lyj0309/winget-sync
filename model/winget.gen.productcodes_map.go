package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _ProductcodesMapMgr struct {
	*_BaseMgr
}

// ProductcodesMapMgr open func
func ProductcodesMapMgr(db *gorm.DB) *_ProductcodesMapMgr {
	if db == nil {
		panic(fmt.Errorf("ProductcodesMapMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ProductcodesMapMgr{_BaseMgr: &_BaseMgr{DB: db.Table("productcodes_map"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_ProductcodesMapMgr) GetTableName() string {
	return "productcodes_map"
}

// Reset 重置gorm会话
func (obj *_ProductcodesMapMgr) Reset() *_ProductcodesMapMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_ProductcodesMapMgr) Get() (result ProductcodesMap, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ProductcodesMap{}).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_ProductcodesMapMgr) Gets() (results []*ProductcodesMap, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ProductcodesMap{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_ProductcodesMapMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(ProductcodesMap{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithManifest manifest获取
func (obj *_ProductcodesMapMgr) WithManifest(manifest int) Option {
	return optionFunc(func(o *options) { o.query["manifest"] = manifest })
}

// WithProductcode productcode获取
func (obj *_ProductcodesMapMgr) WithProductcode(productcode int) Option {
	return optionFunc(func(o *options) { o.query["productcode"] = productcode })
}

// GetByOption 功能选项模式获取
func (obj *_ProductcodesMapMgr) GetByOption(opts ...Option) (result ProductcodesMap, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(ProductcodesMap{}).Where(options.query).Find(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_ProductcodesMapMgr) GetByOptions(opts ...Option) (results []*ProductcodesMap, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(ProductcodesMap{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromManifest 通过manifest获取内容
func (obj *_ProductcodesMapMgr) GetFromManifest(manifest int) (results []*ProductcodesMap, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ProductcodesMap{}).Where("`manifest` = ?", manifest).Find(&results).Error

	return
}

// GetBatchFromManifest 批量查找
func (obj *_ProductcodesMapMgr) GetBatchFromManifest(manifests []int) (results []*ProductcodesMap, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ProductcodesMap{}).Where("`manifest` IN (?)", manifests).Find(&results).Error

	return
}

// GetFromProductcode 通过productcode获取内容
func (obj *_ProductcodesMapMgr) GetFromProductcode(productcode int) (results []*ProductcodesMap, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ProductcodesMap{}).Where("`productcode` = ?", productcode).Find(&results).Error

	return
}

// GetBatchFromProductcode 批量查找
func (obj *_ProductcodesMapMgr) GetBatchFromProductcode(productcodes []int) (results []*ProductcodesMap, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ProductcodesMap{}).Where("`productcode` IN (?)", productcodes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchUniqueIndexByProductcodesMapPkindex primary or index 获取唯一内容
func (obj *_ProductcodesMapMgr) FetchUniqueIndexByProductcodesMapPkindex(manifest int, productcode int) (result ProductcodesMap, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ProductcodesMap{}).Where("`manifest` = ? AND `productcode` = ?", manifest, productcode).Find(&result).Error

	return
}

// FetchIndexByProductcodesMapIndex  获取多个内容
func (obj *_ProductcodesMapMgr) FetchIndexByProductcodesMapIndex(manifest int) (results []*ProductcodesMap, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ProductcodesMap{}).Where("`manifest` = ?", manifest).Find(&results).Error

	return
}
