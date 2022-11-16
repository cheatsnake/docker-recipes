const apiService = require("../services/api.service");

class ApiController {
    getRandomNumber(req, res) {
        const { max, min } = req.query;
        const result = apiService.getRandomNumber(+max, +min);

        res.json({ result });
    }
}

module.exports = new ApiController();
