import mongoose from "mongoose";

let spendingScheama = new mongoose.Schema(
  {
    internalId: {
        type: String,
        required: true,
    },
    item_name: {
        type: String,
        require: true,
    },
    price: {
        type: Number,
        required: true,
    },
    quantity: {
        type: String,
        required: false,
    },
    note: {
        type: String,
        required: false,
    },
    is_trash: {
        type: Boolean,
        required: false,
        default: false
    },
  },
  { timestamps: true }
);

export const Spending = mongoose.model("spendings", spendingScheama);