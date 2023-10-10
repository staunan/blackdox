import express from "express";
const router = express.Router();

import {
    home,
    test,
} from "../controllers/index.js";

router.get(
    "/",
    home
);

router.get(
    "/test",
    test
);

export const indexRoute = router;