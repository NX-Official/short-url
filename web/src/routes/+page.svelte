<script>
    import { getShortenURL } from "./request.js"; // 确保路径正确

    let inputUrl = "";
    let shortUrl = "";
    let errorMessage = "";

    async function handleShorten() {
        try {
            shortUrl = await getShortenURL(inputUrl);
        } catch (error) {
            errorMessage = error.message;
        }
    }

    async function copyToClipboard() {
        if (shortUrl) {
            try {
                await navigator.clipboard.writeText(shortUrl);
                // 可以添加一些反馈，比如显示“已复制”消息
            } catch (err) {
                errorMessage = "Failed to copy to clipboard :" + err.message;
            }
        }
    }
</script>

<div>
    <input type="text" bind:value={inputUrl} placeholder="Enter URL here" />
    <button on:click={handleShorten}>Shorten URL</button>
</div>

{#if errorMessage}
    <p class="error">{errorMessage}</p>
{/if}

<div>
    <input
        type="text"
        value={shortUrl}
        readonly
        placeholder="Shortened URL will appear here"
    />
    {#if shortUrl}
        <button on:click={copyToClipboard}>Copy to Clipboard</button>
    {/if}
</div>

<style>
    /* 样式可根据需要调整 */
    .error {
        color: red;
    }
</style>
