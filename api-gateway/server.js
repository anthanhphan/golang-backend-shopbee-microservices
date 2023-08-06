require("dotenv").config();
const express = require("express");
const axios = require("axios");
const httpProxy = require("http-proxy");
const urls = require("./urls");

const app = express();
const proxy = httpProxy.createProxyServer({});
const { pushToLogDiscord } = require("./middleware/logger");
app.use(pushToLogDiscord);

let routes = [];

async function fetchRoutes() {
    const requests = urls.map((url) => axios.get(url).catch((error) => null));

    const responses = await axios.all(requests);

    const newRoutes = [];

    responses.forEach((resp, index) => {
        if (resp) {
            const parts = resp.config.url.split("/");
            const lastPart = parts.pop();
            const route = {
                path: "/api/v1/" + lastPart,
                target: resp.data["serviceUrl"],
            };

            newRoutes.push(route);
        } else {
            console.error(`Failed to fetch data for URL: ${urls[index]}`);
        }
    });

    routes = newRoutes;
    console.log(routes);
}

fetchRoutes().then(() => {
    setInterval(fetchRoutes, 60000); // Update routes every 2 minutes

    app.all("*", (req, res) => {
        const matchedRoute = routes.find((route) =>
            req.path.startsWith(route.path)
        );

        if (matchedRoute) {
            proxy.web(req, res, { target: matchedRoute.target });
        } else {
            res.status(404).send("Route not found");
        }
    });

    app.listen(80, () => {
        console.log("API Gateway is listening on port 80");
    });
});
