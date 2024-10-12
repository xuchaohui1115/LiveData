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

func newViolation(db *gorm.DB, opts ...gen.DOOption) violation {
	_violation := violation{}

	_violation.violationDo.UseDB(db, opts...)
	_violation.violationDo.UseModel(&model.Violation{})

	tableName := _violation.violationDo.TableName()
	_violation.ALL = field.NewAsterisk(tableName)
	_violation.ID = field.NewInt64(tableName, "id")
	_violation.HostNickname = field.NewString(tableName, "host_nickname")
	_violation.HostID = field.NewInt64(tableName, "host_id")
	_violation.SecurityCode = field.NewString(tableName, "security_code")
	_violation.AccountStatus = field.NewString(tableName, "account_status")
	_violation.ViolationNumber = field.NewString(tableName, "violation_number")
	_violation.ViolationTime = field.NewInt64(tableName, "violation_time")
	_violation.ViolationReason = field.NewString(tableName, "violation_reason")
	_violation.ViolationImpact = field.NewString(tableName, "violation_impact")
	_violation.ViolationLiveRoomID = field.NewString(tableName, "violation_live_room_id")
	_violation.CreatedAt = field.NewInt64(tableName, "created_at")

	_violation.fillFieldMap()

	return _violation
}

// violation 违规记录表
type violation struct {
	violationDo

	ALL                 field.Asterisk
	ID                  field.Int64  // 主键ID
	HostNickname        field.String // 主播昵称
	HostID              field.Int64  // 主播ID
	SecurityCode        field.String // 安全码
	AccountStatus       field.String // 账号状态
	ViolationNumber     field.String // 违规编号
	ViolationTime       field.Int64  // 违规时间（Unix 时间戳）
	ViolationReason     field.String // 违规原因
	ViolationImpact     field.String // 违规影响
	ViolationLiveRoomID field.String // 违规对象是直播间的ID
	CreatedAt           field.Int64  // 创建时间（Unix 时间戳）

	fieldMap map[string]field.Expr
}

func (v violation) Table(newTableName string) *violation {
	v.violationDo.UseTable(newTableName)
	return v.updateTableName(newTableName)
}

func (v violation) As(alias string) *violation {
	v.violationDo.DO = *(v.violationDo.As(alias).(*gen.DO))
	return v.updateTableName(alias)
}

func (v *violation) updateTableName(table string) *violation {
	v.ALL = field.NewAsterisk(table)
	v.ID = field.NewInt64(table, "id")
	v.HostNickname = field.NewString(table, "host_nickname")
	v.HostID = field.NewInt64(table, "host_id")
	v.SecurityCode = field.NewString(table, "security_code")
	v.AccountStatus = field.NewString(table, "account_status")
	v.ViolationNumber = field.NewString(table, "violation_number")
	v.ViolationTime = field.NewInt64(table, "violation_time")
	v.ViolationReason = field.NewString(table, "violation_reason")
	v.ViolationImpact = field.NewString(table, "violation_impact")
	v.ViolationLiveRoomID = field.NewString(table, "violation_live_room_id")
	v.CreatedAt = field.NewInt64(table, "created_at")

	v.fillFieldMap()

	return v
}

func (v *violation) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := v.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (v *violation) fillFieldMap() {
	v.fieldMap = make(map[string]field.Expr, 11)
	v.fieldMap["id"] = v.ID
	v.fieldMap["host_nickname"] = v.HostNickname
	v.fieldMap["host_id"] = v.HostID
	v.fieldMap["security_code"] = v.SecurityCode
	v.fieldMap["account_status"] = v.AccountStatus
	v.fieldMap["violation_number"] = v.ViolationNumber
	v.fieldMap["violation_time"] = v.ViolationTime
	v.fieldMap["violation_reason"] = v.ViolationReason
	v.fieldMap["violation_impact"] = v.ViolationImpact
	v.fieldMap["violation_live_room_id"] = v.ViolationLiveRoomID
	v.fieldMap["created_at"] = v.CreatedAt
}

func (v violation) clone(db *gorm.DB) violation {
	v.violationDo.ReplaceConnPool(db.Statement.ConnPool)
	return v
}

func (v violation) replaceDB(db *gorm.DB) violation {
	v.violationDo.ReplaceDB(db)
	return v
}

type violationDo struct{ gen.DO }

type IViolationDo interface {
	gen.SubQuery
	Debug() IViolationDo
	WithContext(ctx context.Context) IViolationDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IViolationDo
	WriteDB() IViolationDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IViolationDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IViolationDo
	Not(conds ...gen.Condition) IViolationDo
	Or(conds ...gen.Condition) IViolationDo
	Select(conds ...field.Expr) IViolationDo
	Where(conds ...gen.Condition) IViolationDo
	Order(conds ...field.Expr) IViolationDo
	Distinct(cols ...field.Expr) IViolationDo
	Omit(cols ...field.Expr) IViolationDo
	Join(table schema.Tabler, on ...field.Expr) IViolationDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IViolationDo
	RightJoin(table schema.Tabler, on ...field.Expr) IViolationDo
	Group(cols ...field.Expr) IViolationDo
	Having(conds ...gen.Condition) IViolationDo
	Limit(limit int) IViolationDo
	Offset(offset int) IViolationDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IViolationDo
	Unscoped() IViolationDo
	Create(values ...*model.Violation) error
	CreateInBatches(values []*model.Violation, batchSize int) error
	Save(values ...*model.Violation) error
	First() (*model.Violation, error)
	Take() (*model.Violation, error)
	Last() (*model.Violation, error)
	Find() ([]*model.Violation, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Violation, err error)
	FindInBatches(result *[]*model.Violation, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Violation) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IViolationDo
	Assign(attrs ...field.AssignExpr) IViolationDo
	Joins(fields ...field.RelationField) IViolationDo
	Preload(fields ...field.RelationField) IViolationDo
	FirstOrInit() (*model.Violation, error)
	FirstOrCreate() (*model.Violation, error)
	FindByPage(offset int, limit int) (result []*model.Violation, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IViolationDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (v violationDo) Debug() IViolationDo {
	return v.withDO(v.DO.Debug())
}

func (v violationDo) WithContext(ctx context.Context) IViolationDo {
	return v.withDO(v.DO.WithContext(ctx))
}

func (v violationDo) ReadDB() IViolationDo {
	return v.Clauses(dbresolver.Read)
}

func (v violationDo) WriteDB() IViolationDo {
	return v.Clauses(dbresolver.Write)
}

func (v violationDo) Session(config *gorm.Session) IViolationDo {
	return v.withDO(v.DO.Session(config))
}

func (v violationDo) Clauses(conds ...clause.Expression) IViolationDo {
	return v.withDO(v.DO.Clauses(conds...))
}

func (v violationDo) Returning(value interface{}, columns ...string) IViolationDo {
	return v.withDO(v.DO.Returning(value, columns...))
}

func (v violationDo) Not(conds ...gen.Condition) IViolationDo {
	return v.withDO(v.DO.Not(conds...))
}

func (v violationDo) Or(conds ...gen.Condition) IViolationDo {
	return v.withDO(v.DO.Or(conds...))
}

func (v violationDo) Select(conds ...field.Expr) IViolationDo {
	return v.withDO(v.DO.Select(conds...))
}

func (v violationDo) Where(conds ...gen.Condition) IViolationDo {
	return v.withDO(v.DO.Where(conds...))
}

func (v violationDo) Order(conds ...field.Expr) IViolationDo {
	return v.withDO(v.DO.Order(conds...))
}

func (v violationDo) Distinct(cols ...field.Expr) IViolationDo {
	return v.withDO(v.DO.Distinct(cols...))
}

func (v violationDo) Omit(cols ...field.Expr) IViolationDo {
	return v.withDO(v.DO.Omit(cols...))
}

func (v violationDo) Join(table schema.Tabler, on ...field.Expr) IViolationDo {
	return v.withDO(v.DO.Join(table, on...))
}

func (v violationDo) LeftJoin(table schema.Tabler, on ...field.Expr) IViolationDo {
	return v.withDO(v.DO.LeftJoin(table, on...))
}

func (v violationDo) RightJoin(table schema.Tabler, on ...field.Expr) IViolationDo {
	return v.withDO(v.DO.RightJoin(table, on...))
}

func (v violationDo) Group(cols ...field.Expr) IViolationDo {
	return v.withDO(v.DO.Group(cols...))
}

func (v violationDo) Having(conds ...gen.Condition) IViolationDo {
	return v.withDO(v.DO.Having(conds...))
}

func (v violationDo) Limit(limit int) IViolationDo {
	return v.withDO(v.DO.Limit(limit))
}

func (v violationDo) Offset(offset int) IViolationDo {
	return v.withDO(v.DO.Offset(offset))
}

func (v violationDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IViolationDo {
	return v.withDO(v.DO.Scopes(funcs...))
}

func (v violationDo) Unscoped() IViolationDo {
	return v.withDO(v.DO.Unscoped())
}

func (v violationDo) Create(values ...*model.Violation) error {
	if len(values) == 0 {
		return nil
	}
	return v.DO.Create(values)
}

func (v violationDo) CreateInBatches(values []*model.Violation, batchSize int) error {
	return v.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (v violationDo) Save(values ...*model.Violation) error {
	if len(values) == 0 {
		return nil
	}
	return v.DO.Save(values)
}

func (v violationDo) First() (*model.Violation, error) {
	if result, err := v.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Violation), nil
	}
}

func (v violationDo) Take() (*model.Violation, error) {
	if result, err := v.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Violation), nil
	}
}

func (v violationDo) Last() (*model.Violation, error) {
	if result, err := v.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Violation), nil
	}
}

func (v violationDo) Find() ([]*model.Violation, error) {
	result, err := v.DO.Find()
	return result.([]*model.Violation), err
}

func (v violationDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Violation, err error) {
	buf := make([]*model.Violation, 0, batchSize)
	err = v.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (v violationDo) FindInBatches(result *[]*model.Violation, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return v.DO.FindInBatches(result, batchSize, fc)
}

func (v violationDo) Attrs(attrs ...field.AssignExpr) IViolationDo {
	return v.withDO(v.DO.Attrs(attrs...))
}

func (v violationDo) Assign(attrs ...field.AssignExpr) IViolationDo {
	return v.withDO(v.DO.Assign(attrs...))
}

func (v violationDo) Joins(fields ...field.RelationField) IViolationDo {
	for _, _f := range fields {
		v = *v.withDO(v.DO.Joins(_f))
	}
	return &v
}

func (v violationDo) Preload(fields ...field.RelationField) IViolationDo {
	for _, _f := range fields {
		v = *v.withDO(v.DO.Preload(_f))
	}
	return &v
}

func (v violationDo) FirstOrInit() (*model.Violation, error) {
	if result, err := v.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Violation), nil
	}
}

func (v violationDo) FirstOrCreate() (*model.Violation, error) {
	if result, err := v.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Violation), nil
	}
}

func (v violationDo) FindByPage(offset int, limit int) (result []*model.Violation, count int64, err error) {
	result, err = v.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = v.Offset(-1).Limit(-1).Count()
	return
}

func (v violationDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = v.Count()
	if err != nil {
		return
	}

	err = v.Offset(offset).Limit(limit).Scan(result)
	return
}

func (v violationDo) Scan(result interface{}) (err error) {
	return v.DO.Scan(result)
}

func (v violationDo) Delete(models ...*model.Violation) (result gen.ResultInfo, err error) {
	return v.DO.Delete(models)
}

func (v *violationDo) withDO(do gen.Dao) *violationDo {
	v.DO = *do.(*gen.DO)
	return v
}
