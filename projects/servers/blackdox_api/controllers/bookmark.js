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
		if(!title){
			throw Error("Please provide title");
		}
		if(!domain){
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
		let bookmarkObj = new Bookmark(bookmarkData);
    	await bookmarkObj.save();

		return res.status(200).json({message: "Bookmark has been created successfully!", bookmark: bookmarkObj});
	} catch (err) {
		console.log(err);
		return res.status(400).json({message: err.message});
	}
};

export const updateBookmark = async (req, res) => {
	try {
		let {
			bookmark_id,
			title,
			domain,
			username,
			password,
			note
		} = req.body;

		// Validation --
		if(!bookmark_id){
			throw Error("Please provide bookmark id");
		}
		if(!title){
			throw Error("Please provide title");
		}
		if(!domain){
			throw Error("Please provide domain");
		}

		let bookmark = await Bookmark.findOne({internalId: bookmark_id});

		if(!bookmark){
			throw Error("Bookmark not found with the provided bookmark ID: " + bookmark_id);
		}
		console.log(bookmark);

		// Update the data --
		bookmark.title = title;
		bookmark.domain = domain;
		bookmark.username = username;
		bookmark.password = password;
		bookmark.note = note;
    	await bookmark.save();

		return res.status(200).json({message: "Bookmark has been updated successfully!", bookmark: bookmark});
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
		let condition = {is_trash: false};

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

export const deleteBookmark = async (req, res) => {
	try {
		let {
			bookmark_id
		} = req.body;

		if(!bookmark_id){
			throw Error("Please provide bookmark id");
		}

		let bookmark = await Bookmark.findOne({ internalId: bookmark_id });
		if(!bookmark){
			throw Error("Cannot find bookmark with the provided id");
		}

		bookmark.is_trash = true;
		bookmark.save();

		return res.status(200).json({message: "Bookmark has been deleted successfully!"});
	} catch (err) {
		console.log(err);
		return res.status(400).json({message: err.message});
	}
};