const amqp = require('amqplib');
const Book = require('../models/Book');
const connectDB = require('../config/db');

const RABBITMQ_URL = 'amqp://user:password@34.205.157.11:5672/';

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
      console.log('[✓] Recibido:', content);

      if (content.evento === 'checkout') {
        for (let item of content.items) {
          // Aquí deberías tener campo `stock` en el modelo
          const book = await Book.findById(item.product_id);
          if (book) {
            book.stock = Math.max(0, book.stock - item.quantity);
            await book.save();
          }
        }
      }

      channel.ack(msg);
    }
  });
}

startConsumer().catch(console.error);
