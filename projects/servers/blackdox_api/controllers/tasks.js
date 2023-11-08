// Import Packages --
import { v4 } from "uuid";

// Import Models --
import { Task } from "../models/tasks/task.js";

export const createTask = async (req, res) => {
	try {
		let {
			title,
			details
		} = req.body;

		// Validation --
		if(!title){
			throw Error("Please provide task title");
		}

		let taskData = {
			internalId: v4(),
			title: title,
			details: details
		};
		let taskObj = new Task(taskData);
    	await taskObj.save();

		return res.status(200).json({message: "Task has been created successfully!", task: taskObj});
	} catch (err) {
		console.log(err);
		return res.status(400).json({message: err.message});
	}
};

export const updateTask = async (req, res) => {
	try {
		let {
			task_id,
			title,
			details
		} = req.body;

		// Validation --
		if(!task_id){
			throw Error("Please provide task id");
		}
		if(!title){
			throw Error("Please provide task title");
		}

		let task = await Task.findOne({internalId: task_id});

		if(!task){
			throw Error("Task not found with the provided task ID: " + task_id);
		}
		console.log(task);

		// Update the data --
		task.title = title;
		task.details = details;
    	await task.save();

		return res.status(200).json({message: "Task has been updated successfully!", task: task});
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
		let condition = {is_trash: false};

		let resultSet = await Task.aggregate([
			{ $match: condition },
			{ $skip: skip },
			{ $limit: perPage },
		]);

		return res.status(200).json({tasks: resultSet});
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