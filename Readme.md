# **APP WITH GRAPHQL AND GO**

## 📋 Table of Contents

1. [📖 About the Project](#-about-the-project)
2. [🛠️ Tools Used](#%EF%B8%8F-tools-used)
3. [📋 Prerequisites](#-prerequisites)
4. [🚀 Project Usage](#-project-usage)
5. [📜 Preview](#-preview)

---

## 📖 About the Project

This project is a simple **Hello World** example that demonstrates the usage of **GraphQL** in a Go backend application. The project sets up a GraphQL API server, processes queries, and returns data.

## 🛠️ Tools Used

- **Go**: The programming language used to build the GraphQL API.
- **GraphQL-go**: A Go package for implementing GraphQL queries.
- **GraphiQL**: A web-based interface that allows you to write and test GraphQL queries.

## 📋 Prerequisites

Before you begin, make sure you have:

1. **Go** installed in your machine. You can download it from [here](https://golang.org/dl/).

2. Basic knowledge of **GraphQL** and its principles.

---

## 🚀 Project Usage

### 1. Clone the Repository
First, clone the repository to your local machine:
```bash
git clone https://github.com/EAllaucaD/graphql_app.git
```

### 2. Commands

You need to run the following commands
```
    go mod init graphql
```

```
    go get github.com/graphql-go/graphql
```

```
    go get github.com/graphql-go/handler
```


### 3. Run the server

Once the packages are installed, a go.mod and another go.sum file will be created. Then navigate to the project directory and run the Go application with the following command:

```
    go run src/graphql.go
```

Now you can interact with the GraphQL server using GraphiQL by navigating to http://localhost:8081/graphql in your browser
You can test a simple Hello World query like this:
```
{
  hello
}
```

## 🎨 Preview
