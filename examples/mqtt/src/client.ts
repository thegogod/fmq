import mqtt from 'mqtt';
import { faker } from '@faker-js/faker';

export class Client {
  private readonly _topics: string[] = [];
  private readonly _client: mqtt.MqttClient;

  private _onProduce = () => {};
  private _onConsume = () => {};

  constructor(i: number, topics: string[], url: string) {
    this._topics = topics;
    this._client = mqtt.connect(url, { clientId: `client/${i}` });
    this._client.on('connect', this._onConnect.bind(this));
    this._client.on('message', this._onConsume);
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
      this._client.publish(topic, faker.internet.email(), (err) => {
        if (err) console.error(err);
        this._onProduce();
      });
    }, 10);
  }

  onProduce(callback: () => void) {
    this._onProduce = callback;
  }

  onConsume(callback: () => void) {
    this._onConsume = callback;
    this._client.on('message', callback);
  }
}
