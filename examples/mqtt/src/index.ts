import mqtt from 'mqtt';
import { faker } from '@faker-js/faker';

(async () => {
  const client = mqtt.connect('mqtt://localhost:9876');

  client.on('connect', () => {
    console.log('connected...');
    client.subscribe('^[a-zA-Z0-9_]*$');

    setInterval(() => {
      client.publish('test', faker.internet.email());
    }, 100);
  });

  client.on('message', (topic, payload) => {
    console.log(topic, payload.toString());
  });
})();
