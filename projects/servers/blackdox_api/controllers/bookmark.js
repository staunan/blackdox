// Import Packages --
import { v4 } from "uuid";

// Import Models --
import { Bookmark } from "../models/bookmark/bookmark.js";

export const createBookmark = async (req, res) => {
	try {
		let {
			title,
			domain,
			username,
			password,
			note
		} = req.body;

		// Validation --
		if(!req.body.title){
			throw Error("Please provide title");
		}
		if(!req.body.domain){
			throw Error("Please provide domain");
		}

		let bookmarkData = {
			internalId: v4(),
			title: title,
			domain: domain,
			username: username,
			password: password,
			note: note
		};
		let bookmark = new Bookmark(bookmarkData);
    	await bookmark.save();

		return res.status(200).json({message: "Bookmark has been created successfully!"});
	} catch (err) {
		console.log(err);
		return res.status(400).json({message: err.message});
	}
};

export const getBookmarks = async (req, res) => {
	try {
		const perPage = req.query.perPage ? req.query.perPage : 10;
      	const page = req.query.pageNum ? req.query.pageNum : 1;
      	const skip = (page - 1) * perPage;
		let condition = {};

		let resultSet = await Bookmark.aggregate([
			{ $match: condition },
			{ $skip: skip },
			{ $limit: perPage },
		]);

		return res.status(200).json({bookmarks: resultSet});
	} catch (err) {
		console.log(err);
		return res.status(400).json({message: err.message});
	}
};