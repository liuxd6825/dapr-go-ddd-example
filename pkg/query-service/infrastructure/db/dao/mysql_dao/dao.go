package mysql_dao

import (
	"context"
	"errors"
	"fmt"
	"github.com/liuxd6825/components-contrib/liuxd/eventstorage/impl/gorm_impl/db"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd/ddd_repository"
	"github.com/liuxd6825/dapr-go-ddd-sdk/rsql"
	"github.com/liuxd6825/dapr-go-ddd-sdk/utils/stringutils"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/gorm"
	"strings"
	"time"
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

type Process interface {
	OnAndItem()
	OnAndStart()
	OnAndEnd()
	OnOrItem()
	OnOrStart()
	OnOrEnd()
	OnEquals(name string, value interface{}, rValue rsql.Value)
	OnNotEquals(name string, value interface{}, rValue rsql.Value)
	OnLike(name string, value interface{}, rValue rsql.Value)
	OnNotLike(name string, value interface{}, rValue rsql.Value)
	OnGreaterThan(name string, value interface{}, rValue rsql.Value)
	OnGreaterThanOrEquals(name string, value interface{}, rValue rsql.Value)
	OnLessThan(name string, value interface{}, rValue rsql.Value)
	OnLessThanOrEquals(name string, value interface{}, rValue rsql.Value)
	OnIn(name string, value interface{}, rValue rsql.Value)
	OnNotIn(name string, value interface{}, rValue rsql.Value)
	GetSqlWhere(tenantId string) interface{}
	GetFilter(tenantId string) interface{}
}

type process struct {
	str string
}

func NewSqlProcess() Process {
	return &process{}
}

func (p *process) OnNotEquals(name string, value interface{}, rValue rsql.Value) {
	p.str = fmt.Sprintf("%s %s != (%v)", p.str, name, value)
}

func (p *process) OnLike(name string, value interface{}, rValue rsql.Value) {
	p.str = fmt.Sprintf("%s %s like (%v)", p.str, name, value)
}

func (p *process) OnNotLike(name string, value interface{}, rValue rsql.Value) {
	p.str = fmt.Sprintf("%s %s not like %v", p.str, name, value)
}

func (p *process) OnGreaterThan(name string, value interface{}, rValue rsql.Value) {
	p.str = fmt.Sprintf("%s %s>%v", p.str, name, value)
}

func (p *process) OnGreaterThanOrEquals(name string, value interface{}, rValue rsql.Value) {
	p.str = fmt.Sprintf("%s %s>=%v", p.str, name, value)
}

func (p *process) OnLessThan(name string, value interface{}, rValue rsql.Value) {
	p.str = fmt.Sprintf("%s %s<%v", p.str, name, value)
}

func (p *process) OnLessThanOrEquals(name string, value interface{}, rValue rsql.Value) {
	p.str = fmt.Sprintf("%s %s <= %v", p.str, name, value)
}

func (p *process) OnIn(name string, value interface{}, rValue rsql.Value) {
	p.str = fmt.Sprintf("%s %s in %v", p.str, name, value)
}

func (p *process) OnNotIn(name string, value interface{}, rValue rsql.Value) {
	p.str = fmt.Sprintf("%s %s not in %v", p.str, name, value)
}

func (p *process) OnEquals(name string, value interface{}, rValue rsql.Value) {
	p.str = fmt.Sprintf("%s %s=%v", p.str, name, value)
}

func (p *process) NotEquals(name string, value interface{}, rValue rsql.Value) {
	p.str = fmt.Sprintf("%s %s=%v", p.str, name, value)
}

func (p *process) OnAndItem() {
	p.str = fmt.Sprintf("%s and ", p.str)
}
func (p *process) OnAndStart() {
	p.str = fmt.Sprintf("%s(", p.str)
}
func (p *process) OnAndEnd() {
	p.str = fmt.Sprintf("%s)", p.str)
}
func (p *process) OnOrItem() {
	p.str = fmt.Sprintf("%s or ", p.str)
}
func (p *process) OnOrStart() {
	p.str = fmt.Sprintf("%s(", p.str)
}
func (p *process) OnOrEnd() {
	p.str = fmt.Sprintf("%s)", p.str)
}

func (p *process) GetFilter(tenantId string) interface{} {
	return p.str
}

func (p *process) GetSqlWhere(tenantId string) interface{} {
	return p.str
}
func (p *process) Print() {
	fmt.Print(p.str)
}

func ParseProcess(input string, process Process) error {
	if len(input) == 0 {
		return nil
	}
	expr, err := rsql.Parse(input)
	if err != nil {
		return errors.New(fmt.Sprintf("rsql %s expression error, %s", input, err.Error()))
	}
	err = parseProcess(expr, process)
	if err != nil {
		return errors.New(fmt.Sprintf("rsql %s parseProcess error, %s", input, err.Error()))
	}
	return nil
}

func parseProcess(expr rsql.Expression, process rsql.Process) error {
	switch expr.(type) {
	case rsql.AndExpression:
		ex, _ := expr.(rsql.AndExpression)
		process.OnAndStart()
		for i, e := range ex.Items {
			_ = parseProcess(e, process)
			if i < len(ex.Items)-1 {
				process.OnAndItem()
			}
		}
		process.OnAndEnd()
		break
	case rsql.OrExpression:
		ex, _ := expr.(rsql.OrExpression)
		process.OnOrStart()
		for i, e := range ex.Items {
			_ = parseProcess(e, process)
			if i < len(ex.Items)-1 {
				process.OnOrItem()
			}
		}
		process.OnOrEnd()
		break
	case rsql.NotEqualsComparison:
		ex, _ := expr.(rsql.NotEqualsComparison)
		name := ex.Comparison.Identifier.Val
		name = stringutils.AsFieldName(name)
		value := getValue(ex.Comparison.Val)
		process.OnNotEquals(name, value, ex.Comparison.Val)
		break
	case rsql.EqualsComparison:
		ex, _ := expr.(rsql.EqualsComparison)
		name := ex.Comparison.Identifier.Val
		name = stringutils.AsFieldName(name)
		value := getValue(ex.Comparison.Val)
		process.OnEquals(name, value, ex.Comparison.Val)
		break
	case rsql.LikeComparison:
		ex, _ := expr.(rsql.LikeComparison)
		name := stringutils.AsFieldName(ex.Comparison.Identifier.Val)
		value := getValue(ex.Comparison.Val)
		process.OnLike(name, value, ex.Comparison.Val)
		break
	case rsql.NotLikeComparison:
		ex, _ := expr.(rsql.NotLikeComparison)
		name := ex.Comparison.Identifier.Val
		name = stringutils.AsFieldName(name)
		value := getValue(ex.Comparison.Val)
		process.OnNotLike(name, value, ex.Comparison.Val)
		break
	case rsql.GreaterThanComparison:
		ex, _ := expr.(rsql.GreaterThanComparison)
		name := ex.Comparison.Identifier.Val
		name = stringutils.AsFieldName(name)
		value := getValue(ex.Comparison.Val)
		process.OnGreaterThan(name, value, ex.Comparison.Val)
		break
	case rsql.GreaterThanOrEqualsComparison:
		ex, _ := expr.(rsql.GreaterThanOrEqualsComparison)
		name := ex.Comparison.Identifier.Val
		name = stringutils.AsFieldName(name)
		value := getValue(ex.Comparison.Val)
		process.OnGreaterThanOrEquals(name, value, ex.Comparison.Val)
		break
	case rsql.LessThanComparison:
		ex, _ := expr.(rsql.LessThanComparison)
		name := ex.Comparison.Identifier.Val
		value := getValue(ex.Comparison.Val)
		process.OnLessThan(name, value, ex.Comparison.Val)
		break
	case rsql.LessThanOrEqualsComparison:
		ex, _ := expr.(rsql.LessThanOrEqualsComparison)
		name := ex.Comparison.Identifier.Val
		name = stringutils.AsFieldName(name)
		value := getValue(ex.Comparison.Val)
		process.OnLessThanOrEquals(name, value, ex.Comparison.Val)
		break
	case rsql.InComparison:
		ex, _ := expr.(rsql.InComparison)
		name := ex.Comparison.Identifier.Val
		name = stringutils.AsFieldName(name)
		value := getValue(ex.Comparison.Val)
		process.OnIn(name, value, ex.Comparison.Val)
		break
	case rsql.NotInComparison:
		ex, _ := expr.(rsql.NotInComparison)
		name := ex.Comparison.Identifier.Val
		name = stringutils.AsFieldName(name)
		value := getValue(ex.Comparison.Val)
		process.OnNotIn(name, value, ex.Comparison.Val)
		break
	}
	return nil
}

func getValue(val rsql.Value) interface{} {
	var value interface{}
	switch val.(type) {
	case rsql.IntegerValue:
		value = val.(rsql.IntegerValue).Value
		break
	case rsql.BooleanValue:
		value = val.(rsql.BooleanValue).Value
		break
	case rsql.StringValue:
		value = fmt.Sprintf(`"%v"`, val.(rsql.StringValue).Value)
		break
	case rsql.DateTimeValue:
		value = fmt.Sprintf(`"%v"`, val.(rsql.DateTimeValue).Value)
		break
	case rsql.DoubleValue:
		value = val.(rsql.DoubleValue).Value
		break
	}
	return value
}

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

func GetMySqlDB() *gorm.DB {
	return nil
}

func NewSession(isWrite bool) ddd_repository.Session {
	return nil
}

func IsErrRecordNotFound(err error) bool {
	if err == gorm.ErrRecordNotFound {
		return true
	}
	return false
}

func newFindOptions(opt *Options) *options.FindOptions {
	findOptions := &options.FindOptions{}
	findOptions.MaxTime = opt.MaxTime
	findOptions.Sort = opt.Sort
	return findOptions
}

func getSqlWhere(tenantId, filter string) (string, error) {
	p := NewSqlProcess()
	if err := ParseProcess(filter, p); err != nil {
		return "", err
	}
	return p.GetSqlWhere(tenantId).(string), nil
}
