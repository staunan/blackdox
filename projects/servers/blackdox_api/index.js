import sirv from 'sirv';
import express from "express";
import compression from 'compression';
import bodyParser from "body-parser";
import mongoose from "mongoose";
import cors from "cors";

// Import ENV --
import "./loadEnv.js";

// Import DB URL --
import { DB_URL } from "./configs/mongodb.js";

// Import routes --
import { indexRoute } from './routes/index.js';

const app = express();
const dev = process.env.NODE_ENV === 'development';

// Connecting to mongodb --
mongoose.connect(DB_URL(), {});

// BodyParser Middleware --
app.use(bodyParser.urlencoded({
	extended: true
}));
app.use(bodyParser.json({ 
	limit: "50mb",
	verify: function(req, res, buf, encoding) {
        req.rawBody = buf.toString(encoding || 'utf8');
    }
}));

// Setup View Directory --
app.set('view engine', 'pug');

// Setup Cors --
app.use(cors({ origin: [], credentials: true }));

// Register Routes --
app.use("/", indexRoute);

app.use(
	compression({ threshold: 0 }),
	sirv('public', { dev })
);

const appInstance = app.listen(process.env.PORT || 80, () => {
	console.log(`🚀 Server ready at port ${process.env.PORT || 80}`);
});

export {appInstance};