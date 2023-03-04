<script lang="ts">
    import { goto } from "$app/navigation";

    let name = '';
    let email = '';
    let password = '';
    let confirmPassword = '';
    let formError = '';

    const handleSubmit = async () => {
        if (password !== confirmPassword) {
            formError = 'Passwords do not match';
        } else {
            try {
                const res = await fetch("http://localhost:7100/api/user/create", {
                    method: "POST",
                    headers: {
                        "Content-Type": "application/json"
                    },
                    body: JSON.stringify({
                        "name": name,
                        "email": email,
                        "password": password
                    })
                });

                const data = await res.json();

                if (res.ok) {
                    await goto("/login");
                } else {
                    formError = data.message;
                }

            } catch (e: Error) {
                formError = e.message;
            }
        }
    }

    $: passwordMatch = password === confirmPassword;
    $: formValid = name && email && password && confirmPassword && passwordMatch;
</script>

<style>
    .signup-form {
        display: flex;
        flex-direction: column;
        align-items: center;
        margin-top: 2rem;
    }

    .signup-form label {
        display: flex;
        flex-direction: column;
        margin-bottom: 1rem;
    }

    .signup-form label span {
        margin-bottom: 0.5rem;
        font-weight: bold;
    }

    .signup-form input {
        padding: 0.5rem;
        border-radius: 0.25rem;
        border: 1px solid #ccc;
    }

    .signup-form input:invalid {
        border: 1px solid red;
    }

    .signup-form .error-message {
        color: red;
        margin-top: 0.5rem;
    }

    .signup-form button {
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

    .signup-form button:disabled {
        background-color: #ccc;
        cursor: default;
    }

    .signup-form button:hover:not(:disabled) {
        background-color: #3e8e41;
    }
</style>

<div class="signup-form">
    <h1>Create an account</h1>

    {#if formError}
        <p class="error-message">{formError}</p>
    {/if}

    <form on:submit|preventDefault={handleSubmit}>
        <label>
            <span>Name</span>
            <input type="text" bind:value={name} required>
        </label>

        <label>
            <span>Email</span>
            <input type="email" bind:value={email} required>
        </label>

        <label>
            <span>Password</span>
            <input type="password" bind:value={password} required>
        </label>

        <label>
            <span>Confirm Password</span>
            <input type="password" bind:value={confirmPassword} required>
        </label>

        {#if password && !passwordMatch}
            <p class="error-message">Passwords do not match</p>
        {/if}

        <button type="submit" disabled={!formValid}>Sign up</button>
    </form>
</div>

