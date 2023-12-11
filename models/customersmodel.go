package models

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/cahiman99/crud_golang/config"
	"github.com/cahiman99/crud_golang/entities"
)

type CustomersModel struct {
	conn *sql.DB
}

func NewCustomersModel() *CustomersModel {
	conn, err := config.DBConnection()
	if err != nil {
		panic(err)
	}

	return &CustomersModel{
		conn: conn,
	}
}

func (p *CustomersModel) FindAll() ([]entities.Customers, error) {

	rows, err := p.conn.Query("select * from customers")
	if err != nil {
		return []entities.Customers{}, err
	}
	defer rows.Close()

	var dataCustomers []entities.Customers
	for rows.Next() {
		var customers entities.Customers
		rows.Scan(&customers.Id,
			&customers.Name,
			&customers.NIK,
			&customers.JenisKelamin,
			&customers.Tempat_lahir,
			&customers.Tanggal_lahir,
			&customers.Alamat,
			&customers.No_hp)

		if customers.JenisKelamin == "1" {
			customers.JenisKelamin = "Laki-laki"
		} else {
			customers.JenisKelamin = "Perempuan"
		}
		// 2006-01-02 => yyyy-mm-dd
		tgl_lahir, _ := time.Parse("2006-01-02", customers.Tanggal_lahir)
		// 02-01-2006 => dd-mm-yyyy
		customers.Tanggal_lahir = tgl_lahir.Format("02-01-2006")

		dataCustomers = append(dataCustomers, customers)
	}

	return dataCustomers, nil

}

func (p *CustomersModel) Create(customers entities.Customers) bool {

	result, err := p.conn.Exec("insert into customers (nama_lengkap, nik, jenis_kelamin, tempat_lahir, tanggal_lahir, alamat, no_hp) values(?,?,?,?,?,?,?)",
		customers.Name, customers.NIK, customers.JenisKelamin, customers.Tempat_lahir, customers.Tanggal_lahir, customers.Alamat, customers.No_hp)

	if err != nil {
		fmt.Println(err)
		return false
	}

	lastInsertId, _ := result.LastInsertId()

	return lastInsertId > 0
}

func (p *CustomersModel) Find(id int64, customers *entities.Customers) error {

	return p.conn.QueryRow("select * from customers where id = ?", id).Scan(
		&customers.Id,
		&customers.Name,
		&customers.NIK,
		&customers.JenisKelamin,
		&customers.Tempat_lahir,
		&customers.Tanggal_lahir,
		&customers.Alamat,
		&customers.No_hp)
}

func (p *CustomersModel) Update(customers entities.Customers) error {

	_, err := p.conn.Exec(
		"update customers set name = ?, nik = ?, jenis_kelamin = ?, tempat_lahir = ?, tanggal_lahir = ?, alamat = ?, no_hp = ? where id = ?",
		customers.Name, customers.NIK, customers.JenisKelamin, customers.Tempat_lahir, customers.Tanggal_lahir, customers.Alamat, customers.No_hp, customers.Id)

	if err != nil {
		return err
	}

	return nil
}

func (p *CustomersModel) Delete(id int64) {
	p.conn.Exec("delete from customers where id = ?", id)
}
