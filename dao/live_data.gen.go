// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dao

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"demo20220320/model"
)

func newLiveDatum(db *gorm.DB, opts ...gen.DOOption) liveDatum {
	_liveDatum := liveDatum{}

	_liveDatum.liveDatumDo.UseDB(db, opts...)
	_liveDatum.liveDatumDo.UseModel(&model.LiveDatum{})

	tableName := _liveDatum.liveDatumDo.TableName()
	_liveDatum.ALL = field.NewAsterisk(tableName)
	_liveDatum.ID = field.NewInt64(tableName, "id")
	_liveDatum.HostNickname = field.NewString(tableName, "host_nickname")
	_liveDatum.HostID = field.NewInt64(tableName, "host_id")
	_liveDatum.LiveStatus = field.NewString(tableName, "live_status")
	_liveDatum.HostStatus = field.NewString(tableName, "host_status")
	_liveDatum.Remarks = field.NewString(tableName, "remarks")
	_liveDatum.LiveDuration = field.NewInt64(tableName, "live_duration")
	_liveDatum.PaymentAmount = field.NewFloat64(tableName, "payment_amount")
	_liveDatum.ViewCount = field.NewInt64(tableName, "view_count")
	_liveDatum.ClickThroughRate = field.NewFloat64(tableName, "click_through_rate")
	_liveDatum.ConversionRate = field.NewFloat64(tableName, "conversion_rate")
	_liveDatum.TransactionCount = field.NewInt64(tableName, "transaction_count")
	_liveDatum.TransactionItems = field.NewInt64(tableName, "transaction_items")
	_liveDatum.YesterdaySales = field.NewFloat64(tableName, "yesterday_sales")
	_liveDatum.PauseCount = field.NewInt64(tableName, "pause_count")
	_liveDatum.PauseRecordID = field.NewInt64(tableName, "pause_record_id")
	_liveDatum.CreatedAt = field.NewInt64(tableName, "created_at")
	_liveDatum.UpdatedAt = field.NewInt64(tableName, "updated_at")

	_liveDatum.fillFieldMap()

	return _liveDatum
}

// liveDatum 直播数据表
type liveDatum struct {
	liveDatumDo

	ALL              field.Asterisk
	ID               field.Int64   // 主键ID
	HostNickname     field.String  // 主播昵称
	HostID           field.Int64   // 主播ID
	LiveStatus       field.String  // 直播状态
	HostStatus       field.String  // 主播状态
	Remarks          field.String  // 备注
	LiveDuration     field.Int64   // 直播时长（以分钟为单位）
	PaymentAmount    field.Float64 // 支付金额
	ViewCount        field.Int64   // 观看次数
	ClickThroughRate field.Float64 // 点击率
	ConversionRate   field.Float64 // 成交率
	TransactionCount field.Int64   // 成交人数
	TransactionItems field.Int64   // 成交件数
	YesterdaySales   field.Float64 // 昨日成交
	PauseCount       field.Int64   // 暂停次数
	PauseRecordID    field.Int64   // 暂停记录表ID，用于关联
	CreatedAt        field.Int64   // 创建时间（Unix 时间戳）
	UpdatedAt        field.Int64   // 修改时间（Unix 时间戳）

	fieldMap map[string]field.Expr
}

func (l liveDatum) Table(newTableName string) *liveDatum {
	l.liveDatumDo.UseTable(newTableName)
	return l.updateTableName(newTableName)
}

func (l liveDatum) As(alias string) *liveDatum {
	l.liveDatumDo.DO = *(l.liveDatumDo.As(alias).(*gen.DO))
	return l.updateTableName(alias)
}

func (l *liveDatum) updateTableName(table string) *liveDatum {
	l.ALL = field.NewAsterisk(table)
	l.ID = field.NewInt64(table, "id")
	l.HostNickname = field.NewString(table, "host_nickname")
	l.HostID = field.NewInt64(table, "host_id")
	l.LiveStatus = field.NewString(table, "live_status")
	l.HostStatus = field.NewString(table, "host_status")
	l.Remarks = field.NewString(table, "remarks")
	l.LiveDuration = field.NewInt64(table, "live_duration")
	l.PaymentAmount = field.NewFloat64(table, "payment_amount")
	l.ViewCount = field.NewInt64(table, "view_count")
	l.ClickThroughRate = field.NewFloat64(table, "click_through_rate")
	l.ConversionRate = field.NewFloat64(table, "conversion_rate")
	l.TransactionCount = field.NewInt64(table, "transaction_count")
	l.TransactionItems = field.NewInt64(table, "transaction_items")
	l.YesterdaySales = field.NewFloat64(table, "yesterday_sales")
	l.PauseCount = field.NewInt64(table, "pause_count")
	l.PauseRecordID = field.NewInt64(table, "pause_record_id")
	l.CreatedAt = field.NewInt64(table, "created_at")
	l.UpdatedAt = field.NewInt64(table, "updated_at")

	l.fillFieldMap()

	return l
}

func (l *liveDatum) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := l.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (l *liveDatum) fillFieldMap() {
	l.fieldMap = make(map[string]field.Expr, 18)
	l.fieldMap["id"] = l.ID
	l.fieldMap["host_nickname"] = l.HostNickname
	l.fieldMap["host_id"] = l.HostID
	l.fieldMap["live_status"] = l.LiveStatus
	l.fieldMap["host_status"] = l.HostStatus
	l.fieldMap["remarks"] = l.Remarks
	l.fieldMap["live_duration"] = l.LiveDuration
	l.fieldMap["payment_amount"] = l.PaymentAmount
	l.fieldMap["view_count"] = l.ViewCount
	l.fieldMap["click_through_rate"] = l.ClickThroughRate
	l.fieldMap["conversion_rate"] = l.ConversionRate
	l.fieldMap["transaction_count"] = l.TransactionCount
	l.fieldMap["transaction_items"] = l.TransactionItems
	l.fieldMap["yesterday_sales"] = l.YesterdaySales
	l.fieldMap["pause_count"] = l.PauseCount
	l.fieldMap["pause_record_id"] = l.PauseRecordID
	l.fieldMap["created_at"] = l.CreatedAt
	l.fieldMap["updated_at"] = l.UpdatedAt
}

func (l liveDatum) clone(db *gorm.DB) liveDatum {
	l.liveDatumDo.ReplaceConnPool(db.Statement.ConnPool)
	return l
}

func (l liveDatum) replaceDB(db *gorm.DB) liveDatum {
	l.liveDatumDo.ReplaceDB(db)
	return l
}

type liveDatumDo struct{ gen.DO }

type ILiveDatumDo interface {
	gen.SubQuery
	Debug() ILiveDatumDo
	WithContext(ctx context.Context) ILiveDatumDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ILiveDatumDo
	WriteDB() ILiveDatumDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ILiveDatumDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ILiveDatumDo
	Not(conds ...gen.Condition) ILiveDatumDo
	Or(conds ...gen.Condition) ILiveDatumDo
	Select(conds ...field.Expr) ILiveDatumDo
	Where(conds ...gen.Condition) ILiveDatumDo
	Order(conds ...field.Expr) ILiveDatumDo
	Distinct(cols ...field.Expr) ILiveDatumDo
	Omit(cols ...field.Expr) ILiveDatumDo
	Join(table schema.Tabler, on ...field.Expr) ILiveDatumDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ILiveDatumDo
	RightJoin(table schema.Tabler, on ...field.Expr) ILiveDatumDo
	Group(cols ...field.Expr) ILiveDatumDo
	Having(conds ...gen.Condition) ILiveDatumDo
	Limit(limit int) ILiveDatumDo
	Offset(offset int) ILiveDatumDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ILiveDatumDo
	Unscoped() ILiveDatumDo
	Create(values ...*model.LiveDatum) error
	CreateInBatches(values []*model.LiveDatum, batchSize int) error
	Save(values ...*model.LiveDatum) error
	First() (*model.LiveDatum, error)
	Take() (*model.LiveDatum, error)
	Last() (*model.LiveDatum, error)
	Find() ([]*model.LiveDatum, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.LiveDatum, err error)
	FindInBatches(result *[]*model.LiveDatum, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.LiveDatum) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ILiveDatumDo
	Assign(attrs ...field.AssignExpr) ILiveDatumDo
	Joins(fields ...field.RelationField) ILiveDatumDo
	Preload(fields ...field.RelationField) ILiveDatumDo
	FirstOrInit() (*model.LiveDatum, error)
	FirstOrCreate() (*model.LiveDatum, error)
	FindByPage(offset int, limit int) (result []*model.LiveDatum, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ILiveDatumDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (l liveDatumDo) Debug() ILiveDatumDo {
	return l.withDO(l.DO.Debug())
}

func (l liveDatumDo) WithContext(ctx context.Context) ILiveDatumDo {
	return l.withDO(l.DO.WithContext(ctx))
}

func (l liveDatumDo) ReadDB() ILiveDatumDo {
	return l.Clauses(dbresolver.Read)
}

func (l liveDatumDo) WriteDB() ILiveDatumDo {
	return l.Clauses(dbresolver.Write)
}

func (l liveDatumDo) Session(config *gorm.Session) ILiveDatumDo {
	return l.withDO(l.DO.Session(config))
}

func (l liveDatumDo) Clauses(conds ...clause.Expression) ILiveDatumDo {
	return l.withDO(l.DO.Clauses(conds...))
}

func (l liveDatumDo) Returning(value interface{}, columns ...string) ILiveDatumDo {
	return l.withDO(l.DO.Returning(value, columns...))
}

func (l liveDatumDo) Not(conds ...gen.Condition) ILiveDatumDo {
	return l.withDO(l.DO.Not(conds...))
}

func (l liveDatumDo) Or(conds ...gen.Condition) ILiveDatumDo {
	return l.withDO(l.DO.Or(conds...))
}

func (l liveDatumDo) Select(conds ...field.Expr) ILiveDatumDo {
	return l.withDO(l.DO.Select(conds...))
}

func (l liveDatumDo) Where(conds ...gen.Condition) ILiveDatumDo {
	return l.withDO(l.DO.Where(conds...))
}

func (l liveDatumDo) Order(conds ...field.Expr) ILiveDatumDo {
	return l.withDO(l.DO.Order(conds...))
}

func (l liveDatumDo) Distinct(cols ...field.Expr) ILiveDatumDo {
	return l.withDO(l.DO.Distinct(cols...))
}

func (l liveDatumDo) Omit(cols ...field.Expr) ILiveDatumDo {
	return l.withDO(l.DO.Omit(cols...))
}

func (l liveDatumDo) Join(table schema.Tabler, on ...field.Expr) ILiveDatumDo {
	return l.withDO(l.DO.Join(table, on...))
}

func (l liveDatumDo) LeftJoin(table schema.Tabler, on ...field.Expr) ILiveDatumDo {
	return l.withDO(l.DO.LeftJoin(table, on...))
}

func (l liveDatumDo) RightJoin(table schema.Tabler, on ...field.Expr) ILiveDatumDo {
	return l.withDO(l.DO.RightJoin(table, on...))
}

func (l liveDatumDo) Group(cols ...field.Expr) ILiveDatumDo {
	return l.withDO(l.DO.Group(cols...))
}

func (l liveDatumDo) Having(conds ...gen.Condition) ILiveDatumDo {
	return l.withDO(l.DO.Having(conds...))
}

func (l liveDatumDo) Limit(limit int) ILiveDatumDo {
	return l.withDO(l.DO.Limit(limit))
}

func (l liveDatumDo) Offset(offset int) ILiveDatumDo {
	return l.withDO(l.DO.Offset(offset))
}

func (l liveDatumDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ILiveDatumDo {
	return l.withDO(l.DO.Scopes(funcs...))
}

func (l liveDatumDo) Unscoped() ILiveDatumDo {
	return l.withDO(l.DO.Unscoped())
}

func (l liveDatumDo) Create(values ...*model.LiveDatum) error {
	if len(values) == 0 {
		return nil
	}
	return l.DO.Create(values)
}

func (l liveDatumDo) CreateInBatches(values []*model.LiveDatum, batchSize int) error {
	return l.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (l liveDatumDo) Save(values ...*model.LiveDatum) error {
	if len(values) == 0 {
		return nil
	}
	return l.DO.Save(values)
}

func (l liveDatumDo) First() (*model.LiveDatum, error) {
	if result, err := l.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.LiveDatum), nil
	}
}

func (l liveDatumDo) Take() (*model.LiveDatum, error) {
	if result, err := l.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.LiveDatum), nil
	}
}

func (l liveDatumDo) Last() (*model.LiveDatum, error) {
	if result, err := l.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.LiveDatum), nil
	}
}

func (l liveDatumDo) Find() ([]*model.LiveDatum, error) {
	result, err := l.DO.Find()
	return result.([]*model.LiveDatum), err
}

func (l liveDatumDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.LiveDatum, err error) {
	buf := make([]*model.LiveDatum, 0, batchSize)
	err = l.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (l liveDatumDo) FindInBatches(result *[]*model.LiveDatum, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return l.DO.FindInBatches(result, batchSize, fc)
}

func (l liveDatumDo) Attrs(attrs ...field.AssignExpr) ILiveDatumDo {
	return l.withDO(l.DO.Attrs(attrs...))
}

func (l liveDatumDo) Assign(attrs ...field.AssignExpr) ILiveDatumDo {
	return l.withDO(l.DO.Assign(attrs...))
}

func (l liveDatumDo) Joins(fields ...field.RelationField) ILiveDatumDo {
	for _, _f := range fields {
		l = *l.withDO(l.DO.Joins(_f))
	}
	return &l
}

func (l liveDatumDo) Preload(fields ...field.RelationField) ILiveDatumDo {
	for _, _f := range fields {
		l = *l.withDO(l.DO.Preload(_f))
	}
	return &l
}

func (l liveDatumDo) FirstOrInit() (*model.LiveDatum, error) {
	if result, err := l.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.LiveDatum), nil
	}
}

func (l liveDatumDo) FirstOrCreate() (*model.LiveDatum, error) {
	if result, err := l.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.LiveDatum), nil
	}
}

func (l liveDatumDo) FindByPage(offset int, limit int) (result []*model.LiveDatum, count int64, err error) {
	result, err = l.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = l.Offset(-1).Limit(-1).Count()
	return
}

func (l liveDatumDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = l.Count()
	if err != nil {
		return
	}

	err = l.Offset(offset).Limit(limit).Scan(result)
	return
}

func (l liveDatumDo) Scan(result interface{}) (err error) {
	return l.DO.Scan(result)
}

func (l liveDatumDo) Delete(models ...*model.LiveDatum) (result gen.ResultInfo, err error) {
	return l.DO.Delete(models)
}

func (l *liveDatumDo) withDO(do gen.Dao) *liveDatumDo {
	l.DO = *do.(*gen.DO)
	return l
}
