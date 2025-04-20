const { getUserByUsername } = require('./userModel');
const { createToken } = require('./jwtUtil');
const { publishUserEvent } = require('./events');

function loginHandler(req, res) {
  const { username, password } = req.body;

  const user = getUserByUsername(username);

  if (!user || user.password !== password) {
    return res.status(401).json({ error: 'Invalid credentials' });
  }

  const token = createToken(username);

  publishUserEvent('user_logged_in', {
    username,
    timestamp: new Date().toISOString()
  });

  return res.status(200).json({ token });
}

module.exports = {
  loginHandler
};
