package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type InsertTipoDescuento_20210115_105430 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertTipoDescuento_20210115_105430{}
	m.Created = "20210115_105430"

	migration.Register("InsertTipoDescuento_20210115_105430", m)
}

// Run the migrations
func (m *InsertTipoDescuento_20210115_105430) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	file, err := ioutil.ReadFile("../scripts/20210115_105430_insert_tipo_descuento.up.sql")

	if err != nil {
		// handle error
		fmt.Println(err)
	}

	requests := strings.Split(string(file), ";\n")

	for _, request := range requests {
		fmt.Println(request)
		m.SQL(request)
		// do whatever you need with result and error
	}

}

// Reverse the migrations
func (m *InsertTipoDescuento_20210115_105430) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	file, err := ioutil.ReadFile("../scripts/20210115_105430_insert_tipo_descuento.down.sql")

	if err != nil {
		// handle error
		fmt.Println(err)
	}

	requests := strings.Split(string(file), ";\n")

	for _, request := range requests {
		fmt.Println(request)
		m.SQL(request)
		// do whatever you need with result and error
	}
}
