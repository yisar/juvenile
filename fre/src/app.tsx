import { render, useState, h, useEffect, useRef } from 'fre';
import { getUser } from './api';
import Cookie from 'js-cookie';
import './style.css';

let token = Cookie.get('gitlab-token');

function App() {
  const [user, setUser] = useState({} as any);
  useEffect(() => {
    getUser().then((data) => {
      let { username, avatar_url } = data as any;
      console.log(data);
      setUser({ username, avatar_url });
    });
  }, []);
  return (
    <div>
      {token ? (
        <ul class='bio'>
          {/* <li>{user.username}, </li> */}

          <li>
            <img src={user.avatar_url} alt='' />
          </li>
        </ul>
      ) : (
        <a href='http://localhost:4000/login' class='gitlab'>
          <i class='iconfont icon-gitlab-line'></i>
          使用 gitlab 登陆
        </a>
      )}
    </div>
  );
}
render(<App />, document.getElementById('app'));
