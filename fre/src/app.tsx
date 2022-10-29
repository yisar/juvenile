import { render, useState, h, useEffect, useRef } from 'fre';
import { dockerV, getUser } from './api';
import Cookie from 'js-cookie';
import './style.css';

let token = Cookie.get('gitlab-token');

function App() {
  const [user, setUser] = useState({} as any);
  const t = useRef(null);
  useEffect(() => {
    getUser().then((data) => {
      let { username, avatar_url } = data as any;
      console.log(data);
      setUser({ username, avatar_url });
    });
  }, []);
  useEffect(() => {
    const source = new EventSource('http://localhost:4000/events/');
    console.log('链接成功');
    source.onmessage = function (e) {
      const log = document.createElement('li');
      log.textContent = e.data;
      t.current.appendChild(log);
    };
  }, []);

  const sse = () => {
    dockerV().then((res) => {
      // console.log(res)
    });
  };
  return (
    <div>
      {token ? (
        <div>
          <header>
            <button onclick={sse} class='new'>
              新建项目
            </button>
            <ul class='bio'>
              <li>
                <img src={user.avatar_url} alt='' />
              </li>
            </ul>
          </header>
          <main>
            <pre ref={t}></pre>
          </main>
        </div>
      ) : (
        <div>
          <a href='http://localhost:4000/login' class='gitlab'>
            <i class='iconfont icon-gitlab-line'></i>
            使用 gitlab 登陆
          </a>
        </div>
      )}
    </div>
  );
}
render(<App />, document.getElementById('app'));
