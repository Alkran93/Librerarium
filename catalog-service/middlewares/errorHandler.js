const logger = require('../utils/logger');

function errorHandler(err, req, res, next) {
  logger.error({
    message: err.message,
    stack: err.stack,
    path: req.path,
    method: req.method
  });

  res.status(500).json({ error: 'Internal Server Error' });
}

module.exports = errorHandler;
