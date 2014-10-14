Dookio-cli: Dookio Toolbelt
===============================

This script implements some shortcuts for retrieving useful information from your Dookio server.

```
$ dookio apps

*************************************
*       Dookio apps deployed
*************************************
--> apache.git.blabla.com (replicated in 2 containers)
Done.

$ dookio containers git/apache

*************************************
*       Dookio containers
*************************************
User: git
App: apache
[{"node": "http://123.123.123.123", "containers": ["a2ac34cc437a125125d8a0ef6d11e38612f30d5480a6a7b17dd66633fc3a27d0", "eae17e29a17b8578cf2f89522a1035b3dd504761b6461ff1f1c29e1339337b15"]}]
Done.

$ dookio stop git/apache

*************************************
*       Dookio containers
*************************************
User: git
App: apache
Stopping containers...
[{"node": "http://0.0.0.0", "containers": []}]
Done.

$ dookio scale=5 git/apache

*************************************
*       Dookio containers
*************************************
User: git
App: apache
Scaling to 5 containers...
The app can not scale unless is running!

Done.
```

## 1. Installation

Export the `DOOKIO_SERVER_ADDRESS` (you might want to add this line to the `.bashrc` file)

```
export DOOKIO_SERVER_ADDRESS="http://127.127.127.127"
```

Then, make the script executable:

```
sudo chmod +x dookio
```

And finally make it visible to your PATH, for example:

```
sudo mv dookio /usr/local/sbin/
```

## 2. Contribute

Simply create a PR. Easy :)

## 3. TODO

* Add command for removing containers
* Add command for removing deployed apps

## 4. LICENSE
MIT
