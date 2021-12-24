package model

// Channels [...]
type Channels struct {
	Rowid   int    `gorm:"primaryKey;column:rowid;type:int;not null"`
	Channel string `gorm:"column:channel;type:text;not null"`
}

func (Channels) TableName() string {
	return "channels"
}

// ChannelsColumns get sql column name.获取数据库列名
var ChannelsColumns = struct {
	Rowid   string
	Channel string
}{
	Rowid:   "rowid",
	Channel: "channel",
}

// Commands [...]
type Commands struct {
	Rowid   int    `gorm:"primaryKey;column:rowid;type:int;not null"`
	Command string `gorm:"column:command;type:text;not null"`
}

func (Commands) TableName() string {
	return "commands"
}

// CommandsColumns get sql column name.获取数据库列名
var CommandsColumns = struct {
	Rowid   string
	Command string
}{
	Rowid:   "rowid",
	Command: "command",
}

// CommandsMap [...]
type CommandsMap struct {
	Manifest int `gorm:"uniqueIndex:commands_map_pkindex;column:manifest;type:int;not null"`
	Command  int `gorm:"uniqueIndex:commands_map_pkindex;column:command;type:int;not null"`
}

func (CommandsMap) TableName() string {
	return "commandsMap"
}

// CommandsMapColumns get sql column name.获取数据库列名
var CommandsMapColumns = struct {
	Manifest string
	Command  string
}{
	Manifest: "manifest",
	Command:  "command",
}

// IDs [...]
type IDs struct {
	Rowid int    `gorm:"primaryKey;column:rowid;type:int;not null"`
	ID    string `gorm:"column:id;type:text;not null"`
}

func (IDs) TableName() string {
	return "IDs"
}

// IDsColumns get sql column name.获取数据库列名
var IDsColumns = struct {
	Rowid string
	ID    string
}{
	Rowid: "rowid",
	ID:    "id",
}

// Manifest [...]
type Manifest struct {
	Rowid    int    `gorm:"primaryKey;column:rowid;type:int;not null"`
	ID       int    `gorm:"index:manifest_id_index;column:id;type:int;not null"`
	Name     int    `gorm:"index:manifest_name_index;column:name;type:int;not null"`
	Moniker  int    `gorm:"index:manifest_moniker_index;column:moniker;type:int;not null"`
	Version  int    `gorm:"column:version;type:int;not null"`
	Channel  int    `gorm:"column:channel;type:int;not null"`
	Pathpart int    `gorm:"column:pathpart;type:int;not null"`
	Hash     []byte `gorm:"column:hash;type:blob"`
}

func (Manifest) TableName() string {
	return "manifest"
}

// ManifestColumns get sql column name.获取数据库列名
var ManifestColumns = struct {
	Rowid    string
	ID       string
	Name     string
	Moniker  string
	Version  string
	Channel  string
	Pathpart string
	Hash     string
}{
	Rowid:    "rowid",
	ID:       "id",
	Name:     "name",
	Moniker:  "moniker",
	Version:  "version",
	Channel:  "channel",
	Pathpart: "pathpart",
	Hash:     "hash",
}

// Monikers [...]
type Monikers struct {
	Rowid   int    `gorm:"primaryKey;column:rowid;type:int;not null"`
	Moniker string `gorm:"column:moniker;type:text;not null"`
}

func (Monikers) TableName() string {
	return "monikers"
}

// MonikersColumns get sql column name.获取数据库列名
var MonikersColumns = struct {
	Rowid   string
	Moniker string
}{
	Rowid:   "rowid",
	Moniker: "moniker",
}

// Names [...]
type Names struct {
	Rowid int    `gorm:"primaryKey;column:rowid;type:int;not null"`
	Name  string `gorm:"column:name;type:text;not null"`
}

func (Names) TableName() string {
	return "names"
}

// NamesColumns get sql column name.获取数据库列名
var NamesColumns = struct {
	Rowid string
	Name  string
}{
	Rowid: "rowid",
	Name:  "name",
}

// NormNames [...]
type NormNames struct {
	Rowid    int    `gorm:"primaryKey;column:rowid;type:int;not null"`
	NormName string `gorm:"column:norm_name;type:text;not null"`
}

func (NormNames) TableName() string {
	return "normNames"
}

// NormNamesColumns get sql column name.获取数据库列名
var NormNamesColumns = struct {
	Rowid    string
	NormName string
}{
	Rowid:    "rowid",
	NormName: "norm_name",
}

// NormNamesMap [...]
type NormNamesMap struct {
	Manifest int `gorm:"uniqueIndex:norm_names_map_pkindex;index:norm_names_map_index;column:manifest;type:int;not null"`
	NormName int `gorm:"uniqueIndex:norm_names_map_pkindex;column:norm_name;type:int;not null"`
}

func (NormNamesMap) TableName() string {
	return "normNamesMap"
}

// NormNamesMapColumns get sql column name.获取数据库列名
var NormNamesMapColumns = struct {
	Manifest string
	NormName string
}{
	Manifest: "manifest",
	NormName: "norm_name",
}

// NormPublishers [...]
type NormPublishers struct {
	Rowid         int    `gorm:"primaryKey;column:rowid;type:int;not null"`
	NormPublisher string `gorm:"column:norm_publisher;type:text;not null"`
}

// NormPublishersColumns get sql column name.获取数据库列名
var NormPublishersColumns = struct {
	Rowid         string
	NormPublisher string
}{
	Rowid:         "rowid",
	NormPublisher: "norm_publisher",
}

// NormPublishersMap [...]
type NormPublishersMap struct {
	Manifest      int `gorm:"uniqueIndex:norm_publishers_map_pkindex;index:norm_publishers_map_index;column:manifest;type:int;not null"`
	NormPublisher int `gorm:"uniqueIndex:norm_publishers_map_pkindex;column:norm_publisher;type:int;not null"`
}

// NormPublishersMapColumns get sql column name.获取数据库列名
var NormPublishersMapColumns = struct {
	Manifest      string
	NormPublisher string
}{
	Manifest:      "manifest",
	NormPublisher: "norm_publisher",
}

// Pathparts [...]
type Pathparts struct {
	Rowid    int    `gorm:"primaryKey;column:rowid;type:int;not null"`
	Parent   int    `gorm:"column:parent;type:int"`
	Pathpart string `gorm:"column:pathpart;type:text;not null"`
}

// PathpartsColumns get sql column name.获取数据库列名
var PathpartsColumns = struct {
	Rowid    string
	Parent   string
	Pathpart string
}{
	Rowid:    "rowid",
	Parent:   "parent",
	Pathpart: "pathpart",
}

// Pfns [...]
type Pfns struct {
	Rowid int    `gorm:"primaryKey;column:rowid;type:int;not null"`
	Pfn   string `gorm:"column:pfn;type:text;not null"`
}

// PfnsColumns get sql column name.获取数据库列名
var PfnsColumns = struct {
	Rowid string
	Pfn   string
}{
	Rowid: "rowid",
	Pfn:   "pfn",
}

// PfnsMap [...]
type PfnsMap struct {
	Manifest int `gorm:"uniqueIndex:pfns_map_pkindex;index:pfns_map_index;column:manifest;type:int;not null"`
	Pfn      int `gorm:"uniqueIndex:pfns_map_pkindex;column:pfn;type:int;not null"`
}

// PfnsMapColumns get sql column name.获取数据库列名
var PfnsMapColumns = struct {
	Manifest string
	Pfn      string
}{
	Manifest: "manifest",
	Pfn:      "pfn",
}

// Productcodes [...]
type Productcodes struct {
	Rowid       int    `gorm:"primaryKey;column:rowid;type:int;not null"`
	Productcode string `gorm:"column:productcode;type:text;not null"`
}

// ProductcodesColumns get sql column name.获取数据库列名
var ProductcodesColumns = struct {
	Rowid       string
	Productcode string
}{
	Rowid:       "rowid",
	Productcode: "productcode",
}

// ProductcodesMap [...]
type ProductcodesMap struct {
	Manifest    int `gorm:"uniqueIndex:productcodes_map_pkindex;index:productcodes_map_index;column:manifest;type:int;not null"`
	Productcode int `gorm:"uniqueIndex:productcodes_map_pkindex;column:productcode;type:int;not null"`
}

// ProductcodesMapColumns get sql column name.获取数据库列名
var ProductcodesMapColumns = struct {
	Manifest    string
	Productcode string
}{
	Manifest:    "manifest",
	Productcode: "productcode",
}

// Tags [...]
type Tags struct {
	Rowid int    `gorm:"primaryKey;column:rowid;type:int;not null"`
	Tag   string `gorm:"column:tag;type:text;not null"`
}

// TagsColumns get sql column name.获取数据库列名
var TagsColumns = struct {
	Rowid string
	Tag   string
}{
	Rowid: "rowid",
	Tag:   "tag",
}

// TagsMap [...]
type TagsMap struct {
	Manifest int `gorm:"uniqueIndex:tags_map_pkindex;column:manifest;type:int;not null"`
	Tag      int `gorm:"uniqueIndex:tags_map_pkindex;column:tag;type:int;not null"`
}

// TagsMapColumns get sql column name.获取数据库列名
var TagsMapColumns = struct {
	Manifest string
	Tag      string
}{
	Manifest: "manifest",
	Tag:      "tag",
}

// Versions [...]
type Versions struct {
	Rowid   int    `gorm:"primaryKey;column:rowid;type:int;not null"`
	Version string `gorm:"column:version;type:text;not null"`
}

// VersionsColumns get sql column name.获取数据库列名
var VersionsColumns = struct {
	Rowid   string
	Version string
}{
	Rowid:   "rowid",
	Version: "version",
}
