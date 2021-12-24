package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _TagsMgr struct {
	*_BaseMgr
}

// TagsMgr open func
func TagsMgr(db *gorm.DB) *_TagsMgr {
	if db == nil {
		panic(fmt.Errorf("TagsMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_TagsMgr{_BaseMgr: &_BaseMgr{DB: db.Table("tags"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_TagsMgr) GetTableName() string {
	return "tags"
}

// Reset 重置gorm会话
func (obj *_TagsMgr) Reset() *_TagsMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_TagsMgr) Get() (result Tags, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Tags{}).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_TagsMgr) Gets() (results []*Tags, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Tags{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_TagsMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(Tags{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithRowid rowid获取
func (obj *_TagsMgr) WithRowid(rowid int) Option {
	return optionFunc(func(o *options) { o.query["rowid"] = rowid })
}

// WithTag tag获取
func (obj *_TagsMgr) WithTag(tag string) Option {
	return optionFunc(func(o *options) { o.query["tag"] = tag })
}

// GetByOption 功能选项模式获取
func (obj *_TagsMgr) GetByOption(opts ...Option) (result Tags, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Tags{}).Where(options.query).Find(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_TagsMgr) GetByOptions(opts ...Option) (results []*Tags, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Tags{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromRowid 通过rowid获取内容
func (obj *_TagsMgr) GetFromRowid(rowid int) (result Tags, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Tags{}).Where("`rowid` = ?", rowid).Find(&result).Error

	return
}

// GetBatchFromRowid 批量查找
func (obj *_TagsMgr) GetBatchFromRowid(rowids []int) (results []*Tags, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Tags{}).Where("`rowid` IN (?)", rowids).Find(&results).Error

	return
}

// GetFromTag 通过tag获取内容
func (obj *_TagsMgr) GetFromTag(tag string) (results []*Tags, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Tags{}).Where("`tag` = ?", tag).Find(&results).Error

	return
}

// GetBatchFromTag 批量查找
func (obj *_TagsMgr) GetBatchFromTag(tags []string) (results []*Tags, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Tags{}).Where("`tag` IN (?)", tags).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_TagsMgr) FetchByPrimaryKey(rowid int) (result Tags, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Tags{}).Where("`rowid` = ?", rowid).Find(&result).Error

	return
}
