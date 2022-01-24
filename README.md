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
### application/x-www-form-urlencoded
```zsh
curl -X POST http://localhost:8080/messages/create/process -d 'name=Gloin&message=Hello, world!'
```

```zsh
% curl -X POST \
http://localhost:8080/messages/createjson/process \
-H "Content-Type: application/json" -H "Connection: close" \
-d '[{"name":"one", "msg":"..."},{"name":"two","msg":"2"}]' -v
```
or in browser:

```
http://localhost:8080/messages/create/process?name=Bofur&message=To+the+weapons
```

### application/json
```zsh
curl -X POST http://localhost:8080/messages/createjson/process -H "Content-Type: application/json" -d '{"name":"Bombur", "msg":"Boop.Beep."}'
```

```zsh
curl -X POST \
http://localhost:8080/messages/createjson/process \
-H "Content-Type: application/json" \
-d '{"name":"Oin", "msg":"WARGH!"}'
```

## Structure
```
.
├── app
│  ├── assets
│  │  └── stylesheets
│  │     └── style.css
│  ├── controllers
│  │  ├── messages_controller.go
│  │  └── messages_json_controller.go
│  ├── helpers
│  │  └── parse_json.go
│  ├── models
│  │  ├── message.go
│  │  ├── message_json.go
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
├── go.mod
├── go.sum
├── main.go
└── README.md
```

---

```
.
├── app/
│  ├── assets/
│  ├── controllers/
│  ├── helpers/
│  ├── models/
│  └── views/
├── db/
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

- [X] Create JSON
- [ ] Read JSON
- [ ] Update JSON
- [ ] Delete JSON

[^1]: Model View Controller.
[^2]: Create Read Update Delete.
[^3]: GitHub Cli.
