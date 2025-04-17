const Book = require('../models/Book');
const logger = require('../utils/logger');

// Crear un libro
exports.createBook = async (req, res, next) => {
  try {
    const newBook = new Book(req.body);
    const savedBook = await newBook.save();
    logger.info('Libro creado exitosamente', { requestId: req.id, bookId: savedBook._id });
    res.status(201).json(savedBook);
  } catch (error) {
    logger.error('Error al crear el libro', { requestId: req.id, error: error.message });
    next(error);
  }
};

// Actualizar un libro
exports.updateBook = async (req, res, next) => {
  try {
    const updatedBook = await Book.findByIdAndUpdate(req.params.id, req.body, { new: true });
    if (!updatedBook) {
      return res.status(404).json({ error: 'Libro no encontrado' });
    }
    logger.info('Libro actualizado', { requestId: req.id, bookId: updatedBook._id });
    res.json(updatedBook);
  } catch (error) {
    logger.error('Error al actualizar el libro', { requestId: req.id, error: error.message });
    next(error);
  }
};

// Eliminar un libro
exports.deleteBook = async (req, res, next) => {
  try {
    const deletedBook = await Book.findByIdAndDelete(req.params.id);
    if (!deletedBook) {
      return res.status(404).json({ error: 'Libro no encontrado' });
    }
    logger.info('Libro eliminado', { requestId: req.id, bookId: deletedBook._id });
    res.json({ message: 'Libro eliminado' });
  } catch (error) {
    logger.error('Error al eliminar el libro', { requestId: req.id, error: error.message });
    next(error);
  }
};

// Obtener todos los libros
exports.getBooks = async (req, res, next) => {
  try {
    const books = await Book.find();
    logger.info('Listado de libros recuperado', { requestId: req.id, count: books.length });
    res.json(books);
  } catch (error) {
    logger.error('Error al recuperar libros', { requestId: req.id, error: error.message });
    next(error);
  }
};

// Obtener un libro por ID
exports.getBookById = async (req, res, next) => {
  try {
    const book = await Book.findById(req.params.id);
    if (!book) {
      return res.status(404).json({ error: 'Libro no encontrado' });
    }
    logger.info('Libro recuperado', { requestId: req.id, bookId: book._id });
    res.json(book);
  } catch (error) {
    logger.error('Error al recuperar el libro', { requestId: req.id, error: error.message });
    next(error);
  }
};
