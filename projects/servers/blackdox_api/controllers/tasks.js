// Import Packages --
import { v4 } from "uuid";

// Import Models --
import { Task } from "../models/tasks/task.js";
import { Project } from "../models/tasks/project.js";

let STATUS = {
    Draft: "Draft",
    Pending: "Pending",
    Hold: "Hold",
    Running: "Running",
    Finished: "Finished",
    Invalid: "Invalid"
};

export const createTask = async (req, res) => {
	try {
		let {
			title,
			details,
			status,
			project_id
		} = req.body;

		// Validation --
		if(!title){
			throw Error("Please provide task title");
		}
		if(!project_id){
			throw Error("Please provide project id");
		}

		let taskData = {
			internalId: v4(),
			title: title,
			project_id: project_id,
			details: details,
			status: status
		};
		let taskObj = new Task(taskData);
    	await taskObj.save();

		return res.status(200).json({message: "Task has been created successfully!", task: taskObj});
	} catch (err) {
		console.log(err);
		return res.status(400).json({message: err.message});
	}
};

export const createProject = async (req, res) => {
	try {
		let {
			project_name,
			project_details
		} = req.body;

		// Validation --
		if(!project_name){
			throw Error("Please provide task project name");
		}

		let projectData = {
			internalId: v4(),
			project_name: project_name,
			project_details: project_details
		};
		let projectObj = new Project(projectData);
    	await projectObj.save();

		return res.status(200).json({message: "Project has been created successfully!", project: projectObj});
	} catch (err) {
		console.log(err);
		return res.status(400).json({message: err.message});
	}
};

export const updateTask = async (req, res) => {
	try {
		let {
			task_id,
			project_id,
			title,
			details,
			status
		} = req.body;

		// Validation --
		if(!task_id){
			throw Error("Please provide task id");
		}
		if(!project_id){
			throw Error("Please provide project id");
		}
		if(!title){
			throw Error("Please provide task title");
		}
		if(!status){
			throw Error("Please provide task status");
		}


		let task = await Task.findOne({internalId: task_id});

		if(!task){
			throw Error("Task not found with the provided task ID: " + task_id);
		}
		console.log(task);

		// Update the data --
		task.project_id = project_id;
		task.title = title;
		task.details = details;
		task.status = status;
    	await task.save();

		return res.status(200).json({message: "Task has been updated successfully!", task: task});
	} catch (err) {
		console.log(err);
		return res.status(400).json({message: err.message});
	}
};

export const getPendingTasks = async (req, res) => {
	try {
		const perPage = req.query.perPage ? req.query.perPage : 10;
      	const page = req.query.pageNum ? req.query.pageNum : 1;
      	const skip = (page - 1) * perPage;
		let condition = {
			is_trash: false,
			status: STATUS.Pending
		};

		let resultSet = await Task.aggregate([
			{ $match: condition },
			{ $skip: skip },
			{ $limit: perPage },
			{ $sort: { createdAt: -1 } },
		]);

		return res.status(200).json({tasks: resultSet});
	} catch (err) {
		console.log(err);
		return res.status(400).json({message: err.message});
	}
};

export const getTasks = async (req, res) => {
	try {
		const perPage = req.query.perPage ? req.query.perPage : 10;
      	const page = req.query.pageNum ? req.query.pageNum : 1;
      	const skip = (page - 1) * perPage;
		let status = req.query.status;
		let project_id = req.query.project_id;

		let condition = {
			is_trash: false,
		};
		if(status){
			condition.status = status;
			condition.project_id = project_id;
		}

		let resultSet = await Task.aggregate([
			{ $match: condition },
			{ $skip: skip },
			{ $limit: perPage },
			{ $sort: { createdAt: -1 } },
			{ $lookup: {
					from: "projects",
					localField: "project_id",
					foreignField: "internalId",
					as: "project_details"
				}
			}
		]);

		return res.status(200).json({tasks: resultSet});
	} catch (err) {
		console.log(err);
		return res.status(400).json({message: err.message});
	}
};

export const getProjects = async (req, res) => {
	try {
		const perPage = req.query.perPage ? req.query.perPage : 10;
      	const page = req.query.pageNum ? req.query.pageNum : 1;
      	const skip = (page - 1) * perPage;
		let status = req.query.status;

		let condition = {
			is_trash: false,
		};
		if(status){
			condition.status = status;
		}

		let resultSet = await Project.aggregate([
			{ $match: condition },
			{ $skip: skip },
			{ $limit: perPage },
			{ $sort: { createdAt: -1 } },
		]);

		return res.status(200).json({projects: resultSet});
	} catch (err) {
		console.log(err);
		return res.status(400).json({message: err.message});
	}
};

export const deleteTask = async (req, res) => {
	try {
		let {
			task_id
		} = req.body;

		if(!task_id){
			throw Error("Please provide task id");
		}

		let task = await Task.findOne({ internalId: task_id });
		if(!task){
			throw Error("Cannot find task with the provided id: " + task_id);
		}

		task.is_trash = true;
		task.save();

		return res.status(200).json({message: "Task has been deleted successfully!"});
	} catch (err) {
		console.log(err);
		return res.status(400).json({message: err.message});
	}
};