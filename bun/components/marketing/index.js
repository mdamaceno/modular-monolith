import announceSomething from './use_cases/announce_something';
import { data as serialize } from '../../helpers/serializer';

// This is the entrypoint for the component.
// Here you can do some validation, authentication, etc.
// and then pass the args to the component or do something else with it.
// You can follow the restful pattern (create, read (index, show), update, delete) or not, it's up to you.

export const controller = (app) => {
  app.get('/announces', () => serialize(announceSomething()));
  app.post('/announces', () => serialize(announceSomething()));

  return app;
};
