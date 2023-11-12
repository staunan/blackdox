// Import Packages --
import { v4 } from "uuid";

// Import Models --
import { Spending } from "../models/spending/spending.js";

export const newSpending = async (req, res) => {
	try {
		let {
			item_name,
			price,
			quantity,
			note
		} = req.body;

		// Validation --
		if(!item_name){
			throw Error("Please provide item name");
		}
		if(!price){
			throw Error("Please provide price");
		}

		let spendingData = {
			internalId: v4(),
			item_name: item_name,
			price: price,
			quantity: quantity,
			note: note
		};
		let spendingObj = new Spending(spendingData);
    	await spendingObj.save();

		return res.status(200).json({message: "Spending has been added successfully!", spending: spendingObj});
	} catch (err) {
		console.log(err);
		return res.status(400).json({message: err.message});
	}
};

export const updateSpending = async (req, res) => {
	try {
		let {
			spending_id,
			item_name,
			price,
			quantity,
			note
		} = req.body;

		// Validation --
		if(!spending_id){
			throw Error("Please provide spending id");
		}
		if(!item_name){
			throw Error("Please provide item name");
		}
		if(!price){
			throw Error("Please provide price");
		}

		let spending = await Spending.findOne({internalId: spending_id});

		if(!spending){
			throw Error("Spending not found with the provided spending ID: " + spending_id);
		}
		console.log(spending);

		// Update the data --
		spending.item_name = item_name;
		spending.price = price;
		spending.quantity = quantity;
		spending.note = note;
    	await spending.save();

		return res.status(200).json({message: "Spending has been updated successfully!", spending: spending});
	} catch (err) {
		console.log(err);
		return res.status(400).json({message: err.message});
	}
};

export const getSpendings = async (req, res) => {
	try {
		const perPage = req.query.perPage ? req.query.perPage : 10;
      	const page = req.query.pageNum ? req.query.pageNum : 1;
      	const skip = (page - 1) * perPage;
		let condition = {is_trash: false};

		let resultSet = await Spending.aggregate([
			{ $match: condition },
			{ $skip: skip },
			{ $limit: perPage },
			{ $sort: { createdAt: -1 } },
		]);

		return res.status(200).json({spendings: resultSet});
	} catch (err) {
		console.log(err);
		return res.status(400).json({message: err.message});
	}
};

export const deleteSpending = async (req, res) => {
	try {
		let {
			spending_id
		} = req.body;

		if(!spending_id){
			throw Error("Please provide spending id");
		}

		let spending = await Spending.findOne({ internalId: spending_id });
		if(!spending){
			throw Error("Cannot find spending with the provided id: " + spending_id);
		}

		spending.is_trash = true;
		spending.save();

		return res.status(200).json({message: "Spending has been deleted successfully!"});
	} catch (err) {
		console.log(err);
		return res.status(400).json({message: err.message});
	}
};