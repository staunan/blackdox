import mongoose from "mongoose";

let storyScheama = new mongoose.Schema(
  {
    internalId: {
        type: String,
        required: true,
    },
    story: {
        type: Object,
        require: true,
    },
    is_trash: {
        type: Boolean,
        required: false,
        default: false
    },
  },
  { timestamps: true }
);

export const Story = mongoose.model("stories", storyScheama);