// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"4clabs-server/app/4clabs-server/internal/adapter/data/model"
)

func newRegisterNft(db *gorm.DB) registerNft {
	_registerNft := registerNft{}

	_registerNft.registerNftDo.UseDB(db)
	_registerNft.registerNftDo.UseModel(&model.RegisterNft{})

	tableName := _registerNft.registerNftDo.TableName()
	_registerNft.ALL = field.NewField(tableName, "*")
	_registerNft.ID = field.NewInt32(tableName, "id")
	_registerNft.TokenID = field.NewString(tableName, "token_id")
	_registerNft.Name = field.NewString(tableName, "name")
	_registerNft.CollectionName = field.NewString(tableName, "collection_name")
	_registerNft.ContractAddress = field.NewString(tableName, "contract_address")
	_registerNft.UserAddress = field.NewString(tableName, "user_address")
	_registerNft.Price = field.NewFloat64(tableName, "price")
	_registerNft.Image = field.NewString(tableName, "image")
	_registerNft.CreatedAt = field.NewTime(tableName, "created_at")
	_registerNft.UpdatedAt = field.NewTime(tableName, "updated_at")

	_registerNft.fillFieldMap()

	return _registerNft
}

type registerNft struct {
	registerNftDo registerNftDo

	ALL             field.Field
	ID              field.Int32
	TokenID         field.String
	Name            field.String
	CollectionName  field.String
	ContractAddress field.String
	UserAddress     field.String
	Price           field.Float64
	Image           field.String
	CreatedAt       field.Time
	UpdatedAt       field.Time

	fieldMap map[string]field.Expr
}

func (r registerNft) Table(newTableName string) *registerNft {
	r.registerNftDo.UseTable(newTableName)
	return r.updateTableName(newTableName)
}

func (r registerNft) As(alias string) *registerNft {
	r.registerNftDo.DO = *(r.registerNftDo.As(alias).(*gen.DO))
	return r.updateTableName(alias)
}

func (r *registerNft) updateTableName(table string) *registerNft {
	r.ALL = field.NewField(table, "*")
	r.ID = field.NewInt32(table, "id")
	r.TokenID = field.NewString(table, "token_id")
	r.Name = field.NewString(table, "name")
	r.CollectionName = field.NewString(table, "collection_name")
	r.ContractAddress = field.NewString(table, "contract_address")
	r.UserAddress = field.NewString(table, "user_address")
	r.Price = field.NewFloat64(table, "price")
	r.Image = field.NewString(table, "image")
	r.CreatedAt = field.NewTime(table, "created_at")
	r.UpdatedAt = field.NewTime(table, "updated_at")

	r.fillFieldMap()

	return r
}

func (r *registerNft) WithContext(ctx context.Context) *registerNftDo {
	return r.registerNftDo.WithContext(ctx)
}

func (r registerNft) TableName() string { return r.registerNftDo.TableName() }

func (r registerNft) Alias() string { return r.registerNftDo.Alias() }

func (r *registerNft) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := r.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (r *registerNft) fillFieldMap() {
	r.fieldMap = make(map[string]field.Expr, 10)
	r.fieldMap["id"] = r.ID
	r.fieldMap["token_id"] = r.TokenID
	r.fieldMap["name"] = r.Name
	r.fieldMap["collection_name"] = r.CollectionName
	r.fieldMap["contract_address"] = r.ContractAddress
	r.fieldMap["user_address"] = r.UserAddress
	r.fieldMap["price"] = r.Price
	r.fieldMap["image"] = r.Image
	r.fieldMap["created_at"] = r.CreatedAt
	r.fieldMap["updated_at"] = r.UpdatedAt
}

func (r registerNft) clone(db *gorm.DB) registerNft {
	r.registerNftDo.ReplaceDB(db)
	return r
}

type registerNftDo struct{ gen.DO }

func (r registerNftDo) Debug() *registerNftDo {
	return r.withDO(r.DO.Debug())
}

func (r registerNftDo) WithContext(ctx context.Context) *registerNftDo {
	return r.withDO(r.DO.WithContext(ctx))
}

func (r registerNftDo) ReadDB(ctx context.Context) *registerNftDo {
	return r.WithContext(ctx).Clauses(dbresolver.Read)
}

func (r registerNftDo) WriteDB(ctx context.Context) *registerNftDo {
	return r.WithContext(ctx).Clauses(dbresolver.Write)
}

func (r registerNftDo) Clauses(conds ...clause.Expression) *registerNftDo {
	return r.withDO(r.DO.Clauses(conds...))
}

func (r registerNftDo) Returning(value interface{}, columns ...string) *registerNftDo {
	return r.withDO(r.DO.Returning(value, columns...))
}

func (r registerNftDo) Not(conds ...gen.Condition) *registerNftDo {
	return r.withDO(r.DO.Not(conds...))
}

func (r registerNftDo) Or(conds ...gen.Condition) *registerNftDo {
	return r.withDO(r.DO.Or(conds...))
}

func (r registerNftDo) Select(conds ...field.Expr) *registerNftDo {
	return r.withDO(r.DO.Select(conds...))
}

func (r registerNftDo) Where(conds ...gen.Condition) *registerNftDo {
	return r.withDO(r.DO.Where(conds...))
}

func (r registerNftDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *registerNftDo {
	return r.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (r registerNftDo) Order(conds ...field.Expr) *registerNftDo {
	return r.withDO(r.DO.Order(conds...))
}

func (r registerNftDo) Distinct(cols ...field.Expr) *registerNftDo {
	return r.withDO(r.DO.Distinct(cols...))
}

func (r registerNftDo) Omit(cols ...field.Expr) *registerNftDo {
	return r.withDO(r.DO.Omit(cols...))
}

func (r registerNftDo) Join(table schema.Tabler, on ...field.Expr) *registerNftDo {
	return r.withDO(r.DO.Join(table, on...))
}

func (r registerNftDo) LeftJoin(table schema.Tabler, on ...field.Expr) *registerNftDo {
	return r.withDO(r.DO.LeftJoin(table, on...))
}

func (r registerNftDo) RightJoin(table schema.Tabler, on ...field.Expr) *registerNftDo {
	return r.withDO(r.DO.RightJoin(table, on...))
}

func (r registerNftDo) Group(cols ...field.Expr) *registerNftDo {
	return r.withDO(r.DO.Group(cols...))
}

func (r registerNftDo) Having(conds ...gen.Condition) *registerNftDo {
	return r.withDO(r.DO.Having(conds...))
}

func (r registerNftDo) Limit(limit int) *registerNftDo {
	return r.withDO(r.DO.Limit(limit))
}

func (r registerNftDo) Offset(offset int) *registerNftDo {
	return r.withDO(r.DO.Offset(offset))
}

func (r registerNftDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *registerNftDo {
	return r.withDO(r.DO.Scopes(funcs...))
}

func (r registerNftDo) Unscoped() *registerNftDo {
	return r.withDO(r.DO.Unscoped())
}

func (r registerNftDo) Create(values ...*model.RegisterNft) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Create(values)
}

func (r registerNftDo) CreateInBatches(values []*model.RegisterNft, batchSize int) error {
	return r.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (r registerNftDo) Save(values ...*model.RegisterNft) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Save(values)
}

func (r registerNftDo) First() (*model.RegisterNft, error) {
	if result, err := r.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.RegisterNft), nil
	}
}

func (r registerNftDo) Take() (*model.RegisterNft, error) {
	if result, err := r.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.RegisterNft), nil
	}
}

func (r registerNftDo) Last() (*model.RegisterNft, error) {
	if result, err := r.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.RegisterNft), nil
	}
}

func (r registerNftDo) Find() ([]*model.RegisterNft, error) {
	result, err := r.DO.Find()
	return result.([]*model.RegisterNft), err
}

func (r registerNftDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.RegisterNft, err error) {
	buf := make([]*model.RegisterNft, 0, batchSize)
	err = r.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (r registerNftDo) FindInBatches(result *[]*model.RegisterNft, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return r.DO.FindInBatches(result, batchSize, fc)
}

func (r registerNftDo) Attrs(attrs ...field.AssignExpr) *registerNftDo {
	return r.withDO(r.DO.Attrs(attrs...))
}

func (r registerNftDo) Assign(attrs ...field.AssignExpr) *registerNftDo {
	return r.withDO(r.DO.Assign(attrs...))
}

func (r registerNftDo) Joins(fields ...field.RelationField) *registerNftDo {
	for _, _f := range fields {
		r = *r.withDO(r.DO.Joins(_f))
	}
	return &r
}

func (r registerNftDo) Preload(fields ...field.RelationField) *registerNftDo {
	for _, _f := range fields {
		r = *r.withDO(r.DO.Preload(_f))
	}
	return &r
}

func (r registerNftDo) FirstOrInit() (*model.RegisterNft, error) {
	if result, err := r.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.RegisterNft), nil
	}
}

func (r registerNftDo) FirstOrCreate() (*model.RegisterNft, error) {
	if result, err := r.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.RegisterNft), nil
	}
}

func (r registerNftDo) FindByPage(offset int, limit int) (result []*model.RegisterNft, count int64, err error) {
	result, err = r.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = r.Offset(-1).Limit(-1).Count()
	return
}

func (r registerNftDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = r.Count()
	if err != nil {
		return
	}

	err = r.Offset(offset).Limit(limit).Scan(result)
	return
}

func (r *registerNftDo) withDO(do gen.Dao) *registerNftDo {
	r.DO = *do.(*gen.DO)
	return r
}
