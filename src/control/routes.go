package control

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gocql/gocql"
	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var Session *gocql.Session

func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return router
}

var routes = Routes{
	Route{
		"EmpSave",
		"POST",
		"/empCreate",
		EmpCreate,
	},
	Route{
		"EmpGET",
		"GET",
		"/Get",
		Get,
	},
}

func EmpCreate(w http.ResponseWriter, r *http.Request) {
	log.Println("Inside EmpCreate")
	var emp Emp
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1068487))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &emp); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	SaveEmpToDb(&emp)

}

func Get(w http.ResponseWriter, r *http.Request) {
	var empList []Emp

	m := map[string]interface{}{}
	type AllEmpsResponse struct {
		Emps []Emp `json:"messages"`
	}
	//create query
	query1 := "SELECT name FROM shivapreals.emptable"
	var ses = GetSession()
	defer ses.Close()
	iterable := ses.Query(query1).Iter()
	for iterable.MapScan(m) {
		empList = append(empList, Emp{
			//Id:   m["id"].(gocql.UUID),
			Name: m["name"].(string),
		})
		m = map[string]interface{}{}
	}

	json.NewEncoder(w).Encode(AllEmpsResponse{Emps: empList})
}
func GetSession() *gocql.Session {

	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "shivapreals"
	Session, err := cluster.CreateSession()
	if err != nil {
		fmt.Println("Session NULL")
	}

	return Session
}
func SaveEmpToDb(emp *Emp) {
	var gocqlUuid gocql.UUID
	gocqlUuid = gocql.TimeUUID()

	var ses = GetSession()
	defer ses.Close()
	fmt.Println("cassandra init done")
	// writing data to Cassandra
	if err := ses.Query(`
      INSERT INTO emptable (id, Name) VALUES (?, ?)`,
		gocqlUuid, &emp.Name).Exec(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("yes")
	}
}
