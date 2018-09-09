package control

type Emp struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type emps []Emp
