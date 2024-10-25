import mqtt from 'mqtt';
import { faker } from '@faker-js/faker';

(async () => {
  const client = mqtt.connect('mqtt://localhost');

  client.on('connect', () => {
    console.log('connected...');
    client.subscribe('test');

    setInterval(() => {
      client.publish('test', faker.internet.email());
    }, 100);
  });

  client.on('message', (topic, payload) => {
    console.log(topic, payload.toString());
  });
})();
