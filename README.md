This is a simple web app for creating, updating, deleting tasks. The backend is written in Golang and front end is written using React framework. 

TODOs: 

1. Add persistent storage (SQLite database)
2. GraphQL Integration

### UI Code set up

For the UI, from the project root, run `npx create-react-app ui`. It will create the `ui` folder and the child folders (except `components` folder under `src`). Component code files are written by developer.

### How to run

#### Run the server
Go to the project root directory and run the following in the command line. 

```go
go build
./sample-web-app
```

#### Run the client

Check [client ReadMe](ui/README.md)


### Code Reference Documents
I have mainly referred to   
https://observiq.com/blog/embed-react-in-golang (but there are many differences as well). 

#### Some other useful reference docs for reading: 

1. https://medium.com/womenintechnology/building-a-react-app-from-scratch-a-step-by-step-guide-2a42a4be41fc
2. https://www.allhandsontech.com/programming/golang/web-app-sqlite-go/
3. https://www.allhandsontech.com/programming/golang/how-to-use-sqlite-with-go/
3. https://medium.com/@madhanganesh/golang-react-application-2aaf3bca92b1 - https://github.com/madhanganesh/taskpad/tree/master
4. https://hackteam.io/blog/create-react-project-from-scratch-without-framework/
5. https://www.jetbrains.com/guide/go/tutorials/webapp_go_react_part_two/
6. https://www.dhiwise.com/post/building-scalable-web-applications-with-react-and-golang
7. https://medium.com/@chrischuck35/how-to-build-a-simple-web-app-in-react-graphql-go-e71c79beb1d
8. https://tutorialedge.net/projects/chat-system-in-go-and-react/