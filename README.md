# GO -- MVC[^1] -- CRUD[^2]
This is a simple example of how to structure a go web application using the MVC[^1] pattern.
## CLONE
ssh:
```zsh
git clone git@github.com:lmllr/go-mvc.git
```
gh[^3]:
```zsh
gh repo clone lmllr/go-mvc
```

## DB (PostgreSQL)
```zsh
createdb go-mvc
```

```zsh
psql -d go-mvc
```

```zsh
psql -f db/setup.sql -d go-mvc
```

## CREATE USER
```zsh
curl -X POST 'http://localhost:8080/user/create/process?name=luke&email=luke@skywalker.com&pw=1234'
```

## RUN
```zsh
export PORT=8080
```

```zsh
go run main.go
```

```zsh
curl http://localhost:8080/
```

```zsh
curl http://localhost:8080/messages
```

or visit in your browser: http://localhost:8080

## POST
```zsh
curl -X POST http://localhost:8080/messages/create/process -d 'name=Gloin&message=Hello, world!'
```

or in browser:

```
http://localhost:8080/messages/create/process?name=gloin&message=hello+world
```

## PUT
```zsh
curl -X PUT -d "id=57&name=Oin&message=Er ist wieder DA\!" http://localhost:8080/messages/update/process
```

## DELETE
```zsh
curl -X DELETE http://localhost:8080/messages/delete/process/54
```

```zsh
curl -X DELETE http://localhost:8080/messages/delete/process\?id=55
```

## Structure
```
.
├── app
│  ├── assets
│  │  └── stylesheets
│  │     └── style.css
│  ├── controllers
│  │  └── messages_controller.go
│  ├── models
│  │  ├── message.go
│  │  ├── page_data.go
│  │  └── raw_data.go
│  └── views
│     ├── create_message.gohtml
│     ├── footer_component.gohtml
│     ├── header_component.gohtml
│     ├── index.gohtml
│     ├── show_all_messages.gohtml
│     ├── show_message.gohtml
│     └── update_message.gohtml
├── db
│  ├── db.go
│  └── setup.sql
├── .gitignore
├── go.mod
├── go.sum
├── main.go
└── README.md
```

## TODO
- [X] Create
- [X] Read
- [X] Update
- [X] Delete
- [ ] ~~Frontend~~
- [X] More data for title, ...
- [X] Make some comments!!!

[^1]: Model View Controller.
[^2]: Create Read Update Delete.
[^3]: GitHub Cli.
