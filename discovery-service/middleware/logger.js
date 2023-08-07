"use strict";

const Logger = require("../logger/discord.log");

const pushToLogDiscord = async (req, res, next) => {
    try {
        Logger.sendToFormatCode({
            title: `Method: ${req.method}`,
            code:
                req.method === "GET"
                    ? `${req.method} http://${req.get("host")}${
                          req.originalUrl
                      }`
                    : req.body,
            message: `${req.get("host")}${req.originalUrl}`,
        });

        return next();
    } catch (error) {
        next(error);
    }
};

module.exports = { pushToLogDiscord };