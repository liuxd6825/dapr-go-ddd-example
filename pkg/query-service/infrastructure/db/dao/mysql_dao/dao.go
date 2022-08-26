package mysql_dao

import (
	"context"
	"errors"
	"fmt"
	"github.com/liuxd6825/components-contrib/liuxd/eventstorage/impl/gorm_impl/db"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd/ddd_repository"
	"gorm.io/gorm"
	"strings"
)

const (
	WhereTenantId      = "tenant_id=%v"
	WhereTenantIdAndId = "tenant_id=%v and id=%v"
)

type Dao[T ddd.Entity] struct {
	db          *gorm.DB
	newFunc     func() T
	newListFunc func() []T
}

/*func NewDao2[T ddd.Entity](collectionName string, opts ...*Options) *Dao[T] {
	options := NewRepositoryOptions()
	options.Merge(opts...)
	coll := options.mongoDB.GetCollection(collectionName)
	return &Dao[T]{
		Dao: ddd_mongodb.NewDao[T](options.mongoDB, coll),
	}
}*/

func NewDao[T ddd.Entity](db *gorm.DB, newFunc func() T, newListFunc func() []T) *Dao[T] {
	dao := Dao[T]{
		db:          db,
		newFunc:     newFunc,
		newListFunc: newListFunc,
	}
	return &dao
}

func (d *Dao[T]) Insert(ctx context.Context, v T, opts ...*ddd_repository.SetOptions) error {
	return d.getDB(ctx).Model(v).Create(v).Error
}

func (d *Dao[T]) InsertMany(ctx context.Context, vList []T, opts ...*ddd_repository.SetOptions) error {
	v := d.NewEntity()
	return d.getDB(ctx).Model(v).CreateInBatches(vList, len(vList)).Error
}

func (d *Dao[T]) Update(ctx context.Context, v T, opts ...*ddd_repository.SetOptions) error {
	return d.getDB(ctx).Where(WhereTenantIdAndId, v.GetTenantId(), v.GetId()).Updates(v).Error
}

func (d *Dao[T]) UpdateMany(ctx context.Context, vList []T, opts ...*ddd_repository.SetOptions) error {
	getDB := d.getDB(ctx)
	var model T
	for _, item := range vList {
		model = d.NewEntity()
		model.SetId(item.GetId())
		model.SetTenantId(item.GetTenantId())
		err := getDB.Model(&model).Updates(item).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func (d *Dao[T]) UpdateManyByFilter(ctx context.Context, tenantId, filter string, data interface{}, opts ...*ddd_repository.SetOptions) error {
	where, err := getSqlWhere(tenantId, filter)
	if err != nil {
		return err
	}
	return d.getDB(ctx).Where(where).Updates(data).Error
}

func (d *Dao[T]) UpdateByMap(ctx context.Context, tenantId, id string, data map[string]interface{}, opts ...*Options) error {
	var v T = d.NewEntity()
	return d.getDB(ctx).Model(v).Where(WhereTenantIdAndId, tenantId, id).Updates(data).Error
}

func (d *Dao[T]) DeleteById(ctx context.Context, tenantId string, id string, opts ...*ddd_repository.SetOptions) error {
	var model T
	model.SetId(id)
	model.SetTenantId(tenantId)
	return d.getDB(ctx).Delete(&model).Error
}

func (d *Dao[T]) DeleteByIds(ctx context.Context, tenantId string, ids []string, opts ...*ddd_repository.SetOptions) error {
	var models []T
	models = d.NewEntities()
	return d.getDB(ctx).Where(WhereTenantId, tenantId).Delete(&models, ids).Error
}

func (d *Dao[T]) DeleteAll(ctx context.Context, tenantId string, opts ...*ddd_repository.SetOptions) error {
	var model T
	model = d.NewEntity()
	return d.getDB(ctx).Where(WhereTenantId, tenantId).Delete(&model).Error
}

func (d *Dao[T]) DeleteByFilter(ctx context.Context, tenantId string, filter string, opts ...*ddd_repository.SetOptions) error {
	var model T
	model = d.NewEntity()
	where, err := getSqlWhere(tenantId, filter)
	if err != nil {
		return err
	}
	return d.getDB(ctx).Where(where).Delete(&model).Error
}

func (d *Dao[T]) FindById(ctx context.Context, tenantId string, id string, opts ...*ddd_repository.FindOptions) (T, bool, error) {
	var model T
	model = d.NewEntity()
	tx := d.getDB(ctx).Where(WhereTenantIdAndId, tenantId, id).Find(&model)
	if tx.Error == gorm.ErrRecordNotFound {
		return nil, false, nil
	} else if tx.Error != nil {
		return nil, false, tx.Error
	}
	return model, true, nil
}

func (d *Dao[T]) FindByIds(ctx context.Context, tenantId string, ids []string) ([]T, bool, error) {
	var vList []T
	vList = d.NewEntities()
	tx := d.getDB(ctx).Where(ids).Find(&vList)
	if IsErrRecordNotFound(tx.Error) {
		return vList, false, nil
	} else if tx.Error != nil {
		return vList, false, tx.Error
	}
	return vList, len(vList) > 0, tx.Error
}

func (d *Dao[T]) FindAll(ctx context.Context, tenantId string, opts ...*ddd_repository.FindOptions) *ddd_repository.FindListResult[T] {
	var vList []T
	vList = d.NewEntities()
	tx := d.getDB(ctx).Where(WhereTenantId, tenantId).Find(&vList)
	if IsErrRecordNotFound(tx.Error) {
		return ddd_repository.NewFindListResult(vList, false, nil)
	} else if tx.Error != nil {
		return ddd_repository.NewFindListResult(vList, false, tx.Error)
	}
	return ddd_repository.NewFindListResult(vList, len(vList) > 0, nil)
}

func (d *Dao[T]) FindListByMap(ctx context.Context, tenantId string, filterMap map[string]interface{}, opts ...*ddd_repository.FindOptions) *ddd_repository.FindListResult[T] {
	var vList []T
	vList = d.NewEntities()
	tx := d.getDB(ctx).Where(filterMap).Find(&vList)
	if IsErrRecordNotFound(tx.Error) {
		return ddd_repository.NewFindListResult(vList, false, nil)
	} else if tx.Error != nil {
		return ddd_repository.NewFindListResult(vList, false, tx.Error)
	}
	return ddd_repository.NewFindListResult(vList, len(vList) > 0, nil)
}

func (d *Dao[T]) FindPaging(ctx context.Context, query ddd_repository.FindPagingQuery, opts ...*ddd_repository.FindOptions) *ddd_repository.FindPagingResult[T] {
	return d.findPaging(ctx, query)
}

func (d *Dao[T]) findPaging(ctx context.Context, query ddd_repository.FindPagingQuery) *ddd_repository.FindPagingResult[T] {
	return d.DoFilter(query.GetTenantId(), query.GetFilter(), func(sqlWhere string) (*ddd_repository.FindPagingResult[T], bool, error) {
		var data []T = d.NewEntities()

		tx := d.getDB(ctx)

		if len(sqlWhere) > 0 {
			tx = tx.Where(sqlWhere)
		}

		if query.GetPageSize() > 0 {
			tx = tx.Limit(int(query.GetPageSize()))
		}

		if query.GetPageNum() > 0 {
			tx = tx.Offset(int(query.GetPageSize() * query.GetPageNum()))
		}

		if len(query.GetSort()) > 0 {
			tx = tx.Order(query.GetSort())
		}

		err := tx.Find(&data).Error
		if IsErrRecordNotFound(err) {
			return nil, false, nil
		}

		var totalRows int64 = -1
		if query.GetIsTotalRows() {
			tx := d.getDB(ctx)
			if len(sqlWhere) > 0 {
				tx = tx.Where(sqlWhere)
			}
			if err := tx.Count(&totalRows).Error; err != nil {
				return nil, false, err
			}
		}

		findData := ddd_repository.NewFindPagingResult[T](data, totalRows, query, err)
		return findData, true, err
	})

}

func (d *Dao[T]) findList(ctx context.Context, tenantId string, where string, limit *int64, opts ...*Options) ([]T, bool, error) {
	opt := NewOptions().SetDbId(tenantId).Merge(opts...)
	findOpts := newFindOptions(opt)
	findOpts.Limit = limit
	list := d.NewEntities()
	tx := d.getDB(ctx)
	var err error
	if limit != nil {
		var l int = 0
		l = int(*limit)
		tx = tx.Limit(l)
	}
	if opt.Sort != nil {
		order := *opt.Sort
		tx = tx.Order(order)
	}
	w := "tenant_id=?"
	if len(where) > 0 {
		w = fmt.Sprintf("(%v) and tenant_id=?", where)
	}
	err = tx.Where(w, tenantId).Find(&list).Error
	if IsErrRecordNotFound(err) {
		return list, false, nil
	} else if err != nil {
		return list, false, err
	}
	return list, true, nil
}

func (d *Dao[T]) DoFilter(tenantId, filter string, fun func(sqlWhere string) (*ddd_repository.FindPagingResult[T], bool, error)) *ddd_repository.FindPagingResult[T] {
	p := NewSqlProcess()
	if err := ParseProcess(filter, p); err != nil {
		return ddd_repository.NewFindPagingResultWithError[T](err)
	}
	sqlWhere, err := getSqlWhere(tenantId, filter)
	if err != nil {
		return ddd_repository.NewFindPagingResultWithError[T](err)
	}
	data, _, err := fun(sqlWhere)
	if err != nil {
		if IsErrRecordNotFound(err) {
			err = nil
		}
	}
	return data
}

func (d *Dao[T]) getSort(sort string) (map[string]interface{}, error) {
	if len(sort) == 0 {
		return nil, nil
	}
	//name:desc,id:asc
	res := map[string]interface{}{}
	list := strings.Split(sort, ",")
	for _, s := range list {
		sortItem := strings.Split(s, ":")
		name := sortItem[0]
		name = strings.Trim(name, " ")
		if name == "id" {
			name = "id"
		}
		order := "asc"
		if len(sortItem) > 1 {
			order = sortItem[1]
			order = strings.ToLower(order)
			order = strings.Trim(order, " ")
		}

		// 其中 1 为升序排列，而-1是用于降序排列.
		orderVal := 1
		var oerr error
		switch order {
		case "asc":
			orderVal = 1
		case "desc":
			orderVal = -1
		default:
			oerr = errors.New("order " + order + " is error")
		}
		if oerr != nil {
			return nil, oerr
		}
		res[name] = orderVal
	}
	return res, nil
}

func (d *Dao[T]) getDB(ctx context.Context) *gorm.DB {
	tx := db.GetTransaction(ctx)
	if tx != nil {
		return tx
	}
	return d.db
}

func (d *Dao[T]) NewEntity() T {
	return d.newFunc()
}

func (d *Dao[T]) NewEntities() []T {
	return d.newListFunc()
}
