const amqp = require('amqplib');
require('dotenv').config();

const EXCHANGE = process.env.MOM_EXCHANGE;
const ROUTING_KEY = process.env.MOM_ROUTING_KEYE;
const AMQP_URL =`amqp://${process.env.MOM_USER}:${process.env.MOM_PASSWORD}@${process.env.MOM_HOST}:${process.env.MOM_PORT}/`;

async function publishUserEvent(eventName, payload) {
  try {
    const conn = await amqp.connect(AMQP_URL);
    const channel = await conn.createChannel();
    await channel.assertExchange(EXCHANGE, 'direct', { durable: true });

    const message = JSON.stringify({ evento: eventName, data: payload });
    channel.publish(EXCHANGE, ROUTING_KEY, Buffer.from(message));

    console.log(`[MOM] Evento enviado: ${eventName}`);
    await channel.close();
    await conn.close();
  } catch (err) {
    console.error('[MOM] Error al enviar evento:', err.message);
  }
}

module.exports = {
  publishUserEvent
};
