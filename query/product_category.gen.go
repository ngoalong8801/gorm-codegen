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

	"gorm/model"
)

func newProductCategory(db *gorm.DB) productCategory {
	_productCategory := productCategory{}

	_productCategory.productCategoryDo.UseDB(db)
	_productCategory.productCategoryDo.UseModel(&model.ProductCategory{})

	tableName := _productCategory.productCategoryDo.TableName()
	_productCategory.ALL = field.NewField(tableName, "*")
	_productCategory.ProductID = field.NewInt64(tableName, "product_id")
	_productCategory.CategoryID = field.NewInt64(tableName, "category_id")

	_productCategory.fillFieldMap()

	return _productCategory
}

type productCategory struct {
	productCategoryDo productCategoryDo

	ALL        field.Field
	ProductID  field.Int64
	CategoryID field.Int64

	fieldMap map[string]field.Expr
}

func (p productCategory) Table(newTableName string) *productCategory {
	p.productCategoryDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p productCategory) As(alias string) *productCategory {
	p.productCategoryDo.DO = *(p.productCategoryDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *productCategory) updateTableName(table string) *productCategory {
	p.ALL = field.NewField(table, "*")
	p.ProductID = field.NewInt64(table, "product_id")
	p.CategoryID = field.NewInt64(table, "category_id")

	p.fillFieldMap()

	return p
}

func (p *productCategory) WithContext(ctx context.Context) *productCategoryDo {
	return p.productCategoryDo.WithContext(ctx)
}

func (p productCategory) TableName() string { return p.productCategoryDo.TableName() }

func (p productCategory) Alias() string { return p.productCategoryDo.Alias() }

func (p *productCategory) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *productCategory) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 2)
	p.fieldMap["product_id"] = p.ProductID
	p.fieldMap["category_id"] = p.CategoryID
}

func (p productCategory) clone(db *gorm.DB) productCategory {
	p.productCategoryDo.ReplaceDB(db)
	return p
}

type productCategoryDo struct{ gen.DO }

func (p productCategoryDo) Debug() *productCategoryDo {
	return p.withDO(p.DO.Debug())
}

func (p productCategoryDo) WithContext(ctx context.Context) *productCategoryDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p productCategoryDo) ReadDB() *productCategoryDo {
	return p.Clauses(dbresolver.Read)
}

func (p productCategoryDo) WriteDB() *productCategoryDo {
	return p.Clauses(dbresolver.Write)
}

func (p productCategoryDo) Clauses(conds ...clause.Expression) *productCategoryDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p productCategoryDo) Returning(value interface{}, columns ...string) *productCategoryDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p productCategoryDo) Not(conds ...gen.Condition) *productCategoryDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p productCategoryDo) Or(conds ...gen.Condition) *productCategoryDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p productCategoryDo) Select(conds ...field.Expr) *productCategoryDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p productCategoryDo) Where(conds ...gen.Condition) *productCategoryDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p productCategoryDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *productCategoryDo {
	return p.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (p productCategoryDo) Order(conds ...field.Expr) *productCategoryDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p productCategoryDo) Distinct(cols ...field.Expr) *productCategoryDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p productCategoryDo) Omit(cols ...field.Expr) *productCategoryDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p productCategoryDo) Join(table schema.Tabler, on ...field.Expr) *productCategoryDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p productCategoryDo) LeftJoin(table schema.Tabler, on ...field.Expr) *productCategoryDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p productCategoryDo) RightJoin(table schema.Tabler, on ...field.Expr) *productCategoryDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p productCategoryDo) Group(cols ...field.Expr) *productCategoryDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p productCategoryDo) Having(conds ...gen.Condition) *productCategoryDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p productCategoryDo) Limit(limit int) *productCategoryDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p productCategoryDo) Offset(offset int) *productCategoryDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p productCategoryDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *productCategoryDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p productCategoryDo) Unscoped() *productCategoryDo {
	return p.withDO(p.DO.Unscoped())
}

func (p productCategoryDo) Create(values ...*model.ProductCategory) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p productCategoryDo) CreateInBatches(values []*model.ProductCategory, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p productCategoryDo) Save(values ...*model.ProductCategory) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p productCategoryDo) First() (*model.ProductCategory, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.ProductCategory), nil
	}
}

func (p productCategoryDo) Take() (*model.ProductCategory, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.ProductCategory), nil
	}
}

func (p productCategoryDo) Last() (*model.ProductCategory, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.ProductCategory), nil
	}
}

func (p productCategoryDo) Find() ([]*model.ProductCategory, error) {
	result, err := p.DO.Find()
	return result.([]*model.ProductCategory), err
}

func (p productCategoryDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.ProductCategory, err error) {
	buf := make([]*model.ProductCategory, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p productCategoryDo) FindInBatches(result *[]*model.ProductCategory, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p productCategoryDo) Attrs(attrs ...field.AssignExpr) *productCategoryDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p productCategoryDo) Assign(attrs ...field.AssignExpr) *productCategoryDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p productCategoryDo) Joins(fields ...field.RelationField) *productCategoryDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p productCategoryDo) Preload(fields ...field.RelationField) *productCategoryDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p productCategoryDo) FirstOrInit() (*model.ProductCategory, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.ProductCategory), nil
	}
}

func (p productCategoryDo) FirstOrCreate() (*model.ProductCategory, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.ProductCategory), nil
	}
}

func (p productCategoryDo) FindByPage(offset int, limit int) (result []*model.ProductCategory, count int64, err error) {
	result, err = p.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = p.Offset(-1).Limit(-1).Count()
	return
}

func (p productCategoryDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p productCategoryDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p productCategoryDo) Delete(models ...*model.ProductCategory) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *productCategoryDo) withDO(do gen.Dao) *productCategoryDo {
	p.DO = *do.(*gen.DO)
	return p
}