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

export async function dockerV() {
  const v = await fetch('http://localhost:4000/health')

  return v
}