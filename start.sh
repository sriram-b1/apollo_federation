
RunServer () {
    FILE=$PWD/server.go
    if [ -f "$FILE" ]; then
        go run server.go &
    fi 
}


for d in ./* ; 
do 
    if [ -d $d ]
    then
        (cd "$d" &&  RunServer)
    fi
done
cd Gateway
node index
