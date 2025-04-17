const mongoose = require('mongoose');

const bookSchema = new mongoose.Schema({
  title: { type: String, required: true },
  author: { type: String, required: true },
  isbn: { type: String, unique: true },
  publishedDate: { type: Date },
  stock: { type: Number, default: 10 }
});

module.exports = mongoose.model('Book', bookSchema);
