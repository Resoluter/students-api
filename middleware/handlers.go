package middleware

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/Resoluter/students-api/models"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type response struct {
	ID      int64  `json:"id"`
	Message string `json:"message"`
}

const (
	PORT     = 5432
	USER     = "postgres"
	PASSWORD = "xsiroj1999"
	DBNAME   = "postgres"
)

func createConnection() *sql.DB {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	dbInfo := fmt.Sprintf("port=%s, user=%s, password=%s, dbname=%s", PORT, USER, PASSWORD, DBNAME)

	db, err := sql.Open("postgres", dbInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()

	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

	return db

}

func CreateStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Origin", "POST")
	w.Header().Set("Access-Control-Allow-Origin", "Content-Type")

	var student models.Student

	err := json.NewDecoder(r.Body).Decode(&student)

	if err != nil {
		log.Fatalf("Unable to decode the request body. %v", err)
	}

	insertID := insertStudent(student)

	res := response{
		ID:      insertID,
		Message: "Student created successfully",
	}

	json.NewEncoder(w).Encode(res)
}

func GetStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Unable to convert the string into int. %v", err)
	}

	student, err := getStudent(int64(id))

	if err != nil {
		log.Fatalf("Unable to get student. %v", err)
	}

	json.NewEncoder(w).Encode(student)
}

func GetAllStudents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	students, err := getAllStudents()

	if err != nil {
		log.Fatalf("Unable to get all students. %v", err)
	}

	json.NewEncoder(w).Encode(students)
}

func UpdateStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Origin", "POST")
	w.Header().Set("Access-Control-Allow-Origin", "Content-Type")

	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Unable to conver the string into int. %v", err)
	}

	var student models.Student

	err = json.NewDecoder(r.Body).Decode(&student)

	if err != nil {
		log.Fatalf("Unable to decode the request Body. %v", err)
	}

	updatedRows := updateStudent(int64(id), student)

	msg := fmt.Sprintf("Student updated successfully. Total rows/record affected", updatedRows)

	res := response{
		ID:      int64(id),
		Message: msg,
	}

	json.NewEncoder(w).Encode(res)
}

func DeleteStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Unable to convert the string into int. %v", err)
	}

	deletedRows := deleteStudent(int64(id))

	msg := fmt.Sprintf("Student deleted successfully. Total rows/record affected %v", deletedRows)

	res := response{
		ID:      int64(id),
		Message: msg,
	}

	json.NewEncoder(w).Encode(res)
}

func insertStudent(student models.Student) int64 {

	db := createConnection()

	defer db.Close()

	sqlStatement := `INSERT INTO student (first_name, last_name, email) VALUES ($1, $2, $3) RETURNING userid`

	var id int64

	err := db.QueryRow(sqlStatement, student.FirstName, student.LastName, student.Email).Scan(&id)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	fmt.Printf("Inserted a single record %v", id)

	return id
}

func getStudent(id int64) (models.Student, error) {

	db := createConnection()

	defer db.Close()

	var student models.Student

	sqlStatemnt := `SELECT * FROM students WHERE userid=$1`

	row := db.QueryRow(sqlStatemnt, id)

	err := row.Scan(&student.ID, &student.FirstName, &student.LastName, &student.Email)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned")
		return student, nil
	case nil:
		return student, nil
	default:
		log.Fatalf("Unable to scan the row. %v", err)
	}

	return student, err
}

func getAllStudents() ([]models.Student, error) {

	db := createConnection()

	defer db.Close()

	var students []models.Student

	sqlStatement := `SELECT * FROM students`

	rows, err := db.Query(sqlStatement)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	defer rows.Close()

	for rows.Next() {
		var student models.Student

		err = rows.Scan(&student.ID, &student.FirstName, &student.LastName, &student.Email)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		students = append(students, student)
	}

	return students, err
}

func updateStudent(id int64, student models.Student) int64 {

	db := createConnection()

	defer db.Close()

	sqlStatement := `UPDATE students SET first_name=$2, last_name=$3, email=$4 WHERE userid=$1`

	res, err := db.Exec(sqlStatement, id, student.FirstName, student.LastName, student.Email)
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v", err)
	}
	fmt.Printf("Total rows/record affected %v", rowsAffected)

	return rowsAffected
}

func deleteStudent(id int64) int64 {

	db := createConnection()

	defer db.Close()

	sqlStatement := `DELETE FROM students WHERE userid=$1`

	res, err := db.Exec(sqlStatement, id)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v", err)
	}

	fmt.Printf("Total rows/record affected %v", rowsAffected)
	return rowsAffected
}
