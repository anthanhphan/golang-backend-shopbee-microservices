"use strict";

const { Client, GatewayIntentBits } = require("discord.js");
const CHANNELID_DISCORD = process.env.CHANNELID_DISCORD;
const TOKEN_DISCORD = process.env.TOKEN_DISCORD;

console.log(TOKEN_DISCORD);

class LoggerService {
    constructor() {
        this.client = new Client({
            intents: [
                GatewayIntentBits.DirectMessages,
                GatewayIntentBits.Guilds,
                GatewayIntentBits.GuildMessages,
                GatewayIntentBits.MessageContent,
            ],
        });

        this.channelId = CHANNELID_DISCORD;

        this.client.on("ready", () => {
            console.log(`${this.client.user.tag}`);
        });

        this.client.login(TOKEN_DISCORD);
    }

    sendToFormatCode(logData) {
        const {
            code,
            message = "This is a some additional information about the code",
            title = "Code Example",
        } = logData;
        const codeMessage = {
            content: message,
            embeds: [
                {
                    color: parseInt("00ff00", 16),
                    title,
                    description:
                        "```json\n" + JSON.stringify(code, null, 2) + "\n```",
                },
            ],
        };
        this.sendToMessage(codeMessage);
    }

    sendToMessage(message = "message") {
        const channel = this.client.channels.cache.get(this.channelId);
        if (!channel) {
            console.error(`Counldn't find the channel`);
            return;
        }

        channel.send(message).catch((e) => console.error(e));
    }
}

const loggerService = new LoggerService();

module.exports = loggerService;
