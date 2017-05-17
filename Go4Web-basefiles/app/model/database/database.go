package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" //Driver para uso do MySQL
)

type Database struct {
	*sql.DB
}

var DB Database

type Search map[int]map[string]string

// NewOpen abre uma conexão no banco de dados (utilizada apenas na inicialização do programa)
func NewOpen(dados string) (Database, error) {
	db, err := sql.Open("mysql", dados)
	return Database{db}, err
}

// DB.Get procura os campos requisitados a partir de uma string {"campo", "valor"}
func (DB Database) Get(table string, id []string, campos []string) (resultado Search, err error) {

	var y string
	for index, element := range campos {
		if index > 0 {
			y = y + ", "
		}
		y = y + "`" + element + "`"
	}

	var idQuery string

	for i := 0; i < (len(id) / 3); i = i + 3 {
		if i == 0 {
			idQuery = "`" + id[0] + "` " + id[1] + " '" + id[2] + "'"
		} else {
			idQuery = idQuery + " AND ` " + id[i] + " `" + id[i+1] + "'" + id[i+2] + "'"
		}
	}
	table = "`" + table + "`"
	query := "SELECT " + y + " FROM " + table + " WHERE " + idQuery + ";"

	rows, _ := DB.Query(query)
	cols, _ := rows.Columns()
	row := 0
	r := make(map[int]map[string]string)
	for rows.Next() {
		columns := make([]interface{}, len(cols))
		columnPointers := make([]interface{}, len(cols))
		for i := range columns {
			columnPointers[i] = &columns[i]
		}

		if err := rows.Scan(columnPointers...); err != nil {
			return nil, err
		}

		r[row] = make(map[string]string)
		for i, colName := range cols {
			val := columnPointers[i].(*interface{})
			t := *val
			if t == nil {
				r[row][colName] = ""
			} else {
				r[row][colName] = string(t.([]uint8))
			}

		}
		row++
	}
	resultado = r
	return
}

// DB.Insert insere uma nova linha no banco de dados
func (DB Database) Insert(table string, campos []string, valores []string) (err error) {
	var y string
	for index, element := range campos {
		if index > 0 {
			y = y + ", "
		}
		y = y + element
	}
	var x string
	for index, element := range valores {
		if index > 0 {
			x = x + ", "
		}
		x = x + "'" + element + "'"
	}

	tx, err := DB.Begin()
	if err != nil {
		return
	}

	query := "INSERT INTO " + table + "(" + y + ") VALUES (" + x + ");"

	smt, err := tx.Prepare(query)
	if err != nil {
		return
	}

	defer func() {
		if err == nil {
			tx.Commit()
		} else {
			tx.Rollback()
		}
		smt.Close()
	}()

	_, err = smt.Exec()
	return
}

// DB.Update atualiza uma informação no banco de dados segundo informação passada
func (DB Database) Update(table string, id []string, campos []string, valores []string) (err error) {
	var y string
	for index, element := range campos {
		if index > 0 {
			y = y + ", "
		}
		y = y + "`" + element + "`='" + valores[index] + "'"
	}
	var idQuery string
	for index, element := range id {
		if index > 0 && (len(id)-(index+1) >= 2) {
			idQuery = idQuery + " AND `" + id[index+1] + "`='" + id[index+2] + "'"
		}
		if int(index) < 1 {
			idQuery = "`" + element + "`='" + id[index+1] + "'"
		}
	}

	query := "UPDATE `" + table + "` SET " + y + " WHERE " + idQuery + ";"
	//UPDATE `table1` SET `nome`='teste' WHERE `codigo`='5';
	//`nome`='Mayara'
	tx, err := DB.Begin()
	if err != nil {
		return
	}

	smt, err := tx.Prepare(query)
	if err != nil {
		return
	}

	defer func() {
		if err == nil {
			tx.Commit()
		} else {
			tx.Rollback()
		}
		smt.Close()
	}()

	_, err = smt.Exec()
	return
}

// DB.Delete deleta uma informação do banco de dados
func (DB Database) Delete(table string, campos []string, valores []string) (err error) {
	var y string
	for index, element := range campos {
		if index > 0 {
			y = y + " AND "
		}
		y = y + "`" + element + "`='" + valores[index] + "'"
	}

	tx, err := DB.Begin()
	if err != nil {
		return
	}

	query := "DELETE FROM `" + table + "` WHERE " + y + ";"

	smt, err := tx.Prepare(query)
	if err != nil {
		return
	}

	defer func() {
		if err == nil {
			tx.Commit()
		} else {
			tx.Rollback()
		}
		smt.Close()
	}()

	_, err = smt.Exec()
	return
}
