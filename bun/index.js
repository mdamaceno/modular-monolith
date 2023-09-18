import { Elysia } from 'elysia';
import * as component from './components';

const app = new Elysia();

app.use(app.group('/marketing', component.marketing));
app.use(app.group('/order', component.order));

app.listen(8081);

console.log(`Server running on port ${app.server.port}`);
