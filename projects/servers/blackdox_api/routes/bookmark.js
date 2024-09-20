import express from "express";
const router = express.Router();

import {
    createBookmark,
    getBookmarks,
    deleteBookmark,
    updateBookmark,
    searchBookmark
} from "../controllers/bookmark.js";

router.post(
    "/create_bookmark",
    createBookmark
);

router.post(
    "/update_bookmark",
    updateBookmark
);

router.get(
    "/get_bookmarks",
    getBookmarks
);

router.post(
    "/delete_bookmark",
    deleteBookmark
);

router.post(
    "/search_bookmark",
    searchBookmark
);

export const bookmarkRoute = router;