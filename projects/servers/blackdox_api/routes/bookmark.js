import express from "express";
const router = express.Router();

import {
    createBookmark,
} from "../controllers/bookmark.js";

router.get(
    "/create_bookmark",
    createBookmark
);

export const bookmarkRoute = router;