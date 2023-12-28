<div class="editor-wrapper">
    <div bind:this={editorElement} />
</div>
<script>
import { onMount } from "svelte";
import "quill-emoji/dist/quill-emoji.css";
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
    const { default: Emoji } = await import("quill-emoji");
    Quill.register("modules/emoji", Emoji);
    quillObj = new Quill(editorElement, {
        modules: {
            toolbar: toolbarOptions,
            "emoji-toolbar": true,
            "emoji-textarea": true,
            "emoji-shortname": true,
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
:global(#textarea-emoji) {
    max-width: 70% !important;
}
:global(#tab-panel) {
    padding: 10px;
}
:global(#tab-panel>span) {
    margin-right: 10px;
    margin-bottom: 10px;
}
</style>