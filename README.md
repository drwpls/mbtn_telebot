# Watch transactions from thiennguyen.app and send alerts to telegram

Prebuild docker image [hub.docker.com/r/drwpls/mbtn_telebot](https://hub.docker.com/r/drwpls/mbtn_telebot)

Use this docker compose file to run:

``` 
version: "3.8"

services:
  sccv:
    image: drwpls/sccv:latest
    environment:
      - ADMIN_ID=<telegram id to receive runtime error>
      - USER_ID=<thiennguyen.app user id>
      - GROUPCHAT_ID=<telegram group chat to send transaction alerts>
      - BOT_TOKEN=<telegram bot token>
      - DEBUG=true <set to true or unset, default false>
      - INTERVAL=10 <interval in seconds for bot crawling updates from thiennguyen.app>
    volumes:
      - /etc/localtime:/etc/localtime:ro
```