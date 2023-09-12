import { hello, make } from "./components";

const router = app => {
  app.get('/', () => hello.index());
  app.get('/hello', () => hello.index());
  app.post('/make', (req) => make.create({ req }));

  return app;
};

export default router;
