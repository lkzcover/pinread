//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package table

import (
	"github.com/go-jet/jet/v2/postgres"
)

var Pins = newPinsTable("public", "pins", "")

type pinsTable struct {
	postgres.Table

	// Columns
	UserID   postgres.ColumnInteger
	MsgID    postgres.ColumnInteger
	IsActive postgres.ColumnBool

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type PinsTable struct {
	pinsTable

	EXCLUDED pinsTable
}

// AS creates new PinsTable with assigned alias
func (a PinsTable) AS(alias string) *PinsTable {
	return newPinsTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new PinsTable with assigned schema name
func (a PinsTable) FromSchema(schemaName string) *PinsTable {
	return newPinsTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new PinsTable with assigned table prefix
func (a PinsTable) WithPrefix(prefix string) *PinsTable {
	return newPinsTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new PinsTable with assigned table suffix
func (a PinsTable) WithSuffix(suffix string) *PinsTable {
	return newPinsTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newPinsTable(schemaName, tableName, alias string) *PinsTable {
	return &PinsTable{
		pinsTable: newPinsTableImpl(schemaName, tableName, alias),
		EXCLUDED:  newPinsTableImpl("", "excluded", ""),
	}
}

func newPinsTableImpl(schemaName, tableName, alias string) pinsTable {
	var (
		UserIDColumn   = postgres.IntegerColumn("user_id")
		MsgIDColumn    = postgres.IntegerColumn("msg_id")
		IsActiveColumn = postgres.BoolColumn("is_active")
		allColumns     = postgres.ColumnList{UserIDColumn, MsgIDColumn, IsActiveColumn}
		mutableColumns = postgres.ColumnList{IsActiveColumn}
	)

	return pinsTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		UserID:   UserIDColumn,
		MsgID:    MsgIDColumn,
		IsActive: IsActiveColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}