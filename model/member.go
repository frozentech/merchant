package model

import (
	"fmt"
	"time"

	"github.com/frozentech/merchant/query"
	"github.com/jmoiron/sqlx"
)

const (
	// TableNameMember defines the member table
	TableNameMember = "member"
)

// TxnUpdateDisabledFields ...
var TxnUpdateDisabledFields = []string{}

// Member ...
type Member struct {
	ID         string `json:"id,omitempty" db:"id"`
	MerchantID string `json:"merchant_id,omitempty" db:"merchant_id"`
	Name       string `json:"name" db:"name"`
	Email      string `json:"email" db:"email"`
	CreateAt   string `json:"created_at,omitempty" db:"created_at"`
}

// NewMember ...
func NewMember() *Member {
	return &Member{
		ID:       GenerateUUID(),
		CreateAt: time.Now().Format(MYSQLTimestampFormat),
	}
}

// DB returns the DB used to connect the Member table
func (me Member) DB() *sqlx.DB {
	return DB
}

// FindByMechantID Find By Mechant ID
func (me *Member) FindByMechantID(id string, page int, limit int) (members []Member, err error) {
	members = []Member{}
	sql := fmt.Sprintf(query.SelectBuilder(me, TableNameMember, false)+" WHERE `merchant_id` = ? ORDER BY `created_at` DESC LIMIT %d OFFSET %d", limit, (page-1)*limit)
	if err = me.DB().Select(&members, sql, id); err == nil {
		return
	}

	return
}

// FindID Find By ID
func (me *Member) FindID(id string) (err error) {
	sql := fmt.Sprintf(query.SelectBuilder(me, TableNameMember, false) + " WHERE `id` = ?")
	if err = me.DB().Get(me, sql, id); err == nil {
		return
	}
	return
}

// Delete Delete record
func (me *Member) Delete() (err error) {
	return me.txDelete()
}

func (me *Member) txDelete() error {
	sql := fmt.Sprintf("DELETE FROM `%s` WHERE `%s`.`id` = ?", TableNameMember, TableNameMember)
	_, err := me.DB().Exec(sql, me.ID)
	return err
}

// Update ...
func (me *Member) Update() error {
	return me.txUpdate()
}

func (me *Member) txUpdate() error {
	sql, params := query.UpdateBuilder(me, TableNameMember, TxnUpdateDisabledFields...)
	sql = fmt.Sprintf("%s WHERE `%s`.`id` = ?", sql, TableNameMember)
	params = append(params, me.ID)
	_, err := me.DB().Exec(sql, params...)
	return err
}

// Create registers a new MemberLog to the DB matching Domain
func (me *Member) Create() error {
	return me.txCreate()
}

func (me *Member) txCreate() error {
	sql, params := query.InsertBuilder(me, TableNameMember, false)
	_, err := me.DB().Exec(sql, params...)
	return err
}
