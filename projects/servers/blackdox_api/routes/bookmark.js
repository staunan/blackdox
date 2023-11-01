import express from "express";
const router = express.Router();

import {
    createBookmark,
    getBookmarks
} from "../controllers/bookmark.js";

router.post(
    "/create_bookmark",
    createBookmark
);

router.get(
    "/get_bookmarks",
    getBookmarks
);

export const bookmarkRoute = router;