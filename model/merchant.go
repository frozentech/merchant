package model

import (
	"fmt"
	"time"

	"github.com/frozentech/database"
	"github.com/frozentech/query"
	"github.com/jmoiron/sqlx"
)

const (
	// TableNameMerchant defines the Merchant table
	TableNameMerchant = "merchant"
)

// Merchant ...
type Merchant struct {
	ID       string `json:"id,omitempty" db:"id"`
	Name     string `json:"name" db:"name"`
	Image    string `json:"image,omitempty" db:"image"`
	CreateAt string `json:"created_at,omitempty" db:"created_at"`
}

// NewMerchant ...
func NewMerchant() *Merchant {
	return &Merchant{
		ID:       query.GenerateUUID(),
		CreateAt: time.Now().Format(query.MYSQLTimestampFormat),
	}
}

// DB returns the DB used to connect the Transaction table
func (me Merchant) DB() *sqlx.DB {
	return database.GetConnection()
}

// FindID Find By ID
func (me *Merchant) FindID(id string) (err error) {
	sql := fmt.Sprintf(query.SelectBuilder(me, TableNameMerchant, false) + " WHERE `id` = ?")
	if err = me.DB().Get(me, sql, id); err == nil {
		return
	}
	return
}

// FindAll Find All
func (me *Merchant) FindAll(page int, limit int) (merchants []Merchant, err error) {
	merchants = []Merchant{}
	sql := fmt.Sprintf(query.SelectBuilder(me, TableNameMerchant, false)+" ORDER BY `created_at` DESC LIMIT %d OFFSET %d", limit, (page-1)*limit)
	if err = me.DB().Select(&merchants, sql); err == nil {
		return
	}

	return
}

// Delete Delete record
func (me *Merchant) Delete() (err error) {
	return me.txDelete()
}

func (me *Merchant) txDelete() error {
	sql := fmt.Sprintf("DELETE FROM `%s` WHERE `%s`.`id` = ?", TableNameMerchant, TableNameMerchant)
	_, err := me.DB().Exec(sql, me.ID)
	return err
}

// Update ...
func (me *Merchant) Update() error {
	return me.txUpdate()
}

func (me *Merchant) txUpdate() error {
	sql, params := query.UpdateBuilder(me, TableNameMerchant, TxnUpdateDisabledFields...)
	sql = fmt.Sprintf("%s WHERE `%s`.`id` = ?", sql, TableNameMerchant)
	params = append(params, me.ID)
	_, err := me.DB().Exec(sql, params...)
	return err
}

// Create registers a new Merchant
func (me *Merchant) Create() error {
	return me.txCreate()
}

func (me *Merchant) txCreate() error {
	sql, params := query.InsertBuilder(me, TableNameMerchant, false)
	_, err := me.DB().Exec(sql, params...)
	return err
}
