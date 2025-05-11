# osuScores

osuScores is a command-line app written in Go used to display recent osu! scores, sort them by chosen metrics

## Features

- Print up to 1000 recently set osu! scores to standard output
- Write score data into a file (TODO)
- Sort scores by pp, accuracy, score, and more
- Find users who set recent scores

## Installation

osuScores requires at least [Go 1.24.0](https://go.dev/) installed on your system of choice.

To start with, clone the repository and install the dependencies from go.mod

```sh
git clone https://github.com/czajkub/osuScores.git
cd osuScores
go mod download
```

First, register a new oauth application on your osu! profile
[Registration help](https://osu.ppy.sh/docs/index.html#registering-an-oauth-application)
Set the application name as osuscores and callback URL as http://localhost:8080/callback

Then, create a .env file inside with your client_id and client_secret:
```sh
touch .env
echo "CLIENT_ID={your client_id}" >> .env
echo "CLIENT_SECRET={your client_secret}" >> .env
```

You can set an alias to run the command without typing out the whole path:
``` sh
alias osuscores="go run $(pwd)/main.go"
```

When you first run the command, you will be required to authorise your oauth2 request and paste your access token into the .env file (the same way as with client_id and client_secret)

## Commands

I will add this at some point, for now you can type 
```sh
osuscores
osuscores -h
osuscores --help
```

## Docker

Hopefully docker support will be added soon! (maybe)

