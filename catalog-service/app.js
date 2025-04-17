const express = require('express');
const connectDB = require('./config/db');
const bookRoutes = require('./routes/books');
const errorHandler = require('./middlewares/errorHandler');

const app = express();

// Middleware para asignar un ID único a cada request (útil para trazabilidad)
app.use((req, res, next) => {
  req.id = Date.now().toString();
  next();
});

// Middleware para parsear JSON
app.use(express.json());

// Conectar a la base de datos
connectDB();

// Definir rutas
app.use('/books', bookRoutes);

// Middleware para rutas no encontradas
app.use((req, res, next) => {
  res.status(404).json({ error: 'No encontrado' });
});

// Middleware global para manejo de errores
app.use(errorHandler);

app.use((err, req, res, next) => {
  const logger = require('./utils/logger');
  logger.error('Error global', { requestId: req.id, error: err.message });
  res.status(err.status || 500).json({ error: err.message || 'Error interno del servidor' });
});

// Levantar el servidor
const PORT = process.env.PORT || 3001;
app.listen(PORT, () => console.log(`Catalog Service escuchando en el puerto ${PORT}`));

module.exports = app;
