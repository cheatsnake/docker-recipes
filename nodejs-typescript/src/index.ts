import http from "http";
import findMyWay from "find-my-way";

const router = findMyWay();
const PORT = 4000;

router.on("GET", "/", (req, res, params) => {
    res.writeHead(200, { "Content-Type": "application/json" });
    res.end(
        JSON.stringify({
            message: "hello from server",
        })
    );
});

const server = http.createServer((req, res) => {
    router.lookup(req, res);
});

server.on("error", (e: Error) => {
    console.error(e);
});

server.listen(PORT, () => {
    console.info(`Server listening on: http://localhost:${PORT}`);
});
