Dookio-cli: Dookio Toolbelt
===============================

Dookio tooltbelt with some interesting features as `listing`, `starting`, `stopping` and `scaling` applications.

## Commands:

* `dookio help`: Shows the list of commands with some examples.
* `dookio apps`: Shows the list of deployed applications.
* `dookio containers <user>/<repository>`: Shows the list of containers for a certain application. e.g "dookio containers git/apache"
* `dookio start <user>/<repository>`: Starts a previously stopped application. e.g "dookio start git/apache"
* `dookio stop <user>/<repository>`: Stops a running application. e.g "dookio stop git/apache"
* `dookio scale=x <user>/<repository>`: Scale X different containers to handle incoming requests for a certain application.  e.g "dookio scale=3 git/apache"

## Examples:

* $ ./dookio-cli apps

```
----> Dookio cli: Running the 'apps' command.
----> Dookio cli: Connecting with '123.123.123.123:8000...'

--> apache.git.blabla.com (replicated in 2 containers)
Done.
```

* $ ./dookio-cli containers git/apache

```
----> Dookio cli: Running the 'containers' command.
----> Dookio cli: For the 'git/portfolio' app.
----> Dookio cli: Connecting with '123.123.123.123:8000...'

[{"node": "http://0.0.0.0", "containers": [{"Status": "Up 2 hours", "Created": 1413832159, "Image": "git/apache:latest", "Id": "0fe2f69467e23f019f42992b7d740f9f8609382bd085c5c168c0bd912297b470", "Command": "/bin/sh -c 'uwsgi --http 0.0.0.0:80 --pythonpath /tmp/apache --static-map /static=/tmp/apache/static_media/ --module apache.wsgi'", "Names": ["/git_apache_4567"], "Ports": [{"PublicPort": 4567, "IP": "0.0.0.0", "Type": "tcp", "PrivatePort": 80}]}]}]
Done.
```

## 1. Installation

Export the `DOOKIO_SERVER_ADDRESS` (you might want to add this line to the `.bashrc` file)

```
export DOOKIO_SERVER_ADDRESS="123.123.123.123"
```

## 2. Contribute

Simply create a PR. Easy :)

## 3. TODO

* Add command for removing containers
* Add command for removing deployed apps

## 4. LICENSE
MIT
