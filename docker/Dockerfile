FROM alpine
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
ADD ./golang-discord-bot.exe
ENTRYPOINT ["/golang-discord-bot.exe"]