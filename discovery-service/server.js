const express = require("express");
const axios = require("axios");

const app = express();
const PORT = 8000;

// Store registered services
const services = {};

app.use(express.json());

// Register a service
app.post("/register", (req, res) => {
    const { serviceName, serviceUrl } = req.body;
    services[serviceName] = serviceUrl;
    console.log(`Service ${serviceName} registered at ${serviceUrl}`);
    res.sendStatus(200);
});

// Deregister a service
app.post("/deregister", (req, res) => {
    const { serviceName } = req.body;
    delete services[serviceName];
    console.log(`Service ${serviceName} deregistered`);
    res.sendStatus(200);
});

// Discover a service
app.get("/discover/:serviceName", (req, res) => {
    const { serviceName } = req.params;
    const serviceUrl = services[serviceName];
    if (serviceUrl) {
        res.json({ serviceUrl });
    } else {
        res.status(404).json({ message: "Service not found" });
    }
});

app.listen(PORT, () => {
    console.log(`Discovery server is running on port ${PORT}`);
});
