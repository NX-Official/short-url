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

	"gin-rush-template/internal/model"
)

func newInterview(db *gorm.DB, opts ...gen.DOOption) interview {
	_interview := interview{}

	_interview.interviewDo.UseDB(db, opts...)
	_interview.interviewDo.UseModel(&model.Interview{})

	tableName := _interview.interviewDo.TableName()
	_interview.ALL = field.NewAsterisk(tableName)
	_interview.ID = field.NewUint(tableName, "id")
	_interview.CreatedAt = field.NewTime(tableName, "created_at")
	_interview.UpdatedAt = field.NewTime(tableName, "updated_at")
	_interview.DeletedAt = field.NewField(tableName, "deleted_at")
	_interview.JobID = field.NewUint(tableName, "job_id")
	_interview.SlotID = field.NewUint(tableName, "slot_id")
	_interview.ApplicationID = field.NewUint(tableName, "application_id")
	_interview.IntervieweeID = field.NewUint(tableName, "interviewee_id")
	_interview.StartAt = field.NewTime(tableName, "start_at")
	_interview.EndAt = field.NewTime(tableName, "end_at")

	_interview.fillFieldMap()

	return _interview
}

type interview struct {
	interviewDo

	ALL           field.Asterisk
	ID            field.Uint
	CreatedAt     field.Time
	UpdatedAt     field.Time
	DeletedAt     field.Field
	JobID         field.Uint
	SlotID        field.Uint
	ApplicationID field.Uint
	IntervieweeID field.Uint
	StartAt       field.Time
	EndAt         field.Time

	fieldMap map[string]field.Expr
}

func (i interview) Table(newTableName string) *interview {
	i.interviewDo.UseTable(newTableName)
	return i.updateTableName(newTableName)
}

func (i interview) As(alias string) *interview {
	i.interviewDo.DO = *(i.interviewDo.As(alias).(*gen.DO))
	return i.updateTableName(alias)
}

func (i *interview) updateTableName(table string) *interview {
	i.ALL = field.NewAsterisk(table)
	i.ID = field.NewUint(table, "id")
	i.CreatedAt = field.NewTime(table, "created_at")
	i.UpdatedAt = field.NewTime(table, "updated_at")
	i.DeletedAt = field.NewField(table, "deleted_at")
	i.JobID = field.NewUint(table, "job_id")
	i.SlotID = field.NewUint(table, "slot_id")
	i.ApplicationID = field.NewUint(table, "application_id")
	i.IntervieweeID = field.NewUint(table, "interviewee_id")
	i.StartAt = field.NewTime(table, "start_at")
	i.EndAt = field.NewTime(table, "end_at")

	i.fillFieldMap()

	return i
}

func (i *interview) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := i.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (i *interview) fillFieldMap() {
	i.fieldMap = make(map[string]field.Expr, 10)
	i.fieldMap["id"] = i.ID
	i.fieldMap["created_at"] = i.CreatedAt
	i.fieldMap["updated_at"] = i.UpdatedAt
	i.fieldMap["deleted_at"] = i.DeletedAt
	i.fieldMap["job_id"] = i.JobID
	i.fieldMap["slot_id"] = i.SlotID
	i.fieldMap["application_id"] = i.ApplicationID
	i.fieldMap["interviewee_id"] = i.IntervieweeID
	i.fieldMap["start_at"] = i.StartAt
	i.fieldMap["end_at"] = i.EndAt
}

func (i interview) clone(db *gorm.DB) interview {
	i.interviewDo.ReplaceConnPool(db.Statement.ConnPool)
	return i
}

func (i interview) replaceDB(db *gorm.DB) interview {
	i.interviewDo.ReplaceDB(db)
	return i
}

type interviewDo struct{ gen.DO }

func (i interviewDo) Debug() *interviewDo {
	return i.withDO(i.DO.Debug())
}

func (i interviewDo) WithContext(ctx context.Context) *interviewDo {
	return i.withDO(i.DO.WithContext(ctx))
}

func (i interviewDo) ReadDB() *interviewDo {
	return i.Clauses(dbresolver.Read)
}

func (i interviewDo) WriteDB() *interviewDo {
	return i.Clauses(dbresolver.Write)
}

func (i interviewDo) Session(config *gorm.Session) *interviewDo {
	return i.withDO(i.DO.Session(config))
}

func (i interviewDo) Clauses(conds ...clause.Expression) *interviewDo {
	return i.withDO(i.DO.Clauses(conds...))
}

func (i interviewDo) Returning(value interface{}, columns ...string) *interviewDo {
	return i.withDO(i.DO.Returning(value, columns...))
}

func (i interviewDo) Not(conds ...gen.Condition) *interviewDo {
	return i.withDO(i.DO.Not(conds...))
}

func (i interviewDo) Or(conds ...gen.Condition) *interviewDo {
	return i.withDO(i.DO.Or(conds...))
}

func (i interviewDo) Select(conds ...field.Expr) *interviewDo {
	return i.withDO(i.DO.Select(conds...))
}

func (i interviewDo) Where(conds ...gen.Condition) *interviewDo {
	return i.withDO(i.DO.Where(conds...))
}

func (i interviewDo) Order(conds ...field.Expr) *interviewDo {
	return i.withDO(i.DO.Order(conds...))
}

func (i interviewDo) Distinct(cols ...field.Expr) *interviewDo {
	return i.withDO(i.DO.Distinct(cols...))
}

func (i interviewDo) Omit(cols ...field.Expr) *interviewDo {
	return i.withDO(i.DO.Omit(cols...))
}

func (i interviewDo) Join(table schema.Tabler, on ...field.Expr) *interviewDo {
	return i.withDO(i.DO.Join(table, on...))
}

func (i interviewDo) LeftJoin(table schema.Tabler, on ...field.Expr) *interviewDo {
	return i.withDO(i.DO.LeftJoin(table, on...))
}

func (i interviewDo) RightJoin(table schema.Tabler, on ...field.Expr) *interviewDo {
	return i.withDO(i.DO.RightJoin(table, on...))
}

func (i interviewDo) Group(cols ...field.Expr) *interviewDo {
	return i.withDO(i.DO.Group(cols...))
}

func (i interviewDo) Having(conds ...gen.Condition) *interviewDo {
	return i.withDO(i.DO.Having(conds...))
}

func (i interviewDo) Limit(limit int) *interviewDo {
	return i.withDO(i.DO.Limit(limit))
}

func (i interviewDo) Offset(offset int) *interviewDo {
	return i.withDO(i.DO.Offset(offset))
}

func (i interviewDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *interviewDo {
	return i.withDO(i.DO.Scopes(funcs...))
}

func (i interviewDo) Unscoped() *interviewDo {
	return i.withDO(i.DO.Unscoped())
}

func (i interviewDo) Create(values ...*model.Interview) error {
	if len(values) == 0 {
		return nil
	}
	return i.DO.Create(values)
}

func (i interviewDo) CreateInBatches(values []*model.Interview, batchSize int) error {
	return i.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (i interviewDo) Save(values ...*model.Interview) error {
	if len(values) == 0 {
		return nil
	}
	return i.DO.Save(values)
}

func (i interviewDo) First() (*model.Interview, error) {
	if result, err := i.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Interview), nil
	}
}

func (i interviewDo) Take() (*model.Interview, error) {
	if result, err := i.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Interview), nil
	}
}

func (i interviewDo) Last() (*model.Interview, error) {
	if result, err := i.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Interview), nil
	}
}

func (i interviewDo) Find() ([]*model.Interview, error) {
	result, err := i.DO.Find()
	return result.([]*model.Interview), err
}

func (i interviewDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Interview, err error) {
	buf := make([]*model.Interview, 0, batchSize)
	err = i.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (i interviewDo) FindInBatches(result *[]*model.Interview, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return i.DO.FindInBatches(result, batchSize, fc)
}

func (i interviewDo) Attrs(attrs ...field.AssignExpr) *interviewDo {
	return i.withDO(i.DO.Attrs(attrs...))
}

func (i interviewDo) Assign(attrs ...field.AssignExpr) *interviewDo {
	return i.withDO(i.DO.Assign(attrs...))
}

func (i interviewDo) Joins(fields ...field.RelationField) *interviewDo {
	for _, _f := range fields {
		i = *i.withDO(i.DO.Joins(_f))
	}
	return &i
}

func (i interviewDo) Preload(fields ...field.RelationField) *interviewDo {
	for _, _f := range fields {
		i = *i.withDO(i.DO.Preload(_f))
	}
	return &i
}

func (i interviewDo) FirstOrInit() (*model.Interview, error) {
	if result, err := i.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Interview), nil
	}
}

func (i interviewDo) FirstOrCreate() (*model.Interview, error) {
	if result, err := i.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Interview), nil
	}
}

func (i interviewDo) FindByPage(offset int, limit int) (result []*model.Interview, count int64, err error) {
	result, err = i.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = i.Offset(-1).Limit(-1).Count()
	return
}

func (i interviewDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = i.Count()
	if err != nil {
		return
	}

	err = i.Offset(offset).Limit(limit).Scan(result)
	return
}

func (i interviewDo) Scan(result interface{}) (err error) {
	return i.DO.Scan(result)
}

func (i interviewDo) Delete(models ...*model.Interview) (result gen.ResultInfo, err error) {
	return i.DO.Delete(models)
}

func (i *interviewDo) withDO(do gen.Dao) *interviewDo {
	i.DO = *do.(*gen.DO)
	return i
}
