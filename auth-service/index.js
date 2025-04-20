const express = require('express');
const bodyParser = require('body-parser');
const { loginHandler } = require('./controllers/authController');
const { createTestUser } = require('./models/userModel');

require('dotenv').config();

const app = express();
const PORT = process.env.AUTH_PORT;
const SECRET = process.env.JWT_SECRET;

app.use(bodyParser.json());
createTestUser();

app.post('/login', loginHandler);

app.listen(PORT, () => {
  console.log(`Auth service running on port ${PORT}`);
});
