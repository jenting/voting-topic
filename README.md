# voting-topic

A website that allows user to create new topic and voting (upvote or downvote).

## Language

* F/E: HTML + jQuery

* B/E: Golang

## Setup

* Download the project

```sh
go get github.com/hsiaoairplane/voting-topic
```

* Compile and run the project

```sh
./run.sh
```

## RESTful APIs

* CRUD

|    Method   |     URL     | Description |
|-------------|-------------|-------------|
| GET | <http://localhost:8080/toptopic> | Query top 20 topic informations. |
| GET | <http://localhost:8080/topic?uid={uid}> | Query topic information with specific uid. |
| POST | <http://localhost:8080/topic> | Create topic with JSON body. |
| PUT | <http://localhost:8080/topic/upvote> | Update upvote by 1 with specific uid in JSON body. |
| PUT | <http://localhost:8080/topic/downvote> | Update downvote by 1 with specific uid in JSON body. |

* HTTP POST/PUT JSON body

|    Field     |   Type(Length)    |    Description  |
|--------------|-------------------|-----------------|
|     uid      |  Version 4 UUID   |       UUID      |
|     name     |  String(255)      |    Topic name   |
|    upvote    |  Unsigned Integer |   Upvote count  |
|   downvote   |  Unsigned Integer |  Downvote count |

## TODO

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
