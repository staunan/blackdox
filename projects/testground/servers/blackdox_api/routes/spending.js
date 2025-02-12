import express from "express";
const router = express.Router();

import {
    newSpending,
    getSpendings,
    deleteSpending,
    updateSpending
} from "../controllers/spending.js";

router.post(
    "/new_spending",
    newSpending
);

router.post(
    "/update_spending",
    updateSpending
);

router.get(
    "/get_spendings",
    getSpendings
);

router.post(
    "/delete_spending",
    deleteSpending
);

export const spendingRoute = router;