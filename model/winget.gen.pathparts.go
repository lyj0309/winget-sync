package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _PathpartsMgr struct {
	*_BaseMgr
}

// PathpartsMgr open func
func PathpartsMgr(db *gorm.DB) *_PathpartsMgr {
	if db == nil {
		panic(fmt.Errorf("PathpartsMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_PathpartsMgr{_BaseMgr: &_BaseMgr{DB: db.Table("pathparts"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_PathpartsMgr) GetTableName() string {
	return "pathparts"
}

// Reset 重置gorm会话
func (obj *_PathpartsMgr) Reset() *_PathpartsMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_PathpartsMgr) Get() (result Pathparts, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Pathparts{}).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_PathpartsMgr) Gets() (results []*Pathparts, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Pathparts{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_PathpartsMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(Pathparts{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithRowid rowid获取
func (obj *_PathpartsMgr) WithRowid(rowid int) Option {
	return optionFunc(func(o *options) { o.query["rowid"] = rowid })
}

// WithParent parent获取
func (obj *_PathpartsMgr) WithParent(parent int) Option {
	return optionFunc(func(o *options) { o.query["parent"] = parent })
}

// WithPathpart pathpart获取
func (obj *_PathpartsMgr) WithPathpart(pathpart string) Option {
	return optionFunc(func(o *options) { o.query["pathpart"] = pathpart })
}

// GetByOption 功能选项模式获取
func (obj *_PathpartsMgr) GetByOption(opts ...Option) (result Pathparts, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Pathparts{}).Where(options.query).Find(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_PathpartsMgr) GetByOptions(opts ...Option) (results []*Pathparts, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Pathparts{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromRowid 通过rowid获取内容
func (obj *_PathpartsMgr) GetFromRowid(rowid int) (result Pathparts, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Pathparts{}).Where("`rowid` = ?", rowid).Find(&result).Error

	return
}

// GetBatchFromRowid 批量查找
func (obj *_PathpartsMgr) GetBatchFromRowid(rowids []int) (results []*Pathparts, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Pathparts{}).Where("`rowid` IN (?)", rowids).Find(&results).Error

	return
}

// GetFromParent 通过parent获取内容
func (obj *_PathpartsMgr) GetFromParent(parent int) (results []*Pathparts, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Pathparts{}).Where("`parent` = ?", parent).Find(&results).Error

	return
}

// GetBatchFromParent 批量查找
func (obj *_PathpartsMgr) GetBatchFromParent(parents []int) (results []*Pathparts, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Pathparts{}).Where("`parent` IN (?)", parents).Find(&results).Error

	return
}

// GetFromPathpart 通过pathpart获取内容
func (obj *_PathpartsMgr) GetFromPathpart(pathpart string) (results []*Pathparts, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Pathparts{}).Where("`pathpart` = ?", pathpart).Find(&results).Error

	return
}

// GetBatchFromPathpart 批量查找
func (obj *_PathpartsMgr) GetBatchFromPathpart(pathparts []string) (results []*Pathparts, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Pathparts{}).Where("`pathpart` IN (?)", pathparts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_PathpartsMgr) FetchByPrimaryKey(rowid int) (result Pathparts, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Pathparts{}).Where("`rowid` = ?", rowid).Find(&result).Error

	return
}
