set positional-arguments 

@mc *args='':
    migrate create -seq -ext=.sql -dir=./migrations $1 
