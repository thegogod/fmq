import mqtt from 'mqtt';
import { faker } from '@faker-js/faker';

export class Client {
  private readonly _i: number;
  private readonly _topics: string[] = [];
  private readonly _client: mqtt.MqttClient;

  constructor(i: number, topics: string[], url: string) {
    this._i = i;
    this._topics = topics;
    this._client = mqtt.connect(url, { clientId: `client/${i}` });
    this._client.on('connect', this._onConnect.bind(this));
    this._client.on('message', this._onMessage.bind(this));
  }

  private _onConnect() {
    console.log('connected...');
    const topics: string[] = [];

    for (let i = 0; i < 5; i++) {
      const topic = this._topics[Math.floor(Math.random() * this._topics.length)];
      topics.push(topic);
    }

    this._client.subscribe(topics);

    setInterval(() => {
      const topic = topics[Math.floor(Math.random() * topics.length)];
      this._client.publish(topic, faker.internet.email());
    }, 10);
  }

  private _onMessage(topic: string, payload: Buffer) {
    console.log(`client/${this._i}/${topic}`, payload.toString());
  }
}
