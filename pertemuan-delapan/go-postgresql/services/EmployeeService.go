package services

import (
	"database/sql"
	"pertemuan-delapan/go-postgresql/models"
)

func CreateEmployee(db *sql.DB, name string, email string, age int, devision string) models.Employee {
	var employee models.Employee

	sql := `
	INSERT INTO "Employee" ("FullName", "Email", "Age", "Devision")
	VALUES($1, $2, $3, $4)
	RETURNING *
	`

	err := db.QueryRow(sql, name, email, age, devision).Scan(
		&employee.Id, &employee.FullName, &employee.Email, &employee.Age, &employee.Devision,
	)

	if err != nil {
		panic(err)
	}

	return employee
}

func GetEmployee(db *sql.DB) []models.Employee {
	var result []models.Employee

	sql := `SELECT * FROM "Employee"`

	rows, err := db.Query(sql)
	defer rows.Close()
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var employee models.Employee

		err = rows.Scan(&employee.Id, &employee.FullName, &employee.Email, &employee.Age, &employee.Devision)
		if err != nil {
			panic(err)
		}

		result = append(result, employee)
	}

	return result
}

func UpdateEmployee(db *sql.DB, id int, name string, email string, age int, devision string) string {
	sql := `
	UPDATE "Employee"
		SET "FullName" = $2, "Email" = $3, "Age" = $4, "Devision" = $5
	WHERE "EmployeeId" = $1
	`

	res, err := db.Exec(sql, id, name, email, age, devision)
	if err != nil {
		panic(err)
	}

	count, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}

	if count < 1 {
		return "Gagal mengubah data!"
	}

	return "Berhasil mengubah data!"
}

func DeleteEmployee(db *sql.DB, id int) string {
	var rows int
	sql := `SELECT COUNT(*) FROM "Employee" WHERE "EmployeeId" = $1`

	err := db.QueryRow(sql, id).Scan(&rows)
	if err != nil {
		panic(err)
	}

	if rows < 1 {
		return "Data not found!"
	}

	sql = `DELETE FROM "Employee" WHERE "EmployeeId" = $1`
	res, err := db.Exec(sql, id)
	if err != nil {
		panic(err)
	}

	_, err = res.RowsAffected()
	if err != nil {
		panic(err)
	}

	return "Data successfully deleted!"
}
