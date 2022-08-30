all: bot

.PHONY: bot

bot:
    go build -o . 

image:
    bot
    docker build -f docker/Dockerfile -t llelepipedepyro/discord_go .