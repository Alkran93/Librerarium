const express = require('express');
const bodyParser = require('body-parser');
const Database = require('better-sqlite3');

const app = express();
const db = new Database('auth.db');
const PORT = 3000;

app.use(bodyParser.json());

db.prepare(`
  CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    username TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL
  );
`).run();

const existing = db.prepare('SELECT * FROM users WHERE username = ?').get('admin');

if (!existing) {
  db.prepare('INSERT INTO users (username, password) VALUES (?, ?)').run('admin', '1234');
  console.log('Inserted test user: admin / 1234');
}

app.post('/login', (req, res) => {
  const { username, password } = req.body;

  const user = db.prepare('SELECT * FROM users WHERE username = ?').get(username);

  if (!user || user.password !== password) {
    return res.status(401).json({ error: 'Invalid credentials' });
  }

  return res.status(200).json({ success: true });
});

app.listen(PORT, () => {
  console.log(`Auth service running on port ${PORT}`);
});
