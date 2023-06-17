import express from 'express';
import {register} from "./api/router/router";

function main() {
    const app = express();
    const router = express.Router();

    register(router);

    app.use('/api', router);

    app.listen(3001, () => {
        console.log('Controller started on port 3002');
    });
}

main();