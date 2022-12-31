# backend-collection-go-2

### Version: release-1.0 (2022-12-31)

### Summary:
This is Go backend number 2 from my backend collection project. This backend is a REST API for CRUD news data build with Beego framework, MySQL, and Beego ORM. In this release, there's no test files, test files will coming in later release.

### Requirements:
1. go (tested: v1.19.4 windows/amd64)
2. mysql (tested: v8.0.31)
3. bee (tested: v2.0.2)

### Steps to run the backend server:
1. install requirement go and mysql that can be downloaded from official website
2. After install go, install bee with `go install github.com/beego/bee/v2@v2.0.2`
3. clone repository `https://github.com/reyhanfikridz/backend-collection-go-2` at directory `$GOPATH/src/github.com/reyhanfikridz/`
4. at repository root directory, which is `$GOPATH/src/github.com/reyhanfikridz/backend-collection-go-2`:
    1. switch to branch release-1.0 with `git checkout release-1.0`
    2. download required go modules with `go mod download`
    3. copy all downloaded go modules to the repository with `go mod vendor`
    4. create file app.conf at directory conf (`$GOPATH/src/github.com/reyhanfikridz/backend-collection-go-2/conf/app.conf`) with contents:

    ```
    appname = backend-collection-go-2
    httpport = 8080
    runmode = dev
    autorender = false
    copyrequestbody = true
    EnableDocs = true
    sqlconn = dbuser:dbpass@tcp(dbhost:dbport)/dbname (example => someuser:somepassword@tcp(127.0.0.1:3306)/backend_collection_go_2)
    ```

    5. create mysql database with name same as in app.conf file
    6. migrate database using sqlconn data in app.conf with `bee migrate -driver=mysql -conn="someuser:somepassword@tcp(127.0.0.1:3306)/backend_collection_go_2"`
    7. run server with `bee run`

### API collection:
1. Go to https://www.postman.com/reyhanfikri/workspace/backend-collection-go-2/overview
2. Choose `release-1.0` collection

### License:
This project is MIT license, so basically you can use it for personal or commercial use as long as the original LICENSE.md included in your project.
