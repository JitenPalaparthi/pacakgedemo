### go mod : to create go modules
Go did not have any concept called dependency management. Old versions we used to use go deps

There are 3 kinds of packages in go

1. go standard packages. GOROOT contains go standard packages.
2. user defined packages
3. third party packages.

step-1 : any project you create . initialise that project with a module name by using 
```go mod init demo```
upon giving the name , it creates go.mod file. That file contains all the dependencies required for the project.
If any thirdparty dependencies exists use ```go mod tidy``` in the root directory where go.mod file exists.

### Rules for package creation
1. each package must have a directory
2. the go file can be any name
3. the package name inside go file name must be the directory name
4. go does not have any access specifiers
5. any field or function, method or type that starts with Uppercase is public
6. any field or function, method or type that starts with lowerCase is private
7. a package directory can have any number of files but the name of the pacakge inside all the files must be the directory name which is package name
8. package can have sub directories. The name of the package will be the name of the subdirectory. same way how to use math/rand ; rand.InitN(100)
