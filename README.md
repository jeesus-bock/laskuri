# Laskuri

The name's Finnish for counter and that's pretty much what it is. An ultrasimple
visitor counter meant to be integrated on webpages.

## Depedencies

- **go-sqlite3**, uses sqlite db to save the visitors.
- **gin**, definitely an overkill solution for something this simple, but I asked
  someone what tech to use to handle the HTTP-side of things and they said gin.
  So I did.

## Running

First build the laskuri binary. Run it by just `./laskuri`, it accepts no arguments.
The binary will create the `laskuri.db` database file automagically. Now you have
the thing listening on port 8080.

I've used Nginx, but any proxy should work. The proxy needs to add header called
`"X-Real-IP"` to the request with the original request IP.

## The front-end

`example.html` shows a very simple but comprehensible example on how to use laskuri
from the webpage side of things.

### Assorted

The hashing algorithm is NON-cryptographic, but it's not protecting any too
sensitive information. It could easily be replaced with something more
hardcore if need be.

Gin is in the debug mode by default. It'll printout the environment variable
to set if you have a problem with that. Same goes with the default port which
is set at 8080.

All the error logging is set to fatal, if something goes wrong, laskuri will go
down. It's quick to get back up though, so... :)
