// Copyright (C) 2019 Yasuhiro Matsumoto <mattn.jp@gmail.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

//go:build !cgo
// +build !cgo

package sqlite3

import (
	"database/sql"
	"database/sql/driver"
	"errors"
)

var errorMsg = errors.New("Binary was compiled with 'CGO_ENABLED=0', go-sqlite3 requires cgo to work. This is a stub")
var driverName = "sqlite3"

func init() {
	if driverName != "" && !isRegistered(driverName) {
		sql.Register(driverName, &SQLiteDriver{})
	}
}

func isRegistered(name string) bool {
	for _, v := range sql.Drivers() {
		if v == name {
			return true
		}
	}
	return false
}

type (
	SQLiteDriver struct {
		Extensions  []string
		ConnectHook func(*SQLiteConn) error
	}
	SQLiteConn struct{}
)

func (SQLiteDriver) Open(s string) (driver.Conn, error)                        { return nil, errorMsg }
func (c *SQLiteConn) RegisterAggregator(string, interface{}, bool) error       { return errorMsg }
func (c *SQLiteConn) RegisterAuthorizer(func(int, string, string, string) int) {}
func (c *SQLiteConn) RegisterCollation(string, func(string, string) int) error { return errorMsg }
func (c *SQLiteConn) RegisterCommitHook(func() int)                            {}
func (c *SQLiteConn) RegisterFunc(string, interface{}, bool) error             { return errorMsg }
func (c *SQLiteConn) RegisterRollbackHook(func())                              {}
func (c *SQLiteConn) RegisterUpdateHook(func(int, string, string, int64))      {}
