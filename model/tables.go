package model

// Channels [...]
type Channels struct {
	Rowid   int    `gorm:"primaryKey;column:rowid;type:int;not null"`
	Channel string `gorm:"column:channel;type:text;not null"`
}

func (Channels) TableName() string {
	return "channels"
}

// Commands [...]
type Commands struct {
	Rowid   int    `gorm:"primaryKey;column:rowid;type:int;not null"`
	Command string `gorm:"column:command;type:text;not null"`
}

func (Commands) TableName() string {
	return "commands"
}

// CommandsMap [...]
type CommandsMap struct {
	Manifest int `gorm:"uniqueIndex:commands_map_pkindex;column:manifest;type:int;not null"`
	Command  int `gorm:"uniqueIndex:commands_map_pkindex;column:command;type:int;not null"`
}

func (CommandsMap) TableName() string {
	return "commands_map"
}

// IDs [...]
type IDs struct {
	Rowid int    `gorm:"primaryKey;column:rowid;type:int;not null"`
	ID    string `gorm:"column:id;type:text;not null"`
}

func (IDs) TableName() string {
	return "IDs"
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

// Monikers [...]
type Monikers struct {
	Rowid   int    `gorm:"primaryKey;column:rowid;type:int;not null"`
	Moniker string `gorm:"column:moniker;type:text;not null"`
}

func (Monikers) TableName() string {
	return "monikers"
}

// Names [...]
type Names struct {
	Rowid int    `gorm:"primaryKey;column:rowid;type:int;not null"`
	Name  string `gorm:"column:name;type:text;not null"`
}

func (Names) TableName() string {
	return "names"
}

// NormNames [...]
type NormNames struct {
	Rowid    int    `gorm:"primaryKey;column:rowid;type:int;not null"`
	NormName string `gorm:"column:norm_name;type:text;not null"`
}

func (NormNames) TableName() string {
	return "norm_names"
}

// NormNamesMap [...]
type NormNamesMap struct {
	Manifest int `gorm:"uniqueIndex:norm_names_map_pkindex;index:norm_names_map_index;column:manifest;type:int;not null"`
	NormName int `gorm:"uniqueIndex:norm_names_map_pkindex;column:norm_name;type:int;not null"`
}

func (NormNamesMap) TableName() string {
	return "normNamesMap"
}

// NormPublishers [...]
type NormPublishers struct {
	Rowid         int    `gorm:"primaryKey;column:rowid;type:int;not null"`
	NormPublisher string `gorm:"column:norm_publisher;type:text;not null"`
}

func (NormPublishers) TableName() string {
	return "normPublishers"
}

// NormPublishersMap [...]
type NormPublishersMap struct {
	Manifest      int `gorm:"uniqueIndex:norm_publishers_map_pkindex;index:norm_publishers_map_index;column:manifest;type:int;not null"`
	NormPublisher int `gorm:"uniqueIndex:norm_publishers_map_pkindex;column:norm_publisher;type:int;not null"`
}

func (NormPublishersMap) TableName() string {
	return "normPublishers_map"
}

// Pathparts [...]
type Pathparts struct {
	Rowid    int    `gorm:"primaryKey;column:rowid;type:int;not null"`
	Parent   int    `gorm:"column:parent;type:int"`
	Pathpart string `gorm:"column:pathpart;type:text;not null"`
}

func (Pathparts) TableName() string {
	return "pathparts"
}

// Pfns [...]
type Pfns struct {
	Rowid int    `gorm:"primaryKey;column:rowid;type:int;not null"`
	Pfn   string `gorm:"column:pfn;type:text;not null"`
}

func (Pfns) TableName() string {
	return "pfns"
}

// PfnsMap [...]
type PfnsMap struct {
	Manifest int `gorm:"uniqueIndex:pfns_map_pkindex;index:pfns_map_index;column:manifest;type:int;not null"`
	Pfn      int `gorm:"uniqueIndex:pfns_map_pkindex;column:pfn;type:int;not null"`
}

func (PfnsMap) TableName() string {
	return "pfns_map"
}

// Productcodes [...]
type Productcodes struct {
	Rowid       int    `gorm:"primaryKey;column:rowid;type:int;not null"`
	Productcode string `gorm:"column:productcode;type:text;not null"`
}

func (Productcodes) TableName() string {
	return "productcodes"
}

// ProductcodesMap [...]
type ProductcodesMap struct {
	Manifest    int `gorm:"uniqueIndex:productcodes_map_pkindex;index:productcodes_map_index;column:manifest;type:int;not null"`
	Productcode int `gorm:"uniqueIndex:productcodes_map_pkindex;column:productcode;type:int;not null"`
}

func (ProductcodesMap) TableName() string {
	return "productcodes_map"
}

// Tags [...]
type Tags struct {
	Rowid int    `gorm:"primaryKey;column:rowid;type:int;not null"`
	Tag   string `gorm:"column:tag;type:text;not null"`
}

func (Tags) TableName() string {
	return "tags"
}

// TagsMap [...]
type TagsMap struct {
	Manifest int `gorm:"uniqueIndex:tags_map_pkindex;column:manifest;type:int;not null"`
	Tag      int `gorm:"uniqueIndex:tags_map_pkindex;column:tag;type:int;not null"`
}

func (TagsMap) TableName() string {
	return "tags_map"
}

// Versions [...]
type Versions struct {
	Rowid   int    `gorm:"primaryKey;column:rowid;type:int;not null"`
	Version string `gorm:"column:version;type:text;not null"`
}

func (Versions) TableName() string {
	return "versions"
}
