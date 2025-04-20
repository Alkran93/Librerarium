const amqp = require('amqplib');
const mongoose = require('mongoose');
const Book = require('../models/Book');
const connectDB = require('../config/db');

require('dotenv').config();

const EXCHANGE = process.env.MOM_EXCHANGE;
const ROUTING_KEY = process.env.MOM_ROUTING_KEYE;
const RABBITMQ_URL =`amqp://${process.env.MOM_USER}:${process.env.MOM_PASSWORD}@${process.env.MOM_HOST}:${process.env.MOM_PORT}/`;

async function startConsumer() {
  await connectDB();

  const conn = await amqp.connect(RABBITMQ_URL);
  const channel = await conn.createChannel();

  await channel.assertExchange(EXCHANGE, 'direct', { durable: true });
  const q = await channel.assertQueue('my_app', { durable: true });

  await channel.bindQueue(q.queue, EXCHANGE, ROUTING_KEY);

  console.log('[Consumer] Esperando mensajes...');

  channel.consume(q.queue, async (msg) => {
    if (msg !== null) {
      const content = JSON.parse(msg.content.toString());
      console.log('[✓] Mensaje recibido del MOM:', content);

      if (content.evento === 'checkout') {
        const usuario = content.usuario || 'desconocido';
        const items = content.items || [];

        console.log(`🛒 Checkout procesado por el usuario: ${usuario}`);
        for (let item of items) {
          const book = await Book.findById(item.product_id);
          if (book) {
            book.stock = Math.max(0, book.stock - item.quantity);
            await book.save();
            console.log(`📘 Stock actualizado: ${book.title} → -${item.quantity}`);
          } else {
            console.log(`⚠️ Libro con ID ${item.product_id} no encontrado`);
          }
        }
      }

      channel.ack(msg);
    }
  });
}

module.exports = startConsumer;
