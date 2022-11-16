require("dotenv").config();
const express = require("express");
const cors = require("cors");
const apiRouter = require("./routes/api.router");
const app = express();

// enable middlewares
app.use(cors());
app.use(express.json());

// hosting static HTML page
app.use(express.static(__dirname + "/public"));

// setup API router
app.use("/api", apiRouter);

module.exports = app;
