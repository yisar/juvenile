import { render, useState, h, useEffect, useRef } from 'fre';

function App() {
  const [evnet, setEvent] = useState('');
  const t = useRef(null);

  useEffect(() => {
    const source = new EventSource('http://localhost:8000/events/');
    source.onmessage = function (e) {
      // setEvent(e.data as any);
      const log = document.createElement('li');
      log.textContent = e.data
      t.current.appendChild(log);
    };
  }, []);

  return <pre ref={t}>{evnet}</pre>;
}
render(<App />, document.getElementById('app'));
