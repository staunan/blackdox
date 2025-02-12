// Import Packages --
import { v4 } from "uuid";

// Import Models --
import { Story } from "../models/stories/story.js";

export const createStory = async (req, res) => {
	try {
		let {
			story
		} = req.body;

		// Validation --
		if(!story){
			throw Error("Please provide story content");
		}

		let storyData = {
			internalId: v4(),
			story: story
		};
		let storyObj = new Story(storyData);
    	await storyObj.save();

		return res.status(200).json({message: "Story has been created successfully!", story: storyObj});
	} catch (err) {
		console.log(err);
		return res.status(400).json({message: err.message});
	}
};

export const updateStory = async (req, res) => {
	try {
		let {
			story_id,
			story
		} = req.body;

		// Validation --
		if(!story_id){
			throw Error("Please provide story id");
		}
		if(!story){
			throw Error("Please provide story content");
		}

		let story_obj = await Story.findOne({internalId: story_id});

		if(!story_obj){
			throw Error("Story not found with the provided story ID: " + story_id);
		}
		console.log(story_obj);

		// Update the data --
		story_obj.story = story;
    	await story_obj.save();

		return res.status(200).json({message: "Story has been updated successfully!", story: story_obj});
	} catch (err) {
		console.log(err);
		return res.status(400).json({message: err.message});
	}
};

export const getStories = async (req, res) => {
	try {
		const perPage = req.query.perPage ? req.query.perPage : 10;
      	const page = req.query.pageNum ? req.query.pageNum : 1;
      	const skip = (page - 1) * perPage;
		let condition = {is_trash: false};

		let resultSet = await Story.aggregate([
			{ $match: condition },
			{ $skip: skip },
			{ $limit: perPage },
			{ $sort: { createdAt: -1 } },
		]);

		return res.status(200).json({stories: resultSet});
	} catch (err) {
		console.log(err);
		return res.status(400).json({message: err.message});
	}
};

export const deleteStory = async (req, res) => {
	try {
		let {
			story_id
		} = req.body;

		if(!story_id){
			throw Error("Please provide story id");
		}

		let story_obj = await Story.findOne({ internalId: story_id });
		if(!story_obj){
			throw Error("Cannot find story with the provided id: " + story_id);
		}

		story_obj.is_trash = true;
		story_obj.save();

		return res.status(200).json({message: "Story has been deleted successfully!"});
	} catch (err) {
		console.log(err);
		return res.status(400).json({message: err.message});
	}
};