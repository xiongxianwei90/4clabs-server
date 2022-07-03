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

func newComic(db *gorm.DB) comic {
	_comic := comic{}

	_comic.comicDo.UseDB(db)
	_comic.comicDo.UseModel(&model.Comic{})

	tableName := _comic.comicDo.TableName()
	_comic.ALL = field.NewField(tableName, "*")
	_comic.ID = field.NewInt32(tableName, "id")
	_comic.TokenID = field.NewString(tableName, "token_id")
	_comic.Name = field.NewString(tableName, "name")
	_comic.ContractAddress = field.NewString(tableName, "contract_address")
	_comic.UserAddress = field.NewString(tableName, "user_address")
	_comic.MintLimit = field.NewInt32(tableName, "mint_limit")
	_comic.MintPrice = field.NewFloat64(tableName, "mint_price")
	_comic.Medata = field.NewString(tableName, "medata")
	_comic.CreatedAt = field.NewTime(tableName, "created_at")
	_comic.UpdatedAt = field.NewTime(tableName, "updated_at")

	_comic.fillFieldMap()

	return _comic
}

type comic struct {
	comicDo comicDo

	ALL             field.Field
	ID              field.Int32
	TokenID         field.String
	Name            field.String
	ContractAddress field.String
	UserAddress     field.String
	MintLimit       field.Int32
	MintPrice       field.Float64
	Medata          field.String
	CreatedAt       field.Time
	UpdatedAt       field.Time

	fieldMap map[string]field.Expr
}

func (c comic) Table(newTableName string) *comic {
	c.comicDo.UseTable(newTableName)
	return c.updateTableName(newTableName)
}

func (c comic) As(alias string) *comic {
	c.comicDo.DO = *(c.comicDo.As(alias).(*gen.DO))
	return c.updateTableName(alias)
}

func (c *comic) updateTableName(table string) *comic {
	c.ALL = field.NewField(table, "*")
	c.ID = field.NewInt32(table, "id")
	c.TokenID = field.NewString(table, "token_id")
	c.Name = field.NewString(table, "name")
	c.ContractAddress = field.NewString(table, "contract_address")
	c.UserAddress = field.NewString(table, "user_address")
	c.MintLimit = field.NewInt32(table, "mint_limit")
	c.MintPrice = field.NewFloat64(table, "mint_price")
	c.Medata = field.NewString(table, "medata")
	c.CreatedAt = field.NewTime(table, "created_at")
	c.UpdatedAt = field.NewTime(table, "updated_at")

	c.fillFieldMap()

	return c
}

func (c *comic) WithContext(ctx context.Context) *comicDo { return c.comicDo.WithContext(ctx) }

func (c comic) TableName() string { return c.comicDo.TableName() }

func (c comic) Alias() string { return c.comicDo.Alias() }

func (c *comic) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := c.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (c *comic) fillFieldMap() {
	c.fieldMap = make(map[string]field.Expr, 10)
	c.fieldMap["id"] = c.ID
	c.fieldMap["token_id"] = c.TokenID
	c.fieldMap["name"] = c.Name
	c.fieldMap["contract_address"] = c.ContractAddress
	c.fieldMap["user_address"] = c.UserAddress
	c.fieldMap["mint_limit"] = c.MintLimit
	c.fieldMap["mint_price"] = c.MintPrice
	c.fieldMap["medata"] = c.Medata
	c.fieldMap["created_at"] = c.CreatedAt
	c.fieldMap["updated_at"] = c.UpdatedAt
}

func (c comic) clone(db *gorm.DB) comic {
	c.comicDo.ReplaceDB(db)
	return c
}

type comicDo struct{ gen.DO }

func (c comicDo) Debug() *comicDo {
	return c.withDO(c.DO.Debug())
}

func (c comicDo) WithContext(ctx context.Context) *comicDo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c comicDo) ReadDB(ctx context.Context) *comicDo {
	return c.WithContext(ctx).Clauses(dbresolver.Read)
}

func (c comicDo) WriteDB(ctx context.Context) *comicDo {
	return c.WithContext(ctx).Clauses(dbresolver.Write)
}

func (c comicDo) Clauses(conds ...clause.Expression) *comicDo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c comicDo) Returning(value interface{}, columns ...string) *comicDo {
	return c.withDO(c.DO.Returning(value, columns...))
}

func (c comicDo) Not(conds ...gen.Condition) *comicDo {
	return c.withDO(c.DO.Not(conds...))
}

func (c comicDo) Or(conds ...gen.Condition) *comicDo {
	return c.withDO(c.DO.Or(conds...))
}

func (c comicDo) Select(conds ...field.Expr) *comicDo {
	return c.withDO(c.DO.Select(conds...))
}

func (c comicDo) Where(conds ...gen.Condition) *comicDo {
	return c.withDO(c.DO.Where(conds...))
}

func (c comicDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *comicDo {
	return c.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (c comicDo) Order(conds ...field.Expr) *comicDo {
	return c.withDO(c.DO.Order(conds...))
}

func (c comicDo) Distinct(cols ...field.Expr) *comicDo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c comicDo) Omit(cols ...field.Expr) *comicDo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c comicDo) Join(table schema.Tabler, on ...field.Expr) *comicDo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c comicDo) LeftJoin(table schema.Tabler, on ...field.Expr) *comicDo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c comicDo) RightJoin(table schema.Tabler, on ...field.Expr) *comicDo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c comicDo) Group(cols ...field.Expr) *comicDo {
	return c.withDO(c.DO.Group(cols...))
}

func (c comicDo) Having(conds ...gen.Condition) *comicDo {
	return c.withDO(c.DO.Having(conds...))
}

func (c comicDo) Limit(limit int) *comicDo {
	return c.withDO(c.DO.Limit(limit))
}

func (c comicDo) Offset(offset int) *comicDo {
	return c.withDO(c.DO.Offset(offset))
}

func (c comicDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *comicDo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c comicDo) Unscoped() *comicDo {
	return c.withDO(c.DO.Unscoped())
}

func (c comicDo) Create(values ...*model.Comic) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c comicDo) CreateInBatches(values []*model.Comic, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c comicDo) Save(values ...*model.Comic) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c comicDo) First() (*model.Comic, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Comic), nil
	}
}

func (c comicDo) Take() (*model.Comic, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Comic), nil
	}
}

func (c comicDo) Last() (*model.Comic, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Comic), nil
	}
}

func (c comicDo) Find() ([]*model.Comic, error) {
	result, err := c.DO.Find()
	return result.([]*model.Comic), err
}

func (c comicDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Comic, err error) {
	buf := make([]*model.Comic, 0, batchSize)
	err = c.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (c comicDo) FindInBatches(result *[]*model.Comic, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(result, batchSize, fc)
}

func (c comicDo) Attrs(attrs ...field.AssignExpr) *comicDo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c comicDo) Assign(attrs ...field.AssignExpr) *comicDo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c comicDo) Joins(fields ...field.RelationField) *comicDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Joins(_f))
	}
	return &c
}

func (c comicDo) Preload(fields ...field.RelationField) *comicDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Preload(_f))
	}
	return &c
}

func (c comicDo) FirstOrInit() (*model.Comic, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Comic), nil
	}
}

func (c comicDo) FirstOrCreate() (*model.Comic, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Comic), nil
	}
}

func (c comicDo) FindByPage(offset int, limit int) (result []*model.Comic, count int64, err error) {
	result, err = c.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = c.Offset(-1).Limit(-1).Count()
	return
}

func (c comicDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	err = c.Offset(offset).Limit(limit).Scan(result)
	return
}

func (c *comicDo) withDO(do gen.Dao) *comicDo {
	c.DO = *do.(*gen.DO)
	return c
}
