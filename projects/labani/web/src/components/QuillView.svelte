<div bind:this={displayElement} />
<script>
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
    var tempQuill=new Quill(document.createElement("div"));
    tempQuill.setContents(quillData.ops);
    return tempQuill.root.innerHTML;
}
</script>