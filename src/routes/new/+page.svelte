<script lang="ts">
    import { goto } from "$app/navigation";
    import Cookies from "js-cookie";

    let title = '';
    let body = '';
    let formError = '';

    const handleSubmit = async () => {
        try {
            const res = await fetch("http://localhost:7100/api/note/create", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    "Cookie": "jwt=" + Cookies.get("jwt")
                },
                body: JSON.stringify({
                    "title": title,
                    "body": body
                }),
                credentials: "include"
            });

            const data = await res.json();

            if (res.ok) {
                await goto("/dashboard");
            } else {
                formError = data.error;
            }

        } catch (e: Error) {
            console.log(e.message);
        }
    }

    $: formValid = title && body
</script>

<style>
    .new-note-form {
        display: flex;
        flex-direction: column;
        align-items: center;
        margin-top: 2rem;
    }

    .new-note-form label {
        display: flex;
        flex-direction: column;
        margin-bottom: 1rem;
    }

    .new-note-form label span {
        margin-bottom: 0.5rem;
        font-weight: bold;
    }

    .new-note-form input {
        padding: 0.5rem;
        border-radius: 0.25rem;
        border: 1px solid #ccc;
    }

    .new-note-form input:invalid {
        border: 1px solid red;
    }

    .new-note-form textarea {
        padding: 0.5rem;
        border-radius: 0.25rem;
        border: 1px solid #ccc;
        resize: none;
        height: 10rem;
    }

    .new-note-form textarea:invalid {
        border: 1px solid red;
    }


    .new-note-form .error-message {
        color: red;
        margin-top: 0.5rem;
    }


    .new-note-form button {
        padding: 0.5rem 1rem;
        border-radius: 0.25rem;
        border: none;
        background-color: #4caf50;
        color: #fff;
        font-weight: bold;
        cursor: pointer;
        margin-top: 1rem;
        transition: background-color 0.2s ease-in-out;
        width: 100%;
    }

    .new-note-form button:disabled {
        background-color: #ccc;
        cursor: default;
    }

    .new-note-form button:hover:not(:disabled) {
        background-color: #3e8e41;
    }

</style>

<div class="new-note-form">
    <h1>New Note</h1>

    {#if formError}
        <p class="error-message">{formError}</p>
    {/if}

    <form on:submit|preventDefault={handleSubmit}>
        <label>
            <span>Title</span>
            <input type="text" bind:value={title} required>
        </label>

        <label>
            <span>Body</span>
            <textarea bind:value={body} required></textarea>
        </label>

        <button type="submit" disabled={!formValid}>Create Note</button>
    </form>
</div>

