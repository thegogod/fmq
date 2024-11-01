import { faker } from '@faker-js/faker';

import { Client } from './client';

(async () => {
  let topics: string[] = [];
  const count = +(process.env.COUNT || 20);

  for (let i = 0; i < Math.floor(count / 2); i++) {
    topics.push(faker.lorem.word());
  }

  for (let i = 0; i < count; i++) {
    (async () => {
      new Client(i, topics, process.env.URL || 'tcp://localhost:1883');
    })();
  }
})();
