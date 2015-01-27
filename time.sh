
function timeit() {
    echo $1 >> times.txt
    { time -p ./merge_sort -n $1 --output 0 ; } 2>> times.txt
}

timeit 10
timeit 100
timeit 1000
timeit 10000
timeit 100000
timeit 1000000

