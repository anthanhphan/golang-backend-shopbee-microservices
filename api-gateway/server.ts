const http = require("http");
const routes = require("./routes.json");

const { createProxyMiddleware } = require("http-proxy-middleware");

// Define API routes and corresponding backend services
// const { routes } = require("./routes.js");

// Specify the default backend server URL for invalid routes
const defaultBackendURL = "http://localhost:8000";

// Create proxy middleware
const proxy = createProxyMiddleware({
    changeOrigin: true,
    xfwd: true,
    router: (req) => {
        // Select backend server based on the requested route
        const route = Object.keys(routes).find((r) => req.url.startsWith(r));
        if (route) {
            return routes[route];
        }
        return defaultBackendURL; // Use the default backend URL for invalid routes
    },
});

// Create API gateway server
const server = http.createServer((req, res) => {
    proxy(req, res);
});

const port = 80;
server.listen(port, () => {
    console.log(`Shopbee API gateway is running on port ${port}`);
});
