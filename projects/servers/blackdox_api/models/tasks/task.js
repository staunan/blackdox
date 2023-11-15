import mongoose from "mongoose";

let taskScheama = new mongoose.Schema(
  {
    internalId: {
        type: String,
        required: true,
    },
    title: {
        type: String,
        require: true,
    },
    details: {
        type: String,
        required: false,
    },
    status: {
      type: String,
      required: false,
      default: "Pending"
    },
    is_trash: {
        type: Boolean,
        required: false,
        default: false
    },
  },
  { timestamps: true }
);

export const Task = mongoose.model("tasks", taskScheama);