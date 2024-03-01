package models

import (
	"example.com/studentAPI/db"
)

type Student struct {
	ID        int64
	FirstName string
	LastName  string
	EmailId   string
}

func (st *Student) Save() error {
	query := `INSERT INTO students (FirstName, LastName, EmailId) VALUES ($1, $2, $3) RETURNING ID`

	var id int64
	err := db.DB.QueryRow(query, st.FirstName, st.LastName, st.EmailId).Scan(&id)

	if err != nil {
		return err
	}
	st.ID = id

	return nil
}

func GetAllStudents() ([]Student, error) {
	query := `SELECT * FROM students`

	rows, err := db.DB.Query(query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var students []Student

	for rows.Next() {
		var st Student
		err = rows.Scan(&st.ID, &st.FirstName, &st.LastName, &st.EmailId)

		if err != nil {
			return nil, err
		}

		students = append(students, st)
	}

	return students, err
}

func (st Student) Delete() error {
	query := `DELETE FROM students WHERE id = $1`

	_, err := db.DB.Exec(query, st.ID)

	if err != nil {
		return err
	}

	return nil
}

func (st Student) Update() error {
	query := `
	UPDATE students
	SET FirstName = $1, LastName = $2, EmailId = $3
	WHERE id = $4
	`

	_, err := db.DB.Exec(query, st.FirstName, st.LastName, st.EmailId, st.ID)

	if err != nil {
		return err
	}

	return nil
}

func GetStudentByID(id int64) (*Student, error) {
	query := `SELECT * FROM students WHERE id = $1`

	row := db.DB.QueryRow(query, id)

	var student Student
	err := row.Scan(&student.ID, &student.FirstName, &student.LastName, &student.EmailId)
	if err != nil {
		return nil, err
	}

	return &student, nil
}
