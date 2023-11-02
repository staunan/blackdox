import express from "express";
const router = express.Router();

import {
    newSpending,
    getSpendings,
    deleteBookmark,
    updateBookmark
} from "../controllers/spending.js";

router.post(
    "/new_spending",
    newSpending
);

router.post(
    "/update_bookmark",
    updateBookmark
);

router.get(
    "/get_spendings",
    getSpendings
);

router.post(
    "/delete_bookmark",
    deleteBookmark
);

export const spendingRoute = router;