// Copyright 2024 Willard Lu
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file or at
// https://opensource.org/licenses/MIT.
package flamingos

import (
	"testing"
)

func TestGetPsqlConfig(t *testing.T) {
	// General Tests
	config, err := GetPsqlConfig("testdata/pg_config.toml")
	if err != "" {
		t.Errorf(err)
		// The t.Error() function will continue down the line after execution,
		// so a RETURN is needed.
		return
	}
	correct_config := "host=localhost port=5432 user=starry password=" +
		"belighT928s dbname=shangjiang_hospital sslmode=disable "
	if config != correct_config {
		t.Errorf("expected '"+correct_config+"', got '%s'", config)
	}
}

func TestConnectPsql(t *testing.T) {
	// General Tests
	config, err := GetPsqlConfig("testdata/pg_config.toml")
	if err != "" {
		t.Errorf(err)
		return
	}
	db, err1 := ConnectPsql(config)
	if err1 != "" {
		t.Errorf(err1)
		return
	}
	err = ClosePsql(db)
	if err != "" {
		t.Errorf(err)
	}
}

func TestSelectPsql(t *testing.T) {
	config, err := GetPsqlConfig("testdata/pg_config.toml")
	if err != "" {
		t.Errorf(err)
		return
	}
	db, err1 := ConnectPsql(config)
	if err1 != "" {
		t.Errorf(err1)
		return
	}
	// General Tests
	query := "select title from test"
	rows, err := PsqlSelect(db, query)
	if err != "" {
		t.Errorf(err)
		return
	}
	var title1 string
	var title2 string
	rows.Next()
	rows.Scan(&title1)
	rows.Next()
	rows.Scan(&title2)
	if title1 != "root" {
		err = ClosePsql(db)
		if err != "" {
			t.Errorf(err)
			return
		}
		t.Errorf("expected 'root', got '%s'", title1)
		return
	}
	if title2 != "leaf" {
		err = ClosePsql(db)
		if err != "" {
			t.Errorf(err)
			return
		}
		t.Errorf("expected 'leaf', got '%s'", title2)
		return
	}
	// Error Tests
	query = "select title from test1"
	_, err = PsqlSelect(db, query)
	if err == "" {
		err = ClosePsql(db)
		if err != "" {
			t.Errorf(err)
			return
		}
		t.Errorf("expected error, got no error")
		return
	}
	err = ClosePsql(db)
	if err != "" {
		t.Errorf(err)
	}
}

func TestPsqlExec(t *testing.T) {
	config, err := GetPsqlConfig("testdata/pg_config.toml")
	if err != "" {
		t.Errorf(err)
		return
	}
	db, err1 := ConnectPsql(config)
	if err1 != "" {
		t.Errorf(err1)
		return
	}
	// General Tests
	query := "insert into test (title) values ('test')"
	err = PsqlExec(db, query)
	if err != "" {
		err = ClosePsql(db)
		if err != "" {
			t.Errorf(err)
			return
		}
		t.Errorf(err)
		return
	}
	query = "update test set title = 'test1' where title = 'test'"
	err = PsqlExec(db, query)
	if err != "" {
		err = ClosePsql(db)
		if err != "" {
			t.Errorf(err)
			return
		}
		t.Errorf(err)
		return
	}
	query = "delete from test where title = 'test1'"
	err = PsqlExec(db, query)
	if err != "" {
		err = ClosePsql(db)
		if err != "" {
			t.Errorf(err)
			return
		}
		t.Errorf(err)
		return
	}
	// Error Tests
	query = "insert into test (title1) values ('test')"
	err = PsqlExec(db, query)
	if err == "" {
		err = ClosePsql(db)
		if err != "" {
			t.Errorf(err)
			return
		}
		t.Errorf("expected error, got no error")
		return
	}
	err = ClosePsql(db)
	if err != "" {
		t.Errorf(err)
	}
}
