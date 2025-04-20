const amqp = require('amqplib');

const EXCHANGE = 'my_exchange';
const ROUTING_KEY = 'test';
const AMQP_URL = 'amqp://user:password@3.82.109.178:5672/';

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
