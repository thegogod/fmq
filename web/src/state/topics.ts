import { Topic } from '../models';

import { StateProperty } from './state-property';

interface _TopicState {
  topic?: Topic;
  topics?: Record<string, Topic>;
}

export class TopicState extends StateProperty<_TopicState> {
  constructor(value: _TopicState = { }) {
    super(value);
  }
}
