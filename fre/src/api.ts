import Cookie from 'js-cookie';

export async function getUser() {
  const token = Cookie.get('gitlab-token');
  console.log(token)
  const res = await fetch('https://gitlab.com/api/v4/user', {
    headers: {
      Authorization: `Bearer ${token}`,
    },
  })
  const data = await res.json()

  return data
}
