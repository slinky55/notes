import type { PageLoad } from './$types';

export const load = (async ({fetch, params}) => {
  const res = await fetch("http://localhost:7100/api/user/profile", {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
    },
    credentials: "include",
  })

  const profile = await res.json();

  return {
    profile
  };
}) satisfies PageLoad;