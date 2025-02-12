import mongoose from "mongoose";

let bookmarkScheama = new mongoose.Schema(
  {
    internalId: {
        type: String,
        required: true,
    },
    title: {
        type: String,
        require: true,
    },
    domain: {
        type: String,
        required: true,
    },
    username: {
        type: String,
        required: false,
    },
    password: {
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

export const Bookmark = mongoose.model("bookmarks", bookmarkScheama);