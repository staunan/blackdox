<!-- svelte-ignore a11y-click-events-have-key-events -->
<!-- svelte-ignore a11y-no-static-element-interactions -->
<!-- svelte-ignore a11y-mouse-events-have-key-events -->
<div class="picture_upload_box" on:click={openFileSelector} on:mouseover={openUploadOverlay} on:mouseleave={closeUploadOverlay}>
    <input style="display: none;" type="file" on:change={imageSelectedHandler} bind:this={fileInputElement}>
    <img id="myImg" alt="" src={imageData}>
    <div class="upload_overlay" class:upload_overlay_visible={overlay_visible} class:animate__scaleYUp={overlay_visible}>
        <div>Select Pictures</div>
    </div>
</div>
<script>
import axios from "axios";
import default_image_upload from '$lib/images/default_image_upload.png';
import {config} from "../../../config.js";

let imageData = default_image_upload;
let fileInputElement = null;
let overlay_visible = false;

function openUploadOverlay(){
    overlay_visible = true;
}
function closeUploadOverlay(){
    overlay_visible = false;
}
function openFileSelector(){
    fileInputElement.click();
}
function imageSelectedHandler(event){
    console.log(event);
    if (event.target.files && event.target.files[0]) {
        imageData = URL.createObjectURL(event.target.files[0]);
        uploadImageToServer(event.target.files);
    }
}
async function uploadImageToServer(images){
    let formData = new FormData();
    formData.append('file', images[0]);
    let header_config = {
        headers: {
            'Accept': 'application/json',
            'Accept-Language': 'en-US,en;q=0.8',
            'Content-Type': `multipart/form-data; boundary=${formData._boundary}`,
            "Access-Control-Allow-Origin": true
        },
    };
    let response = await axios.post(config.api_base_url + 'media/upload_picture', formData, header_config);
}
</script>
<style>
@keyframes scaleYUp{
    0% {
        height: 0;
    }
    100%{
        height: 100%;
    }
}
.picture_upload_box{
    width: 200px;
    height: 150px;
    cursor: pointer;
    border: 1px solid #ccc;
    border-radius: 5px;
    position: relative;
}
.picture_upload_box img{
    width: 200px;
    height: 150px;
    object-fit: contain;
    border-radius: 5px;
    background-color: #ddd;
}
.upload_overlay{
    position: absolute;
    left: 0;
    bottom: 0;
    right: 0;
    width: 100%;
    height: 100%;
    background-color: rgba(0, 0, 0, 0.5);
    cursor: pointer;
    display: none;
    justify-content: center;
    align-items: center;
    z-index: 100;
    color: #fff;
}
.upload_overlay_visible{
    display: flex !important;
}
.animate__scaleYUp{
    animation-name: scaleYUp;
    animation-duration: 500ms;
    animation-iteration-count: 1;
}
</style>