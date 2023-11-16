import express from "express";
const router = express.Router();

import {
    createTask,
    createProject,
    updateTask,
    getPendingTasks,
    getTasks,
    getProjects,
    deleteTask,
} from "../controllers/tasks.js";

router.post(
    "/create_task",
    createTask
);

router.post(
    "/create_project",
    createProject
);

router.post(
    "/update_task",
    updateTask
);

router.get(
    "/get_pending_tasks",
    getPendingTasks
);

router.get(
    "/get_tasks",
    getTasks
);

router.get(
    "/get_projects",
    getProjects
);

router.post(
    "/delete_task",
    deleteTask
);

export const tasksRoute = router;