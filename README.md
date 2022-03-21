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