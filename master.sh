PORT=9123
SEP="------------------------"

_run_test () {
    echo "running test"
    echo "  type: $1"
    echo "  data: $2"
    printf "  --\n  response:\n    "

    if [[ "$1" == "get" ]]; then
        curl -N -X GET localhost:$PORT/orders/$2
    else
        curl -N localhost:$PORT/orders -d $2
    fi

    printf "\n\n"
}

_start_server () {
    echo "launching server"

    go run main.go utils.go external.go&

    echo "waiting for server to launch"
    until [[ 0 -eq $(curl -v --silent http://localhost:$PORT 2>&1 | grep "Connection refused" | wc -l ) ]]; do
        printf '.'
        sleep 1
    done

    echo
    echo "server running in the background at http://localhost:$PORT"
}

_shutdown_server () {
    # kill any previous go proc
    echo "killing any service bound to port $PORT"
    fuser -s -k $PORT/tcp
}

restart_server () {
    _shutdown_server
    _start_server
}

run_tests () {
    restart_server
    printf "running all tests\n$SEP\n\n"

    _run_test post '{"data":"my_Order_Data_For_The_XML_Server"}'
    _run_test post '{"data":"some_other_data"}'
    _run_test get 6f6d7657-3a0e-4dbe-9e2a-1a264e794347
    _run_test get aeffb38f-a1a0-48e7-b7a8-2621a2678534

    printf "%s\ntests complete" $SEP
    printf "\nserver is still running at http://localhost:$PORT\n"
}

install () {
    go get -u github.com/gorilla/mux
}

USAGE="
Usage:

    install         - download go dependencies
    restart_server  - run a server; stop any running server first
    run_tests       - run all of the tests
"

echo "$USAGE"
