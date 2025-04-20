function publishUserEvent(eventName, payload) {
    // TODO: Replace with real RabbitMQ
    console.log(`[MOM] Event: ${eventName}`);
    console.log(`[MOM] Payload:`, payload);
  }
  
  module.exports = {
    publishUserEvent
  };
  