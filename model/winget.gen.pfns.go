package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _PfnsMgr struct {
	*_BaseMgr
}

// PfnsMgr open func
func PfnsMgr(db *gorm.DB) *_PfnsMgr {
	if db == nil {
		panic(fmt.Errorf("PfnsMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_PfnsMgr{_BaseMgr: &_BaseMgr{DB: db.Table("pfns"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_PfnsMgr) GetTableName() string {
	return "pfns"
}

// Reset 重置gorm会话
func (obj *_PfnsMgr) Reset() *_PfnsMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_PfnsMgr) Get() (result Pfns, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Pfns{}).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_PfnsMgr) Gets() (results []*Pfns, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Pfns{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_PfnsMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(Pfns{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithRowid rowid获取
func (obj *_PfnsMgr) WithRowid(rowid int) Option {
	return optionFunc(func(o *options) { o.query["rowid"] = rowid })
}

// WithPfn pfn获取
func (obj *_PfnsMgr) WithPfn(pfn string) Option {
	return optionFunc(func(o *options) { o.query["pfn"] = pfn })
}

// GetByOption 功能选项模式获取
func (obj *_PfnsMgr) GetByOption(opts ...Option) (result Pfns, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Pfns{}).Where(options.query).Find(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_PfnsMgr) GetByOptions(opts ...Option) (results []*Pfns, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Pfns{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromRowid 通过rowid获取内容
func (obj *_PfnsMgr) GetFromRowid(rowid int) (result Pfns, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Pfns{}).Where("`rowid` = ?", rowid).Find(&result).Error

	return
}

// GetBatchFromRowid 批量查找
func (obj *_PfnsMgr) GetBatchFromRowid(rowids []int) (results []*Pfns, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Pfns{}).Where("`rowid` IN (?)", rowids).Find(&results).Error

	return
}

// GetFromPfn 通过pfn获取内容
func (obj *_PfnsMgr) GetFromPfn(pfn string) (results []*Pfns, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Pfns{}).Where("`pfn` = ?", pfn).Find(&results).Error

	return
}

// GetBatchFromPfn 批量查找
func (obj *_PfnsMgr) GetBatchFromPfn(pfns []string) (results []*Pfns, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Pfns{}).Where("`pfn` IN (?)", pfns).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_PfnsMgr) FetchByPrimaryKey(rowid int) (result Pfns, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Pfns{}).Where("`rowid` = ?", rowid).Find(&result).Error

	return
}
