const db = require('./db');

function getUserByUsername(username) {
  return db.prepare('SELECT * FROM users WHERE username = ?').get(username);
}

function createTestUser() {
  const existing = getUserByUsername('admin');
  if (!existing) {
    db.prepare('INSERT INTO users (username, password) VALUES (?, ?)')
      .run('admin', '1234');
    console.log('Inserted test user: admin / 1234');
  }
}

module.exports = {
  getUserByUsername,
  createTestUser
};
