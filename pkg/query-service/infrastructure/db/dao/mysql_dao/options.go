package mysql_dao

import (
	"gorm.io/gorm"
	"time"
)

type Options struct {
	Db      *gorm.DB
	DbId    string
	MaxTime *time.Duration
	Sort    *string
}

func NewOptions() *Options {
	return &Options{}
}

func (o *Options) SetDB(db *gorm.DB) *Options {
	o.Db = db
	return o
}

func (o *Options) SetDbId(v string) *Options {
	o.DbId = v
	return o
}

func (o *Options) GetDbId() string {
	return o.DbId
}

func (o *Options) SetSort(v *string) *Options {
	o.Sort = v
	return o
}

func (o *Options) GetSort() interface{} {
	return o.Sort
}

func (o *Options) Merge(opts ...*Options) *Options {
	for _, item := range opts {
		if item == nil {
			continue
		}
		if item.Db != nil {
			o.Db = item.Db
		}
		if len(item.DbId) != 0 {
			o.DbId = item.DbId
		}
		if item.MaxTime != nil {
			o.MaxTime = item.MaxTime
		}
		if item.Sort != nil {
			o.Sort = item.Sort
		}
	}
	if o.Db == nil {
		o.Db = GetMySqlDB()
	}
	return o
}
