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