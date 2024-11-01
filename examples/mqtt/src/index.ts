import { faker } from '@faker-js/faker';

import { Client } from './client';

(async () => {
  let topics: string[] = [];
  const count = +(process.env.COUNT || 20);

  for (let i = 0; i < Math.floor(count / 2); i++) {
    topics.push(faker.lorem.word());
  }

  let produced = 0;
  let consumed = 0;

  async function start(i: number) {
    const client = new Client(i, topics, process.env.URL || 'tcp://localhost:1883');

    client.onProduce(() => {
      produced++;
    });

    client.onConsume(() => {
      consumed++;
    });
  }

  const clients: Promise<void>[] = [];

  for (let i = 0; i < count; i++) {
    clients.push(start(i));
  }

  setInterval(() => {
    console.info(`produced: ${produced}/sec`);
    console.info(`consumed: ${consumed}/sec`);

    produced = 0;
    consumed = 0;
  }, 1000);

  await Promise.all(clients);
})();
