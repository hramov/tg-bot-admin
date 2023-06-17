import express from 'express';
import {register} from "./api/router/router";
import dotenv from 'dotenv';
dotenv.config();

function main() {
    const app = express();
    const router = express.Router();

    register(router);

    app.use('/api', router);

    app.use('/ui', express.static('bot/ui'))

    app.listen(3001, () => {
        console.log('Controller started on port 3001');
    });
}

main();