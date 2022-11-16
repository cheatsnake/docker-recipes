class ApiService {
    getRandomNumber(max = 100, min = 1) {
        return Math.floor(Math.random() * (max - min + 1)) + min;
    }
}

module.exports = new ApiService();
