# counter-service-app

## 1. Introduction
This is an counter app written in Go language. It stores, increments, decrements and resets a counter. For storing the counter a MySQL database is used.

## 2. Setting up GO environment
For this project, Go version 1.15 is used. Define the $GOPATH as working directory. It should contain 3 typical folders inside: bin, pkg and src.

### 2.1 Linux & Mac
Go to  _~/.bash_profile_ and add the following line:

export GOPATH=_your working directory_

### 2.2 Windows
Open up your windows start menu and type _environment_. You'll probably see -> _Edit the system environment variables_. This will lead you automatically to the _Advanced_ tab on _System Properties_ window.

Click _Environment Variables_. On top you'll see _User variables for YOUR USER_.

Click _New_. Type _GOPATH_ for variable name and type _your working directory_ for variable value.

### 2.3 Useful Links
- https://golang.org/doc/gopath_code
- https://www.callicoder.com/golang-installation-setup-gopath-workspace/
- https://medium.com/rungo/working-in-go-workspace-3b0576e0534a

## 3. Setting up MySQL Database
For this project, the MySQL 8.0 is used.

### 3.1 Creating a user
On the installation of MySQL, it is asked to create a root password and additionally a user. Remember that username and password.

### 3.2 Creating a database
A database should be created manually to that the application can connect and create a necessary table.

To do that, simply type the following sql command: _create database counterdb_. This creates a database with the name "counterdb".

## 4. Running the application
First, this repository has to be cloned with it's dependencies. For that, open up a command terminal and type _go get -u -v github.com/utr1903/counter-service-app_. The "-u" tag detects and downloads the dependencies and the "-v" tag logs the download processes to the console which makes easier to find errors.

## 5. Application structure
The app is run through the main.go file. Both the database username and password have to be given as command line arguments in order to securely keep them out of the code & repository. To run the app,
* Open up a terminal
* Go to code directory
  * Linux: cd $GOPATH/github.com/utr1903/counter-service-app
  * Windows: cd %GOPATH%\github.com\utr1903\counter-service-app
* type _go install_
* type _go run main.go "username" "password"

### 5.1 main.go
Simply the start of the application. It parses the database username and password. Creates an app object.

### 5.2 app
The folder contains only app.go which creates the server and handles the requests. This is the place where the database connection is established and the API endpoints are defined.

### 5.3 controllers
For this particular application, only one controller is created which is CounterController.go. This extends ControllerBase where primitive methods are defined (parsing request and creating response). The requests land to the CounterController and forwarded to CounterService.

The result coming from CounterService is packed into a meaningful JSON object and returned (see commons).

### 5.4 commons
This package includes the CustomResult model which is used to deliver meaningful JSON responses to user. It contains;
* Success -> true / false
* Model -> the actual result object (in scope of this app, it stands for the value of the counter)
* Code -> string representation of HTTP status code
* Message -> Meaningful message explaining how the request is processed
* Error -> The Go error object

### 5.5 services
This package is where the business logic lies.

#### 5.5.1 GetCounter
Returns the last value of the counter.

#### 5.5.1 IncreaseCounter
Increments the counter by given value.

#### 5.5.1 DecreaseCounter
Decreases the counter by given value.

#### 5.5.1 ResetCounter
Resets the counter to zero.



