const db = require('../db/db');

function getUserByUsername(username) {
  return db.prepare('SELECT * FROM users WHERE username = ?').get(username);
}

function createTestUser() {
  const user = getUserByUsername('admin');
  if (!user) {
    db.prepare('INSERT INTO users (username, password) VALUES (?, ?)').run('admin', '1234');
    console.log('Inserted test user: admin / 1234');
  }
}

module.exports = {
  getUserByUsername,
  createTestUser
};
