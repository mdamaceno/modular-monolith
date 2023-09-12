import { Elysia } from 'elysia';
import router from './router.js';

const app = new Elysia();

router(app).listen(8080);

console.log(`Server running on port ${app.server.port}`);
