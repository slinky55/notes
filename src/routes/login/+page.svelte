<script lang="ts">
  import { goto } from "$app/navigation";

  let email = '';
  let password = '';
  let formError = '';

  const handleSubmit = async () => {
      try {
        const res = await fetch("http://localhost:7100/api/user/login", {
          method: "POST",
          headers: {
            "Content-Type": "application/json"
          },
          credentials: "include",
          body: JSON.stringify({
            "email": email,
            "password": password
          })
        });

        const data = await res.json();

        if (res.ok) {
          await goto("/dashboard");
        } else {
          formError = data.message;
        }

      } catch (e: Error) {
        console.log(e);
      }
  }

  $: formValid = email && password
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
  <h1>Login</h1>

  {#if formError}
    <p class="error-message">{formError}</p>
  {/if}

  <form on:submit|preventDefault={handleSubmit}>
    <label>
      <span>Email</span>
      <input type="email" bind:value={email} required>
    </label>

    <label>
      <span>Password</span>
      <input type="password" bind:value={password} required>
    </label>

    <button type="submit" disabled={!formValid}>Login</button>
  </form>
</div>

