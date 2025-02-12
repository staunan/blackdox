<div bind:this={displayElement} />
<script>
import "quill-emoji/dist/quill-emoji.css";
export let content = "";
let displayElement;
$: {
    if(content){
        render(content);
    }
}
async function render(content){
    let htmlOutput = await getStoryAsHTML(content);
    displayElement.innerHTML = htmlOutput;
}
async function getStoryAsHTML(quillData){
    const { default: Quill } = await import("quill");
    const { default: Emoji } = await import("quill-emoji");
    Quill.register("modules/emoji", Emoji);
    var tempQuill=new Quill(document.createElement("div"));
    tempQuill.setContents(quillData.ops);
    return tempQuill.root.innerHTML;
}
</script>
<style>
/* :global(h1 .ql-emojiblot .ap) {
    transform: scale(2) translateY(-3px);
} */
</style>