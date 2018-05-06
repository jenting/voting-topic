# voting-topic

A website that allows user to create new topic and voting (upvote or downvote).

**Language:**

* F/E: Vue.js

* B/E: Golang

## RESTful APIs

* CRUD

|    Method   |     URL     | Description |
|-------------|-------------|-------------|
| GET | <http://localhost/toptopic> | Query top 20 topic informations. |
| GET | <http://localhost/topic?name={name}> | Query topic information with specific topic name. |
| POST | <http://localhost/topic> | Create topic with JSON body. |
| PUT | <http://localhost/topic> | Update topic information with JSON body. |

* HTTP POST/PUT JSON body

|    Field     | Type(Length) |  Description |
|--------------|--------------|--------------|
|     Name     |  String(255) |   Topic name |
|    Upvote    |  Unsigned Integer | Upvote count |
|   Downvote   |  Unsigned Integer | Downvote count |

## TODO

* [ ] Add frontend page
* [ ] Support [prometheus](https://prometheus.io) metrics API

## Godep

* Add all dependency `godep save ./...`

* Restore dependency in vendor folderto the $GOPATH `godep restore`

## Limitations

* Topic should not exceed 255 characters.

* Allow user to upvote or downvote the same topic multiple times.

* Homepage lists top 20 topics (sorted by upvotes, descending)

* Keeps the topics in-memory data cache

* Write test cases

* Deploy to Heroku

## Notes

* Do not do AAA

* Do not check duplicate votes

* No needs in-disk database
