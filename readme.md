# myHostel 📝💻

## Overview 🎯

myHostel HTTP API is a dynamic and efficient hostel management system developed using the GoFr framework.

Gone are the days of signing attendance sheets and dealing with piles of paperwork. With our API, everything from checking attendance to managing hostel resources is just a click away.It is designed to modernize and simplify hostel administration by using the power of digital solutions.

## Key Features 🔐

- CRUD operations for student management.
- Attendance tracking for students.
- Real-time message broadcasting system using WebSockets,  keeping students and staff connected and informed.
- Reduces reliance on paper, contributing to environmentally friendly practices.

## Project Structure ⛓
```
/hostel-management-system
|-- /api
|   |-- /students
|   |   |-- add_student.go
|   |   |-- get_students.go
|   |   |-- get_student.go
|   |   |-- update_student.go
|   |   |-- delete_student.go
|   |-- /attendances
|       |-- record_attendance.go
|       |-- get_attendance.go
|       |-- update_attendance.go
|       |-- delete_attendance.go
|-- /models
|   |-- student.go
|   |-- attendance.go
|   |-- message.go
|-- main.go
```

## REST API Endpoints ▶

##### Add a New Student
- **Endpoint**: `/student`
- **Method**: POST


##### Get All Students
- **Endpoint**: `/students`
- **Method**: GET

##### Get a Single Student
- **Endpoint**: `/student/{studentID}`
- **Method**: GET

##### Update a Student
- **Endpoint**: `/student/{studentID}`
- **Method**: PUT

##### Delete a Student
- **Endpoint**: `/student/{studentID}`
- **Method**: DELETE

##### Record Attendance
- **Endpoint**: `/attendances`
- **Method**: POST

##### Get Attendance for a Student
- **Endpoint**: `/attendances/{studentID}`
- **Method**: GET
- **Description**: Retrieve all attendance records for a specific student.

##### Update an Attendance Record
- **Endpoint**: `/attendances/{recordID}`
- **Method**: PUT

##### Delete an Attendance Record
- **Endpoint**: `/attendances/{recordID}`
- **Method**: DELETE


#### WebSocket Communication

##### Real-Time Chat
- **Endpoint**: `/websocket`
- **Method**: WebSocket Connect
- **Description**: Connect to the WebSocket to engage in real-time chat. Send a JSON message with `sender`, `receiver`, and `content` fields, and the message will be broadcasted to all connected clients.



   **Warden sending message**

   ![warden](/images/warden.png)

   **Student-1 receiving message**

   ![student-1](/images/student-1.png)


   **Student-2 receiving message**

   ![student-2](/images/student-2.png)

## Testing 

[<img src="https://run.pstmn.io/button.svg" alt="Run In Postman" style="width: 128px; height: 32px;">](https://documenter.getpostman.com/view/21763260/2s9YknANBf)

## Uml Diagram

![My Image](/images/final.png)

## SQL Tables
1. **Student Table**:

![studtable](/images/studentstable.png)

2. **Attendance Table**:

![attable](/images/attendancetable.png)

## Installation and Setup 🛠


1. Ensure you have go installed 

    [Go](https://golang.org/dl/) (version 1.x or higher)


2. Clone the repository using:

````
   git clone https://github.com/aynashi/myHostel.git
````


3. **Install GoFr Modules**:```go get gofr.dev```



4. **Install Other Dependencies**:```go mod tidy```



5. You can run the mysql server and create a database locally using the following docker command:

``
docker run --name gofr-mysql -e MYSQL_ROOT_PASSWORD=root123 -e MYSQL_DATABASE=test_db -p 3306:3306 -d mysql:8.0.30
``



6. Access test_db database and create table students with columns id and name

``
docker exec -it gofr-mysql mysql -uroot -proot123 test_db -e {sql code }
``



7. GoFr reads configuration via environment variables(.env) in configuration folder.


8. **Run the api**
````
go run main.go
````
This would start the server at 9000 port.


## References 🖨

1.  [Check gofr documetation here](https://gofr.dev/)
2.   [Check go documetation here](https://go.dev/doc) 

