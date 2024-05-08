# beego-example
Beego is an open-source web framework for Go that follows the MVC pattern. It includes built-in support for session management, caching, validation, and internationalization, as well as support for various databases. Beego also provides a command-line tool for generating code scaffolding and running tasks.

`go get -u github.com/astaxie/beego go get -u github.com/beego/bee/v2`

To check whether the `bee` CLI is successfully installed in your system, check the version using: `bee version` 

To create a new project, give this command: `bee new beegoproject` 

This command will create a new project with all the files required for a Beego project and the module name will be beegoproject.
Now you can enter into the project directory and give the `go mod tidy` command to import the required packages.

To start the Beego application, use below command:
`bee run`

This will start the Beego server and your application will be accessible at http://localhost:8080.

Use docker-compose to run a MySQL instance. To start the MySQL instance, use below command:
`docker-compose up -d`