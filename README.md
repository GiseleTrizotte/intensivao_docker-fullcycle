# Curso intensivão docker Full Cycle

## Course content source

- Docker - [curso fullcycle](https://www.youtube.com/watch?v=uroAHS9PuNY)
  

## :hammer: Description features of the project

All project content was administered in the course.
In this example with docker, we will upload an application with three services:

- ``nginx``
- ``mysql``
- ``go``

**nginx** will be to create an image that has volumes, ie data.

it will just be an example of a total zero image creation including volume.

**mysql** will be to store a table and data and show the data on the console, in this application we will not expose any port, because the connection will be made by another service in the go language.

**go** will be a service that we will make the connection with mysql, basically it will have only one file the main.go

## :wrench: Running the local application:

Cloning the local application:

```bash
$ git clone
```

Upload all services specified in the docker-compose.yaml manifest file

```bash
$ docker-compose up
```

## ![MySQL](https://img.shields.io/badge/mysql-%2300f.svg?style=for-the-badge&logo=mysql&logoColor=white) MySQL

With the container up we will need to enter the mysql container to log into the database and create a table.
In another tab of the finish run:

```bash
$ docker-compose exec mysql bash
```
Executes a command in an active container.
Now we are inside the mysql container and we need to log in with the credentials that we set when creating the images

```bash
$ mysql -u root -p fullcycle
```

This data is in the file go/main.go
Create a table with the same name that we created in the go/main file

![image](https://user-images.githubusercontent.com/57105956/163225306-8b334786-ba9a-4552-8af6-be82727da84f.png)


```bash
$ create table exemplo (id int, name varchar(255));
```

## ![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white) Go

Now we need to enter the container that will connect to the active mysql service and insert records into the database.
In another tab of the finish run:

```bash
$ cd ..
$ cd go
$ docker-compose exec app bash
```

** app is the name we gave to the golang service
Run the service:

```bash
$ go run main.go
```
When executing this command in the main file there is an insert instruction, and it will insert into the database:

![image](https://user-images.githubusercontent.com/57105956/163225027-f9cca643-6c1c-4bc5-8e41-7a264a44951c.png)

Accessing the mysql terminal we execute the select on the table we created to see the record:

```bash
$ select * from exemplo;
```
Saída:

![image](https://user-images.githubusercontent.com/57105956/163225194-dce4a01a-342f-41a4-a4a7-3123020470e7.png)



