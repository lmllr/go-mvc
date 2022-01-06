# GO -- MVC[^1] -- CRUD[^2]
## DB
```zsh
createdb go-mvc
```

```zsh
psql -d go-mvc
```

```zsh
psql -f db/setup.sql -d go-mvc
```

## RUN
```zsh
export PORT=8080
```

```zsh
go run main.go
```

## POST
```zsh
curl -X POST http://localhost:8080/messages/create/process -d 'name=Gloin&message=Hello, world!'
```

or in browser:

```
http://localhost:8080/messages/create/process?name=gloin&message=hello+world
```

## TODO
- [X] Create
- [X] Read
- [ ] Update
- [ ] Delete
- [ ] ~~Frontend~~
- [ ] More data for title, ...

[^1]: Model View Controller.
[^2]: Create Read Update Delete.

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
│  │  └── message.go
│  └── views
│     ├── create.gohtml
│     ├── footer_component.gohtml
│     ├── header_component.gohtml
│     ├── index.gohtml
│     ├── messages.gohtml
│     └── show_message.gohtml
├── db
│  └── db.go
├── go.mod
├── go.sum
├── main.go
├── README.md
├── setup.sql
└── TODO.md
```
