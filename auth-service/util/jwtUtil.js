const jwt = require('jsonwebtoken');
const SECRET = 'super-secret-key';

function createToken(username) {
  return jwt.sign({ sub: username, role: 'user' }, SECRET, { expiresIn: '1h' });
}

module.exports = {
  createToken
};
