# amigo-test
Technical test for Amigo


```
If you haven’t already done so, please answer the following questions:
• Which three words describe a great developer?
• What is your most indispensable development workflow tool and why?
• How do you like your coffee?
```

- Curiosity, adaptability, team-work
- Git, mainly because of the versioning, sharing etc. I also love drone, helps a lot because it tests your builds and deploy them through docker. It's now a day hard to work - specially with go - without drone / github
- The same way I like my woman, strong & dark spirit that can be slightly soother with a bit of sugar. Obviously she is lactose intolerant


# Amigo test

```
Disclaimer:
    For my mind the project is NOT production ready. Simply because I would never send any config and docker file through github.
    There are images versioning that can be shared as you shared a repo with github.
    *.json is supposed to be shared off line and simply a structure through documentation should navigate.

    Therefor, I'm sharing it here in order to help a fast installation. I assume, and pardon me if I'm wrong, that we would prefer spending time on the project itself instead of installing it.
    That's why I chose Docker!
```

The project strictly follow the subject and will run following this schema (from a client or directly in a terminal):

`curl http://localhost:8080/messages/ -d 'Your message'`  
And  
`curl http://localhost:8080/messages/<message_id>`

# Installing through Docker
` Requierements :`
    - Docker Version 17.03.1
    - Compiled with Golang version go1.8.1 darwin/amd64

The steps are then simple:
    - git clone the repo
    - Inside the root of the repo `docker-compose build`
    - Once the build is finished `docker-compose up`

And you've the project running.
In order to use the project as shown in the subject, just keep in mind that the configured port is `:8080` for the project.

# Installing without docker
` Requierements :`
    - Compiled with Golang version go1.8.1 darwin/amd64
    - postgresql >= 9.4.4

Then you'll need a JSON file for the config, following those requirements (or use the one attached to the project, which is `config.jon`)
```
{
    "dbname": "",
    "dbuser": "",
    "dbpassword": "",
    "dburl": "",
    "port": (int)
}
```

The final step is simply to `go build`

Then `./amigo-test` and you can even specify an argument that take the path to your config.json (by default it selects `config.json` at the root of the go folder)
