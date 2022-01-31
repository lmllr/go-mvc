# GO -- MVC[^1] -- CRUD[^2]
This is a simple example of how to structure a go web application using the MVC[^1] pattern.

* [Clone the project](#clone-the-project)
* [DB Setup](#setup-db-(postgresql)
* [Quickstart the app](#quickstart-the-app)
* [Make some more HTTP requests](#make-some-more-http-requests)
  * [GET](#get)
  * [POST](#post)
  * [PUT](#put)
  * [DELETE](#delete)
* [Structure](#structure)

## Clone the project
ssh:
```zsh
git clone git@github.com:lmllr/go-mvc.git
```
gh[^3]:
```zsh
gh repo clone lmllr/go-mvc
```

## Setup DB (PostgreSQL)
```zsh
createdb go-mvc
```

```zsh
psql -d go-mvc
```

```zsh
psql -f db/setup.sql -d go-mvc
```

## Quickstart the app
```zsh
export PORT=8080
```

```zsh
go run main.go
```

```zsh
curl http://localhost:8080/
```

or visit in your favourite browser: [http://localhost:8080](http://localhost:8080)

## Make some more HTTP Requests
### GET
#### application/x-www-form-urlencoded
Get all messages:  
```zsh
curl http://localhost:8080/messages
```

Get a single message by ID:  
```zsh
curl 'http://localhost:8080/messages/show?id=1'
```

```zsh
curl http://localhost:8080/messages/show\?id=1
```

#### application/json
Get all messages in JSON ARRAY format:  
```zsh
curl http://localhost:8080/json/messages
```

Get a single message in JSON format by ID:  
```zsh
curl -X GET \
http://localhost:8080/json/messages/showjson \
-H "Content-Type: application/json" \
-d '{"id":1}'
```

Get multiple messages in JSON ARRAY format by ID:  
```zsh
curl -X GET \
http://localhost:8080/json/messages/showjsonarr \
-H "Content-Type: application/json" \
-d '[{"id":3},{"id":4}]'
```

### POST
#### application/x-www-form-urlencoded
Create a single message:  
```zsh
curl -X POST http://localhost:8080/messages/create/process -d 'name=Gandalf&message=Hello, thy little dwarfs!'
```

or in browser:  
```
http://localhost:8080/messages/create/process?name=Frodo&message=To+the+Beutlin
```

#### application/json
Create a single message in JSON format:  
```zsh
curl -X POST http://localhost:8080/json/messages/createjson -H "Content-Type: application/json" -d '{"name":"Aragorn", "msg":"Boop.Beep."}' 
```

Create multiple messages in JSON ARRAY format:  
```zsh
curl -X POST \
http://localhost:8080/json/messages/createjsons \
-H "Content-Type: application/json" \
-H "Connection: close" \
-d '[{"name":"Sauraus", "msg":"EWOIFJNWEIOF"},{"name":"Arwen","msg":"If you want him, come and claim him!"}]'
```

### PUT
#### application/json
Update a single message by ID, in JSON format:  
```zsh
curl -X PUT \
http://localhost:8080/json/messages/updatejson \
-H "Content-Type: application/json" \
-d '{"id":1,"name":"UPDATED TWICE Dwalin","msg":"Okey, STOP IT!"}'
```

Update multiple messages by ID, in JSON ARRAY format:  
```zsh
curl -X PUT \
http://localhost:8080/json/messages/updatejsons \
-H "Content-Type: application/json" \
-d '[{"id":1,"name":"UPDATED Dwalin","msg":"Wooohooo, update me more."},{"id":2,"name":"UPDATED Balin","msg":"grmph..."}]'
```

### DELETE
#### application/x-www-form-urlencoded
>(NO DELETE METHOD)
Delete all messages:  
```zsh
curl http://localhost:8080/messages/deleteall/process
```
#### application/json
Delete a single message by ID, in JSON format:  
```zsh
curl -X DELETE \
http://localhost:8080/json/messages/delete \
-H "Content-Type: application/json" \
-d '{"id":1}'
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
- [X] Read JSON
- [X] Update JSON
  - [X] PUT
  - [ ] PATCH
- [X] Delete JSON

- [ ] Proper error handling

## Footnotes
[^1]: Model View Controller.
[^2]: Create Read Update Delete.
[^3]: GitHub Cli.
