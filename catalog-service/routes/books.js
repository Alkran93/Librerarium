const express = require('express');
const router = express.Router();
const bookController = require('../controllers/bookController');

// Crear un libro
router.post('/', bookController.createBook);

// Actualizar un libro
router.put('/:id', bookController.updateBook);

// Eliminar un libro
router.delete('/:id', bookController.deleteBook);

// Obtener todos los libros
router.get('/', bookController.getBooks);

// Obtener un libro por ID
router.get('/:id', bookController.getBookById);

module.exports = router;
