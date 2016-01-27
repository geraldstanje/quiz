# quiz


Q: Given a list of words like https://github.com/NodePrime/quiz/blob/master/word.list find the longest compound-word in the list, which is also a concatenation of other sub-words that exist in the list. The program should allow the user to input different data. The finished solution shouldn't take more than one hour. Any programming language can be used, but Go is preferred.


Fork this repo, add your solution and documentation on how to compile and run your solution, and then issue a Pull Request. 

Obviously, we are looking for a fresh solution, not based on others' code.

##Set GOPATH
Go the the top directory and set the GOPATH using:
```
$ export GOPATH="$GOPATH":"$PWD"
```

##Build and Run the App
Flags:<br>
  * -f filename
  * -b enable benchmarking

```
$ go build main.go
$ ./main -f words.list -b
Runtime: 169 ms
Longest compound word: antidisestablishmentarianisms
```

##Test the App
```
$ cd src/lcwsolver
$ go test
PASS
ok    lcwsolver 0.006s
```
