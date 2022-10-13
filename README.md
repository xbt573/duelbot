# duelbot
Duelbot is a Telegram bot which adds duel functionality to your chat!

## Build and run
Firstly, setup environment variables:
- `BOT_TOKEN`: Telegram Bot API token
- `CHAT_ID`: Target chat ID
- `BOT_MODE`: Bot mode, all modes explained in `Bot modes` section
- `ADMIN_NAME`: Used only in `PING` bot mode, admin name

Then build project and run duelbot binary!

If you use docker - you can execute this command instead of building project manually:
```bash
$ docker build -t duelbot .
$ docker run -d -e BOT_TOKEN -e CHAT_ID -e BOT_MODE -e ADMIN_NAME duelbot
```

Or using docker-compose:
```bash
$ docker-compose up -d
```

## Bot modes
There is two bot modes:
- `ADMIN`: Mutes users with admin rights given to him
- `PING`: Pings admins
