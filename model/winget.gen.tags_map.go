package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _TagsMapMgr struct {
	*_BaseMgr
}

// TagsMapMgr open func
func TagsMapMgr(db *gorm.DB) *_TagsMapMgr {
	if db == nil {
		panic(fmt.Errorf("TagsMapMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_TagsMapMgr{_BaseMgr: &_BaseMgr{DB: db.Table("tags_map"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_TagsMapMgr) GetTableName() string {
	return "tags_map"
}

// Reset 重置gorm会话
func (obj *_TagsMapMgr) Reset() *_TagsMapMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_TagsMapMgr) Get() (result TagsMap, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TagsMap{}).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_TagsMapMgr) Gets() (results []*TagsMap, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TagsMap{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_TagsMapMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(TagsMap{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithManifest manifest获取
func (obj *_TagsMapMgr) WithManifest(manifest int) Option {
	return optionFunc(func(o *options) { o.query["manifest"] = manifest })
}

// WithTag tag获取
func (obj *_TagsMapMgr) WithTag(tag int) Option {
	return optionFunc(func(o *options) { o.query["tag"] = tag })
}

// GetByOption 功能选项模式获取
func (obj *_TagsMapMgr) GetByOption(opts ...Option) (result TagsMap, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(TagsMap{}).Where(options.query).Find(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_TagsMapMgr) GetByOptions(opts ...Option) (results []*TagsMap, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(TagsMap{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromManifest 通过manifest获取内容
func (obj *_TagsMapMgr) GetFromManifest(manifest int) (results []*TagsMap, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TagsMap{}).Where("`manifest` = ?", manifest).Find(&results).Error

	return
}

// GetBatchFromManifest 批量查找
func (obj *_TagsMapMgr) GetBatchFromManifest(manifests []int) (results []*TagsMap, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TagsMap{}).Where("`manifest` IN (?)", manifests).Find(&results).Error

	return
}

// GetFromTag 通过tag获取内容
func (obj *_TagsMapMgr) GetFromTag(tag int) (results []*TagsMap, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TagsMap{}).Where("`tag` = ?", tag).Find(&results).Error

	return
}

// GetBatchFromTag 批量查找
func (obj *_TagsMapMgr) GetBatchFromTag(tags []int) (results []*TagsMap, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TagsMap{}).Where("`tag` IN (?)", tags).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchUniqueIndexByTagsMapPkindex primary or index 获取唯一内容
func (obj *_TagsMapMgr) FetchUniqueIndexByTagsMapPkindex(manifest int, tag int) (result TagsMap, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TagsMap{}).Where("`manifest` = ? AND `tag` = ?", manifest, tag).Find(&result).Error

	return
}
