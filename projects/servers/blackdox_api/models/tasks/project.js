import mongoose from "mongoose";

let projectScheama = new mongoose.Schema(
  {
    internalId: {
        type: String,
        required: true,
    },
    project_name: {
        type: String,
        require: true,
    },
    project_details: {
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

export const Project = mongoose.model("projects", projectScheama);