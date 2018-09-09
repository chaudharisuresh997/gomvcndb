# gomvcndb
Minimal example of golang with gorilla mux and cassandra as a database

standard template for any starter project
Step1:
Go to GoTemplate directory
run command to set GOPATH

create cassandra keyspace and table
//echo "CREATE KEYSPACE shivapreals WITH replication = {'class': 'SimpleStrategy', 'replication_factor' : 1};" | cqlsh --cqlversion 3.4.2
	//create table shivapreals.emp (id UUID,Name text PRIMARY KEY(id));


export GOPATH=$(pwd)


go build


go run main.go


