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

func newOperationLog(db *gorm.DB, opts ...gen.DOOption) operationLog {
	_operationLog := operationLog{}

	_operationLog.operationLogDo.UseDB(db, opts...)
	_operationLog.operationLogDo.UseModel(&model.OperationLog{})

	tableName := _operationLog.operationLogDo.TableName()
	_operationLog.ALL = field.NewAsterisk(tableName)
	_operationLog.ID = field.NewInt64(tableName, "id")
	_operationLog.CreatedAt = field.NewTime(tableName, "created_at")
	_operationLog.UpdatedAt = field.NewTime(tableName, "updated_at")
	_operationLog.DeletedAt = field.NewField(tableName, "deleted_at")
	_operationLog.Username = field.NewString(tableName, "username")
	_operationLog.IP = field.NewString(tableName, "ip")
	_operationLog.IPLocation = field.NewString(tableName, "ip_location")
	_operationLog.Method = field.NewString(tableName, "method")
	_operationLog.Path = field.NewString(tableName, "path")
	_operationLog.Desc = field.NewString(tableName, "desc")
	_operationLog.Status = field.NewInt64(tableName, "status")
	_operationLog.StartTime = field.NewTime(tableName, "start_time")
	_operationLog.TimeCost = field.NewInt64(tableName, "time_cost")
	_operationLog.UserAgent = field.NewString(tableName, "user_agent")

	_operationLog.fillFieldMap()

	return _operationLog
}

type operationLog struct {
	operationLogDo

	ALL        field.Asterisk
	ID         field.Int64
	CreatedAt  field.Time
	UpdatedAt  field.Time
	DeletedAt  field.Field
	Username   field.String // '用户登录名'
	IP         field.String // 'Ip地址'
	IPLocation field.String // 'Ip所在地'
	Method     field.String // '请求方式'
	Path       field.String // '访问路径'
	Desc       field.String // '说明'
	Status     field.Int64  // '响应状态码'
	StartTime  field.Time   // '发起时间'
	TimeCost   field.Int64  // '请求耗时(ms)'
	UserAgent  field.String // '浏览器标识'

	fieldMap map[string]field.Expr
}

func (o operationLog) Table(newTableName string) *operationLog {
	o.operationLogDo.UseTable(newTableName)
	return o.updateTableName(newTableName)
}

func (o operationLog) As(alias string) *operationLog {
	o.operationLogDo.DO = *(o.operationLogDo.As(alias).(*gen.DO))
	return o.updateTableName(alias)
}

func (o *operationLog) updateTableName(table string) *operationLog {
	o.ALL = field.NewAsterisk(table)
	o.ID = field.NewInt64(table, "id")
	o.CreatedAt = field.NewTime(table, "created_at")
	o.UpdatedAt = field.NewTime(table, "updated_at")
	o.DeletedAt = field.NewField(table, "deleted_at")
	o.Username = field.NewString(table, "username")
	o.IP = field.NewString(table, "ip")
	o.IPLocation = field.NewString(table, "ip_location")
	o.Method = field.NewString(table, "method")
	o.Path = field.NewString(table, "path")
	o.Desc = field.NewString(table, "desc")
	o.Status = field.NewInt64(table, "status")
	o.StartTime = field.NewTime(table, "start_time")
	o.TimeCost = field.NewInt64(table, "time_cost")
	o.UserAgent = field.NewString(table, "user_agent")

	o.fillFieldMap()

	return o
}

func (o *operationLog) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := o.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (o *operationLog) fillFieldMap() {
	o.fieldMap = make(map[string]field.Expr, 14)
	o.fieldMap["id"] = o.ID
	o.fieldMap["created_at"] = o.CreatedAt
	o.fieldMap["updated_at"] = o.UpdatedAt
	o.fieldMap["deleted_at"] = o.DeletedAt
	o.fieldMap["username"] = o.Username
	o.fieldMap["ip"] = o.IP
	o.fieldMap["ip_location"] = o.IPLocation
	o.fieldMap["method"] = o.Method
	o.fieldMap["path"] = o.Path
	o.fieldMap["desc"] = o.Desc
	o.fieldMap["status"] = o.Status
	o.fieldMap["start_time"] = o.StartTime
	o.fieldMap["time_cost"] = o.TimeCost
	o.fieldMap["user_agent"] = o.UserAgent
}

func (o operationLog) clone(db *gorm.DB) operationLog {
	o.operationLogDo.ReplaceConnPool(db.Statement.ConnPool)
	return o
}

func (o operationLog) replaceDB(db *gorm.DB) operationLog {
	o.operationLogDo.ReplaceDB(db)
	return o
}

type operationLogDo struct{ gen.DO }

type IOperationLogDo interface {
	gen.SubQuery
	Debug() IOperationLogDo
	WithContext(ctx context.Context) IOperationLogDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IOperationLogDo
	WriteDB() IOperationLogDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IOperationLogDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IOperationLogDo
	Not(conds ...gen.Condition) IOperationLogDo
	Or(conds ...gen.Condition) IOperationLogDo
	Select(conds ...field.Expr) IOperationLogDo
	Where(conds ...gen.Condition) IOperationLogDo
	Order(conds ...field.Expr) IOperationLogDo
	Distinct(cols ...field.Expr) IOperationLogDo
	Omit(cols ...field.Expr) IOperationLogDo
	Join(table schema.Tabler, on ...field.Expr) IOperationLogDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IOperationLogDo
	RightJoin(table schema.Tabler, on ...field.Expr) IOperationLogDo
	Group(cols ...field.Expr) IOperationLogDo
	Having(conds ...gen.Condition) IOperationLogDo
	Limit(limit int) IOperationLogDo
	Offset(offset int) IOperationLogDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IOperationLogDo
	Unscoped() IOperationLogDo
	Create(values ...*model.OperationLog) error
	CreateInBatches(values []*model.OperationLog, batchSize int) error
	Save(values ...*model.OperationLog) error
	First() (*model.OperationLog, error)
	Take() (*model.OperationLog, error)
	Last() (*model.OperationLog, error)
	Find() ([]*model.OperationLog, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.OperationLog, err error)
	FindInBatches(result *[]*model.OperationLog, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.OperationLog) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IOperationLogDo
	Assign(attrs ...field.AssignExpr) IOperationLogDo
	Joins(fields ...field.RelationField) IOperationLogDo
	Preload(fields ...field.RelationField) IOperationLogDo
	FirstOrInit() (*model.OperationLog, error)
	FirstOrCreate() (*model.OperationLog, error)
	FindByPage(offset int, limit int) (result []*model.OperationLog, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IOperationLogDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (o operationLogDo) Debug() IOperationLogDo {
	return o.withDO(o.DO.Debug())
}

func (o operationLogDo) WithContext(ctx context.Context) IOperationLogDo {
	return o.withDO(o.DO.WithContext(ctx))
}

func (o operationLogDo) ReadDB() IOperationLogDo {
	return o.Clauses(dbresolver.Read)
}

func (o operationLogDo) WriteDB() IOperationLogDo {
	return o.Clauses(dbresolver.Write)
}

func (o operationLogDo) Session(config *gorm.Session) IOperationLogDo {
	return o.withDO(o.DO.Session(config))
}

func (o operationLogDo) Clauses(conds ...clause.Expression) IOperationLogDo {
	return o.withDO(o.DO.Clauses(conds...))
}

func (o operationLogDo) Returning(value interface{}, columns ...string) IOperationLogDo {
	return o.withDO(o.DO.Returning(value, columns...))
}

func (o operationLogDo) Not(conds ...gen.Condition) IOperationLogDo {
	return o.withDO(o.DO.Not(conds...))
}

func (o operationLogDo) Or(conds ...gen.Condition) IOperationLogDo {
	return o.withDO(o.DO.Or(conds...))
}

func (o operationLogDo) Select(conds ...field.Expr) IOperationLogDo {
	return o.withDO(o.DO.Select(conds...))
}

func (o operationLogDo) Where(conds ...gen.Condition) IOperationLogDo {
	return o.withDO(o.DO.Where(conds...))
}

func (o operationLogDo) Order(conds ...field.Expr) IOperationLogDo {
	return o.withDO(o.DO.Order(conds...))
}

func (o operationLogDo) Distinct(cols ...field.Expr) IOperationLogDo {
	return o.withDO(o.DO.Distinct(cols...))
}

func (o operationLogDo) Omit(cols ...field.Expr) IOperationLogDo {
	return o.withDO(o.DO.Omit(cols...))
}

func (o operationLogDo) Join(table schema.Tabler, on ...field.Expr) IOperationLogDo {
	return o.withDO(o.DO.Join(table, on...))
}

func (o operationLogDo) LeftJoin(table schema.Tabler, on ...field.Expr) IOperationLogDo {
	return o.withDO(o.DO.LeftJoin(table, on...))
}

func (o operationLogDo) RightJoin(table schema.Tabler, on ...field.Expr) IOperationLogDo {
	return o.withDO(o.DO.RightJoin(table, on...))
}

func (o operationLogDo) Group(cols ...field.Expr) IOperationLogDo {
	return o.withDO(o.DO.Group(cols...))
}

func (o operationLogDo) Having(conds ...gen.Condition) IOperationLogDo {
	return o.withDO(o.DO.Having(conds...))
}

func (o operationLogDo) Limit(limit int) IOperationLogDo {
	return o.withDO(o.DO.Limit(limit))
}

func (o operationLogDo) Offset(offset int) IOperationLogDo {
	return o.withDO(o.DO.Offset(offset))
}

func (o operationLogDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IOperationLogDo {
	return o.withDO(o.DO.Scopes(funcs...))
}

func (o operationLogDo) Unscoped() IOperationLogDo {
	return o.withDO(o.DO.Unscoped())
}

func (o operationLogDo) Create(values ...*model.OperationLog) error {
	if len(values) == 0 {
		return nil
	}
	return o.DO.Create(values)
}

func (o operationLogDo) CreateInBatches(values []*model.OperationLog, batchSize int) error {
	return o.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (o operationLogDo) Save(values ...*model.OperationLog) error {
	if len(values) == 0 {
		return nil
	}
	return o.DO.Save(values)
}

func (o operationLogDo) First() (*model.OperationLog, error) {
	if result, err := o.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.OperationLog), nil
	}
}

func (o operationLogDo) Take() (*model.OperationLog, error) {
	if result, err := o.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.OperationLog), nil
	}
}

func (o operationLogDo) Last() (*model.OperationLog, error) {
	if result, err := o.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.OperationLog), nil
	}
}

func (o operationLogDo) Find() ([]*model.OperationLog, error) {
	result, err := o.DO.Find()
	return result.([]*model.OperationLog), err
}

func (o operationLogDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.OperationLog, err error) {
	buf := make([]*model.OperationLog, 0, batchSize)
	err = o.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (o operationLogDo) FindInBatches(result *[]*model.OperationLog, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return o.DO.FindInBatches(result, batchSize, fc)
}

func (o operationLogDo) Attrs(attrs ...field.AssignExpr) IOperationLogDo {
	return o.withDO(o.DO.Attrs(attrs...))
}

func (o operationLogDo) Assign(attrs ...field.AssignExpr) IOperationLogDo {
	return o.withDO(o.DO.Assign(attrs...))
}

func (o operationLogDo) Joins(fields ...field.RelationField) IOperationLogDo {
	for _, _f := range fields {
		o = *o.withDO(o.DO.Joins(_f))
	}
	return &o
}

func (o operationLogDo) Preload(fields ...field.RelationField) IOperationLogDo {
	for _, _f := range fields {
		o = *o.withDO(o.DO.Preload(_f))
	}
	return &o
}

func (o operationLogDo) FirstOrInit() (*model.OperationLog, error) {
	if result, err := o.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.OperationLog), nil
	}
}

func (o operationLogDo) FirstOrCreate() (*model.OperationLog, error) {
	if result, err := o.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.OperationLog), nil
	}
}

func (o operationLogDo) FindByPage(offset int, limit int) (result []*model.OperationLog, count int64, err error) {
	result, err = o.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = o.Offset(-1).Limit(-1).Count()
	return
}

func (o operationLogDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = o.Count()
	if err != nil {
		return
	}

	err = o.Offset(offset).Limit(limit).Scan(result)
	return
}

func (o operationLogDo) Scan(result interface{}) (err error) {
	return o.DO.Scan(result)
}

func (o operationLogDo) Delete(models ...*model.OperationLog) (result gen.ResultInfo, err error) {
	return o.DO.Delete(models)
}

func (o *operationLogDo) withDO(do gen.Dao) *operationLogDo {
	o.DO = *do.(*gen.DO)
	return o
}
