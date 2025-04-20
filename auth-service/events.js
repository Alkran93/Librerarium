const amqp = require('amqplib');

let channel;

async function connectRabbitMQ() {
  const conn = await amqp.connect('amqp://localhost');
  channel = await conn.createChannel();
  await channel.assertQueue('auth_events');
  console.log('[MOM] Connected to RabbitMQ (auth_events)');
}

async function publishUserEvent(event, payload) {
  if (!channel) {
    console.warn('[MOM] Channel not ready. Skipping publish.');
    return;
  }

  const message = JSON.stringify({ event, data: payload });
  channel.sendToQueue('auth_events', Buffer.from(message));
  console.log(`[MOM] Published event: ${event}`, payload);
}

module.exports = {
  connectRabbitMQ,
  publishUserEvent
};
