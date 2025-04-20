const express = require('express');
const bodyParser = require('body-parser');
const { loginHandler } = require('./controllers/authController');
const { createTestUser } = require('./models/userModel');

const app = express();
const PORT = 3000;
const SECRET = 'super-secret-key';

app.use(bodyParser.json());
createTestUser();

app.post('/login', loginHandler);

app.listen(PORT, () => {
  console.log(`Auth service running on port ${PORT}`);
});
