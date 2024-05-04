# GraphQL Bridge

An idea I've had recenetly, so starting a project on it (in golang!) to learn more fundamental programming tasks.

The goal of this project is to try learn how to parse & compile a "generc data behaviour mapping" language
that constructs a GraphQL Schema & Server that is mapped (in terms of data retrieval) to all our your data sources
already defined and handled by the engine.

```graphql
### redisCache @cache {
###   "key": "music-cache"
### }

### musicDatabase @adapter {
###   "key": "music-database"
### }

### auth0Management @adapter {
###   "key": "auth0-management"
### }

type Query {
	getUsers: [User!]!
	### @adapter(auth0Management) retrieve {
	###   path: "/api/v2/users"
	### }

	getUsersLatest: [User!]!
	### @adapter(auth0Management) retrieve {
	###   path: "/api/v2/users"
	###   query: {
	###     sort: "date_added"
	###     order: "desc"
	###   }
	### }

	getMe: User!
	### @argument:set { userID: ":loggedInUserID" }

	getUserById(
		userID: String! ### @argument:for { key: "userID" }
	): User!

	getArtistById(
		artistID: String! ### @argument:for { key: "artistID" }
	): Artist!

	getAlbumById(
		albumID: String! ### @argument:for { key: "albumID" }
	): Album!
}

type User {
	### @argument declare { key: "userID" }

	### getAuth0User @adapter(Auth0Management) retrieve {
	###   "path": "/api/v2/users/:userID"
	### }

	userID: String!
	### @adapter(auth0User) field {
	###   key: "user_id"
	### }

	name: String!
	### adapter(Auth0Management):field {
	###   key: "name"
	### }

	artists: [Artist!]!
	### adapter(MusicDatabase):manyToMany {
	###   table: "users_artists"
	###   where:
	###     user_id: ":userID"
	###   orderBy: "date_added DESC"
	###   columns:
	###     :artistID: "artist_id"
	###   arguments: ":userID"
	### }
}

type Artist {
	### @argument:required { key: "artistID" }

	### @argument:optional {
	###   key: "userID"
	###   default: ":loggedInUserID"
	### }

	### adapter(MusicDatabase):retrieve {
	###   table: "artists"
	###   columns: {
	###     artistID: ":artistID"
	###   }
	### }

	artistID: String!
	### adapter(MusicDatabase):field {
	###   column: "artist_id"
	### }

	name: String!
	### adapter(MusicDatabase):field {
	###   column: "name_id"
	### }

	dateAdded: Int!
	### adapter(MusicDatabase):sql.oneToOne {
	###   table: "users_artists"
	###   where: {
	###     artist_id: ":artistID"
	###     user_id: ":userID"
	###   }
	### }

	albums: [Album!]!
	### adapter(MusicDatabase):manyToMany {
	###   table: "albums_artists"
	###   where: {
	###     artist_id: ":artistID"
	###   }
	###   orderBy: "created_at DESC"
	###   argumentsToColumn: {
	###     ":albumID": "album_id"
	###     ":userID": "user_id"
	###   }
	### }
}

type Album {
	### argument:required { key: "albumID" }
	### argument:optional { key: "userID", default: ":loggedInUserID" }

	### adapter(MusicDatabase):retrieve {
	###   table: "albums"
	###   columns: {
	###     album_id: ":albumID"
	###   }
	### }

	albumID: String!
	### adapter(MusicDatabase):field {
	###   column: "album_id"
	### }

	title: String!
	### adapter(MusicDatabase):field {
	###   column: "title"
	### }

	dateAdded: Int!
	### adapter(MusicDatabase):oneToOne {
	###   table: "users_albums"
	###   where: {
	###     album_id: ":albumID"
	###     user_id: ":userID"
	###   }
	### }

	artists: [Artist!]!
	### adapter(MusicDatabase):manyToMany {
	###   table: "albums_artists"
	###   where: {
	###     album_id: ":albumID"
	###   }
	###   orderBy: "created_at DESC"
	###   "argumentColumnMap": {
	###     ":artistID": "artist_id",
	###   }
	### }
}
```
