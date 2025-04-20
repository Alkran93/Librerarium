const { getUserByUsername } = require('../models/userModel');
const { createToken } = require('../util/jwtUtil');
const { publishUserEvent } = require('../mom/publisher');

async function loginHandler(req, res) {
  const { username, password } = req.body;

  const user = getUserByUsername(username);
  if (!user || user.password !== password) {
    return res.status(401).json({ error: 'Invalid credentials' });
  }

  const token = createToken(username);

  await publishUserEvent('user_logged_in', {
    username,
    timestamp: new Date().toISOString()
  });

  return res.status(200).json({ token });
}

module.exports = {
  loginHandler
};
