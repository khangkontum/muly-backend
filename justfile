set positional-arguments 

mc *args='':
    migrate create -seq -ext=.sql -dir=./migrations $1 

mup *args='':
    migrate -path=./migrations \
    -database='postgres://muly:mulythegoat@localhost:5432/postgres?sslmode=disable' up

mdown *args='':
    migrate -path=./migrations -database=muly down
