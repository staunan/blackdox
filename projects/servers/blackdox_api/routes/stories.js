import express from "express";
const router = express.Router();

import {
    createStory,
    updateStory,
    getStories,
    deleteStory, 
} from "../controllers/stories.js";

router.post(
    "/create_story",
    createStory
);

router.post(
    "/update_story",
    updateStory
);

router.get(
    "/get_stories",
    getStories
);

router.post(
    "/delete_story",
    deleteStory
);

export const storiesRoute = router;