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

func newComicsNft(db *gorm.DB) comicsNft {
	_comicsNft := comicsNft{}

	_comicsNft.comicsNftDo.UseDB(db)
	_comicsNft.comicsNftDo.UseModel(&model.ComicsNft{})

	tableName := _comicsNft.comicsNftDo.TableName()
	_comicsNft.ALL = field.NewField(tableName, "*")
	_comicsNft.ID = field.NewInt32(tableName, "id")
	_comicsNft.ComicsID = field.NewString(tableName, "comics_id")
	_comicsNft.Owner = field.NewString(tableName, "owner")
	_comicsNft.Author = field.NewString(tableName, "author")
	_comicsNft.Comic = comicsNftBelongsToComic{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Comic", "model.Comic"),
	}

	_comicsNft.fillFieldMap()

	return _comicsNft
}

type comicsNft struct {
	comicsNftDo comicsNftDo

	ALL      field.Field
	ID       field.Int32
	ComicsID field.String
	Owner    field.String
	Author   field.String
	Comic    comicsNftBelongsToComic

	fieldMap map[string]field.Expr
}

func (c comicsNft) Table(newTableName string) *comicsNft {
	c.comicsNftDo.UseTable(newTableName)
	return c.updateTableName(newTableName)
}

func (c comicsNft) As(alias string) *comicsNft {
	c.comicsNftDo.DO = *(c.comicsNftDo.As(alias).(*gen.DO))
	return c.updateTableName(alias)
}

func (c *comicsNft) updateTableName(table string) *comicsNft {
	c.ALL = field.NewField(table, "*")
	c.ID = field.NewInt32(table, "id")
	c.ComicsID = field.NewString(table, "comics_id")
	c.Owner = field.NewString(table, "owner")
	c.Author = field.NewString(table, "author")

	c.fillFieldMap()

	return c
}

func (c *comicsNft) WithContext(ctx context.Context) *comicsNftDo {
	return c.comicsNftDo.WithContext(ctx)
}

func (c comicsNft) TableName() string { return c.comicsNftDo.TableName() }

func (c comicsNft) Alias() string { return c.comicsNftDo.Alias() }

func (c *comicsNft) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := c.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (c *comicsNft) fillFieldMap() {
	c.fieldMap = make(map[string]field.Expr, 5)
	c.fieldMap["id"] = c.ID
	c.fieldMap["comics_id"] = c.ComicsID
	c.fieldMap["owner"] = c.Owner
	c.fieldMap["author"] = c.Author

}

func (c comicsNft) clone(db *gorm.DB) comicsNft {
	c.comicsNftDo.ReplaceDB(db)
	return c
}

type comicsNftBelongsToComic struct {
	db *gorm.DB

	field.RelationField
}

func (a comicsNftBelongsToComic) Where(conds ...field.Expr) *comicsNftBelongsToComic {
	if len(conds) == 0 {
		return &a
	}

	exprs := make([]clause.Expression, 0, len(conds))
	for _, cond := range conds {
		exprs = append(exprs, cond.BeCond().(clause.Expression))
	}
	a.db = a.db.Clauses(clause.Where{Exprs: exprs})
	return &a
}

func (a comicsNftBelongsToComic) WithContext(ctx context.Context) *comicsNftBelongsToComic {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a comicsNftBelongsToComic) Model(m *model.ComicsNft) *comicsNftBelongsToComicTx {
	return &comicsNftBelongsToComicTx{a.db.Model(m).Association(a.Name())}
}

type comicsNftBelongsToComicTx struct{ tx *gorm.Association }

func (a comicsNftBelongsToComicTx) Find() (result *model.Comic, err error) {
	return result, a.tx.Find(&result)
}

func (a comicsNftBelongsToComicTx) Append(values ...*model.Comic) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a comicsNftBelongsToComicTx) Replace(values ...*model.Comic) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a comicsNftBelongsToComicTx) Delete(values ...*model.Comic) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a comicsNftBelongsToComicTx) Clear() error {
	return a.tx.Clear()
}

func (a comicsNftBelongsToComicTx) Count() int64 {
	return a.tx.Count()
}

type comicsNftDo struct{ gen.DO }

func (c comicsNftDo) Debug() *comicsNftDo {
	return c.withDO(c.DO.Debug())
}

func (c comicsNftDo) WithContext(ctx context.Context) *comicsNftDo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c comicsNftDo) ReadDB(ctx context.Context) *comicsNftDo {
	return c.WithContext(ctx).Clauses(dbresolver.Read)
}

func (c comicsNftDo) WriteDB(ctx context.Context) *comicsNftDo {
	return c.WithContext(ctx).Clauses(dbresolver.Write)
}

func (c comicsNftDo) Clauses(conds ...clause.Expression) *comicsNftDo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c comicsNftDo) Returning(value interface{}, columns ...string) *comicsNftDo {
	return c.withDO(c.DO.Returning(value, columns...))
}

func (c comicsNftDo) Not(conds ...gen.Condition) *comicsNftDo {
	return c.withDO(c.DO.Not(conds...))
}

func (c comicsNftDo) Or(conds ...gen.Condition) *comicsNftDo {
	return c.withDO(c.DO.Or(conds...))
}

func (c comicsNftDo) Select(conds ...field.Expr) *comicsNftDo {
	return c.withDO(c.DO.Select(conds...))
}

func (c comicsNftDo) Where(conds ...gen.Condition) *comicsNftDo {
	return c.withDO(c.DO.Where(conds...))
}

func (c comicsNftDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *comicsNftDo {
	return c.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (c comicsNftDo) Order(conds ...field.Expr) *comicsNftDo {
	return c.withDO(c.DO.Order(conds...))
}

func (c comicsNftDo) Distinct(cols ...field.Expr) *comicsNftDo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c comicsNftDo) Omit(cols ...field.Expr) *comicsNftDo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c comicsNftDo) Join(table schema.Tabler, on ...field.Expr) *comicsNftDo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c comicsNftDo) LeftJoin(table schema.Tabler, on ...field.Expr) *comicsNftDo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c comicsNftDo) RightJoin(table schema.Tabler, on ...field.Expr) *comicsNftDo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c comicsNftDo) Group(cols ...field.Expr) *comicsNftDo {
	return c.withDO(c.DO.Group(cols...))
}

func (c comicsNftDo) Having(conds ...gen.Condition) *comicsNftDo {
	return c.withDO(c.DO.Having(conds...))
}

func (c comicsNftDo) Limit(limit int) *comicsNftDo {
	return c.withDO(c.DO.Limit(limit))
}

func (c comicsNftDo) Offset(offset int) *comicsNftDo {
	return c.withDO(c.DO.Offset(offset))
}

func (c comicsNftDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *comicsNftDo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c comicsNftDo) Unscoped() *comicsNftDo {
	return c.withDO(c.DO.Unscoped())
}

func (c comicsNftDo) Create(values ...*model.ComicsNft) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c comicsNftDo) CreateInBatches(values []*model.ComicsNft, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c comicsNftDo) Save(values ...*model.ComicsNft) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c comicsNftDo) First() (*model.ComicsNft, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.ComicsNft), nil
	}
}

func (c comicsNftDo) Take() (*model.ComicsNft, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.ComicsNft), nil
	}
}

func (c comicsNftDo) Last() (*model.ComicsNft, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.ComicsNft), nil
	}
}

func (c comicsNftDo) Find() ([]*model.ComicsNft, error) {
	result, err := c.DO.Find()
	return result.([]*model.ComicsNft), err
}

func (c comicsNftDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.ComicsNft, err error) {
	buf := make([]*model.ComicsNft, 0, batchSize)
	err = c.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (c comicsNftDo) FindInBatches(result *[]*model.ComicsNft, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(result, batchSize, fc)
}

func (c comicsNftDo) Attrs(attrs ...field.AssignExpr) *comicsNftDo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c comicsNftDo) Assign(attrs ...field.AssignExpr) *comicsNftDo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c comicsNftDo) Joins(fields ...field.RelationField) *comicsNftDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Joins(_f))
	}
	return &c
}

func (c comicsNftDo) Preload(fields ...field.RelationField) *comicsNftDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Preload(_f))
	}
	return &c
}

func (c comicsNftDo) FirstOrInit() (*model.ComicsNft, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.ComicsNft), nil
	}
}

func (c comicsNftDo) FirstOrCreate() (*model.ComicsNft, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.ComicsNft), nil
	}
}

func (c comicsNftDo) FindByPage(offset int, limit int) (result []*model.ComicsNft, count int64, err error) {
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

func (c comicsNftDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	err = c.Offset(offset).Limit(limit).Scan(result)
	return
}

func (c *comicsNftDo) withDO(do gen.Dao) *comicsNftDo {
	c.DO = *do.(*gen.DO)
	return c
}
