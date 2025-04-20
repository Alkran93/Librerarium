const express = require('express');
const bodyParser = require('body-parser');
const { loginHandler } = require('./authController');
const { createTestUser } = require('./userModel');

const app = express();
const PORT = 3000;

app.use(bodyParser.json());

createTestUser();

app.post('/login', loginHandler);

app.listen(PORT, () => {
  console.log(`Auth service running on port ${PORT}`);
});
