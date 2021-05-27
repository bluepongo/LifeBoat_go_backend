package models

import (
	"database/sql"
	"github.com/bluepongo/lifeboat_go_backend/drivers"
	"log"
)

// models package's db obj
// all db operation should be done in models pkg
// so db is a pkg inner var
var db *sql.DB = drivers.MysqlDb

type Role struct {
	ID       int    `json:"id" form:"id" primaryKey:"true"`
	RoleName string `json:"role_name" form:"role_name" binding:"required"`
	HP       int    `json:"hp" form:"hp" binding:"required"`
	SuvScore int    `json:"suv_score" form:"suv_score" binding:"required"`
	Seat     int    `json:"seat" form:"seat" binding:"required"`
	Describe string `json:"des" form:"des" binding:"required"`
}

// Get role by id
func (model *Role) GetRoleByName(name string) (role Role, err error) {
	// find the role by id
	sql := `
		select id, role_name, hp, suv_score, seat, des
		from role_info
		where role_name = ?`
	err = db.QueryRow(sql, name).Scan(
		&role.ID,
		&role.RoleName,
		&role.HP,
		&role.SuvScore,
		&role.Seat,
		&role.Describe)
	if err != nil {
		log.Println(err.Error())
		return
	}
	return
}

// Get all roles
func (model *Role) GetAllRoles() (roles []Role, err error) {
	sql := `
		select id, role_name, hp, suv_score, seat, des
		from role_info`
	roles = make([]Role, 0)
	rows, err := db.Query(sql)
	defer rows.Close()
	if err != nil {
		log.Println(err.Error())
	}
	var role Role
	for rows.Next() {
		rows.Scan(
			&role.ID,
			&role.RoleName,
			&role.HP,
			&role.SuvScore,
			&role.Seat,
			&role.Describe)
		roles = append(roles, role)
	}
	if err = rows.Err(); err != nil {
		log.Println(err.Error())
	}
	return
}

// Add a role
func (model *Role) AddRole() (id int64, err error) {
	sql := `
		insert into role_info(role_name, hp, suv_score, seat, des)
		values(?, ?, ?, ?, ?)`
	result, err := db.Exec(sql,
		model.RoleName,
		model.HP,
		model.SuvScore,
		model.Seat,
		model.Describe,
	)
	if err != nil {
		log.Println(err.Error())
		return
	}
	id, err = result.LastInsertId()
	return
}

// Update the role info by id
func (model *Role) UpdateRoleByID(id int) (afr int64, err error) {
	sql := `
		update role_info
		set role_name = ?, hp = ?, suv_score = ?, seat = ?, des = ?
		where id = ?`
	result, err := db.Exec(sql,
		model.RoleName,
		model.HP,
		model.SuvScore,
		model.Seat,
		model.Describe,
		id,
	)
	if err != nil {
		log.Println(err.Error())
		return
	}
	afr, err = result.RowsAffected()
	return
}

// Delete a role by id
func (model *Role) DeleteRoleByID(id int) (afr int64, err error) {
	sql := `
		delete from role_info
		where id = ?`
	result, err := db.Exec(sql, id)
	if err != nil {
		log.Println(err.Error())
		return
	}
	afr, err = result.RowsAffected()
	return
}
