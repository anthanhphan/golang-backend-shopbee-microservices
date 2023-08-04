const express = require("express");
const axios = require("axios");
const urls = require("./urls");

const app = express();
const port = 3055;

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

// Fetch routes initially and then set an interval to update every 10 seconds
fetchRoutes().then(() => {
    setInterval(fetchRoutes, 10000); // Update routes every 10 seconds

    app.listen(port, () => {
        console.log(`API Gateway is listening at http://localhost:${port}`);
    });
});
