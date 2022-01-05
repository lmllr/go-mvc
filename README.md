# GO -- MVC[^1] -- CRUD[^2]
## POST
```zsh
curl -X POST http://localhost:8080/messages/create/process -d 'name=Gloin&message=Hello, world!'
```

or in browser:

```
http://localhost:8080/messages/create/process?name=Gloin&message=Hello, world!
```

## TODO
- [X] Create
- [ ] Read
- [ ] Update
- [ ] Delete

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
│     ├── footer_component.gohtml
│     ├── header_component.gohtml
│     ├── index.gohtml
│     └── messages.gohtml
├── db
│  └── db.go
├── go.mod
├── go.sum
├── main.go
├── README.md
└── setup.sql
```
