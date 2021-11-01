# Test-Driven Development

## Learn Go with Tests

I came across this useful Gitbook which educates new Go programmers with Go tests. Some of the tests written in this folder came from there.

https://quii.gitbook.io/learn-go-with-tests/ 

Read the guide for more information!

## Things to Note

All the tests are placed in their individual directory because **each package must be defined in its own directory**.
A package is a component that you can use in more than one program, that you can publish, import, get from an URL, etc. So it makes sense for it to have its own directory as much as a program can have a directory.

Else, the `go test` command will not work because it will not know which package to test.