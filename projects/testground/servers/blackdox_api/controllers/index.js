export const home = async (req, res) => {
	try {
		res.render('pages/home/home');
	} catch (err) {
		console.log(err);
		return res.status(400).json({message: err.message});
	}
};
export const test = async (req, res) => {
	try {
		console.log("Hello World");
		return res.status(200).json({message: "Success"});
	} catch (err) {
		console.log(err);
		return res.status(400).json({message: err.message});
	}
};