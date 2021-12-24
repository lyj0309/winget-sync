package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _ChannelsMgr struct {
	*_BaseMgr
}

// ChannelsMgr open func
func ChannelsMgr(db *gorm.DB) *_ChannelsMgr {
	if db == nil {
		panic(fmt.Errorf("ChannelsMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ChannelsMgr{_BaseMgr: &_BaseMgr{DB: db.Table("channels"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_ChannelsMgr) GetTableName() string {
	return "channels"
}

// Reset 重置gorm会话
func (obj *_ChannelsMgr) Reset() *_ChannelsMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_ChannelsMgr) Get() (result Channels, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Channels{}).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_ChannelsMgr) Gets() (results []*Channels, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Channels{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_ChannelsMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(Channels{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithRowid rowid获取
func (obj *_ChannelsMgr) WithRowid(rowid int) Option {
	return optionFunc(func(o *options) { o.query["rowid"] = rowid })
}

// WithChannel channel获取
func (obj *_ChannelsMgr) WithChannel(channel string) Option {
	return optionFunc(func(o *options) { o.query["channel"] = channel })
}

// GetByOption 功能选项模式获取
func (obj *_ChannelsMgr) GetByOption(opts ...Option) (result Channels, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Channels{}).Where(options.query).Find(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_ChannelsMgr) GetByOptions(opts ...Option) (results []*Channels, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Channels{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromRowid 通过rowid获取内容
func (obj *_ChannelsMgr) GetFromRowid(rowid int) (result Channels, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Channels{}).Where("`rowid` = ?", rowid).Find(&result).Error

	return
}

// GetBatchFromRowid 批量查找
func (obj *_ChannelsMgr) GetBatchFromRowid(rowids []int) (results []*Channels, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Channels{}).Where("`rowid` IN (?)", rowids).Find(&results).Error

	return
}

// GetFromChannel 通过channel获取内容
func (obj *_ChannelsMgr) GetFromChannel(channel string) (results []*Channels, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Channels{}).Where("`channel` = ?", channel).Find(&results).Error

	return
}

// GetBatchFromChannel 批量查找
func (obj *_ChannelsMgr) GetBatchFromChannel(channels []string) (results []*Channels, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Channels{}).Where("`channel` IN (?)", channels).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_ChannelsMgr) FetchByPrimaryKey(rowid int) (result Channels, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Channels{}).Where("`rowid` = ?", rowid).Find(&result).Error

	return
}
