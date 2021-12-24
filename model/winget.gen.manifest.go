package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _ManifestMgr struct {
	*_BaseMgr
}

// ManifestMgr open func
func ManifestMgr(db *gorm.DB) *_ManifestMgr {
	if db == nil {
		panic(fmt.Errorf("ManifestMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ManifestMgr{_BaseMgr: &_BaseMgr{DB: db.Table("manifest"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_ManifestMgr) GetTableName() string {
	return "manifest"
}

// Reset 重置gorm会话
func (obj *_ManifestMgr) Reset() *_ManifestMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_ManifestMgr) Get() (result Manifest, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Manifest{}).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_ManifestMgr) Gets() (results []*Manifest, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Manifest{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_ManifestMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(Manifest{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithRowid rowid获取
func (obj *_ManifestMgr) WithRowid(rowid int) Option {
	return optionFunc(func(o *options) { o.query["rowid"] = rowid })
}

// WithID id获取
func (obj *_ManifestMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithName name获取
func (obj *_ManifestMgr) WithName(name int) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithMoniker moniker获取
func (obj *_ManifestMgr) WithMoniker(moniker int) Option {
	return optionFunc(func(o *options) { o.query["moniker"] = moniker })
}

// WithVersion version获取
func (obj *_ManifestMgr) WithVersion(version int) Option {
	return optionFunc(func(o *options) { o.query["version"] = version })
}

// WithChannel channel获取
func (obj *_ManifestMgr) WithChannel(channel int) Option {
	return optionFunc(func(o *options) { o.query["channel"] = channel })
}

// WithPathpart pathpart获取
func (obj *_ManifestMgr) WithPathpart(pathpart int) Option {
	return optionFunc(func(o *options) { o.query["pathpart"] = pathpart })
}

// WithHash hash获取
func (obj *_ManifestMgr) WithHash(hash []byte) Option {
	return optionFunc(func(o *options) { o.query["hash"] = hash })
}

// GetByOption 功能选项模式获取
func (obj *_ManifestMgr) GetByOption(opts ...Option) (result Manifest, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Manifest{}).Where(options.query).Find(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_ManifestMgr) GetByOptions(opts ...Option) (results []*Manifest, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Manifest{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromRowid 通过rowid获取内容
func (obj *_ManifestMgr) GetFromRowid(rowid int) (result Manifest, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Manifest{}).Where("`rowid` = ?", rowid).Find(&result).Error

	return
}

// GetBatchFromRowid 批量查找
func (obj *_ManifestMgr) GetBatchFromRowid(rowids []int) (results []*Manifest, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Manifest{}).Where("`rowid` IN (?)", rowids).Find(&results).Error

	return
}

// GetFromID 通过id获取内容
func (obj *_ManifestMgr) GetFromID(id int) (results []*Manifest, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Manifest{}).Where("`id` = ?", id).Find(&results).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_ManifestMgr) GetBatchFromID(ids []int) (results []*Manifest, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Manifest{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromName 通过name获取内容
func (obj *_ManifestMgr) GetFromName(name int) (results []*Manifest, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Manifest{}).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找
func (obj *_ManifestMgr) GetBatchFromName(names []int) (results []*Manifest, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Manifest{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromMoniker 通过moniker获取内容
func (obj *_ManifestMgr) GetFromMoniker(moniker int) (results []*Manifest, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Manifest{}).Where("`moniker` = ?", moniker).Find(&results).Error

	return
}

// GetBatchFromMoniker 批量查找
func (obj *_ManifestMgr) GetBatchFromMoniker(monikers []int) (results []*Manifest, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Manifest{}).Where("`moniker` IN (?)", monikers).Find(&results).Error

	return
}

// GetFromVersion 通过version获取内容
func (obj *_ManifestMgr) GetFromVersion(version int) (results []*Manifest, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Manifest{}).Where("`version` = ?", version).Find(&results).Error

	return
}

// GetBatchFromVersion 批量查找
func (obj *_ManifestMgr) GetBatchFromVersion(versions []int) (results []*Manifest, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Manifest{}).Where("`version` IN (?)", versions).Find(&results).Error

	return
}

// GetFromChannel 通过channel获取内容
func (obj *_ManifestMgr) GetFromChannel(channel int) (results []*Manifest, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Manifest{}).Where("`channel` = ?", channel).Find(&results).Error

	return
}

// GetBatchFromChannel 批量查找
func (obj *_ManifestMgr) GetBatchFromChannel(channels []int) (results []*Manifest, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Manifest{}).Where("`channel` IN (?)", channels).Find(&results).Error

	return
}

// GetFromPathpart 通过pathpart获取内容
func (obj *_ManifestMgr) GetFromPathpart(pathpart int) (results []*Manifest, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Manifest{}).Where("`pathpart` = ?", pathpart).Find(&results).Error

	return
}

// GetBatchFromPathpart 批量查找
func (obj *_ManifestMgr) GetBatchFromPathpart(pathparts []int) (results []*Manifest, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Manifest{}).Where("`pathpart` IN (?)", pathparts).Find(&results).Error

	return
}

// GetFromHash 通过hash获取内容
func (obj *_ManifestMgr) GetFromHash(hash []byte) (results []*Manifest, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Manifest{}).Where("`hash` = ?", hash).Find(&results).Error

	return
}

// GetBatchFromHash 批量查找
func (obj *_ManifestMgr) GetBatchFromHash(hashs [][]byte) (results []*Manifest, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Manifest{}).Where("`hash` IN (?)", hashs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_ManifestMgr) FetchByPrimaryKey(rowid int) (result Manifest, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Manifest{}).Where("`rowid` = ?", rowid).Find(&result).Error

	return
}

// FetchIndexByManifestIDIndex  获取多个内容
func (obj *_ManifestMgr) FetchIndexByManifestIDIndex(id int) (results []*Manifest, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Manifest{}).Where("`id` = ?", id).Find(&results).Error

	return
}

// FetchIndexByManifestNameIndex  获取多个内容
func (obj *_ManifestMgr) FetchIndexByManifestNameIndex(name int) (results []*Manifest, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Manifest{}).Where("`name` = ?", name).Find(&results).Error

	return
}

// FetchIndexByManifestMonikerIndex  获取多个内容
func (obj *_ManifestMgr) FetchIndexByManifestMonikerIndex(moniker int) (results []*Manifest, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Manifest{}).Where("`moniker` = ?", moniker).Find(&results).Error

	return
}
