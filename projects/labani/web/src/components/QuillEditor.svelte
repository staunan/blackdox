<div class="editor-wrapper">
    <div bind:this={editorElement} />
</div>
<script>
import { onMount } from "svelte";
import {createEventDispatcher} from 'svelte';
export let placeholder = "Write your story...";
export let toolbarOptions = [
    [{ header: 1 }, { header: 2 }, "blockquote", "link", "image", "video"],
    ["bold", "italic", "underline", "strike"],
    [{ list: "ordered" }, { list: "ordered" }],
    [{ align: [] }],
    ["clean"]
];
export let content = "";
const dispatch = createEventDispatcher();
let editorElement;
let quillObj;

$: {
    if(quillObj){
        console.log("Content Changed", content);
        quillObj.setContents(content);
    }
}

onMount(async () => {
    const { default: Quill } = await import("quill");
    quillObj = new Quill(editorElement, {
        modules: {
            toolbar: toolbarOptions
        },
        theme: "snow",
        placeholder: placeholder
    });
    quillObj.on('text-change', function(delta, oldDelta, source) {
        if (source == 'api') {
            console.log("An API call triggered this change.");
        } else if (source == 'user') {
            onContentChangeHandler();
        }
    });
    console.log(quillObj);
});
function onContentChangeHandler(){
    let editor_content = quillObj.getContents();
    dispatch('change', editor_content);
}
</script>
<style>
@import 'https://cdn.quilljs.com/1.3.6/quill.snow.css';
:global(.editor-wrapper .ql-container) {
	border-bottom: none;
}
:global(.editor-wrapper .ql-toolbar) {
	border-top: none;
}
</style>