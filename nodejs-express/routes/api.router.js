const { Router } = require("express");
const apiController = require("../controllers/api.controller");
const apiRouter = new Router();

apiRouter.get("/random", apiController.getRandomNumber);

module.exports = apiRouter;
