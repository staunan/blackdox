import express from "express";
const router = express.Router();

import {
    createTask,
    updateTask,
    getTasks,
    deleteTask, 
} from "../controllers/tasks.js";

router.post(
    "/create_task",
    createTask
);

router.post(
    "/update_task",
    updateTask
);

router.get(
    "/get_tasks",
    getTasks
);

router.post(
    "/delete_task",
    deleteTask
);

export const tasksRoute = router;