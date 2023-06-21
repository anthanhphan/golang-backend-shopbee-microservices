const http = require("http");
const routes = require("./routes.json");

// Specify the default backend server URL for invalid routes
const defaultBackendURL = "http://localhost:8000";

// Create API gateway server
const server = http.createServer((req, res) => {
  // Select backend server based on the requested route
  const route = Object.keys(routes).find((r) => req.url.startsWith(r));
  const backendURL = route ? routes[route] : defaultBackendURL;

  // Create a proxy request to the selected backend server
  const proxyReq = http.request(backendURL, { method: req.method, headers: req.headers }, (proxyRes) => {
    // Forward the response from the backend server to the client
    res.writeHead(proxyRes.statusCode, proxyRes.headers);
    proxyRes.pipe(res);
  });

  // Forward the client request body to the backend server
  req.pipe(proxyReq);
});

const port = 3000;
server.listen(port, () => {
  console.log(`Shopbee API gateway is running on port ${port}`);
});
