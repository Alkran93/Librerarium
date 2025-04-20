const amqp = require('amqplib');
const mongoose = require('mongoose');
const Book = require('./models/Book');
const connectDB = require('./db');

const RABBITMQ_URL = 'amqp://user:password@3.82.109.178:5672/';

async function startConsumer() {
  await connectDB();

  const conn = await amqp.connect(RABBITMQ_URL);
  const channel = await conn.createChannel();

  await channel.assertExchange('my_exchange', 'direct', { durable: true });
  const q = await channel.assertQueue('my_app', { durable: true });

  await channel.bindQueue(q.queue, 'my_exchange', 'test');

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

startConsumer().catch(console.error);
