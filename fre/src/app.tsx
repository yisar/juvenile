import { render, useState, h, useEffect } from 'fre';

function App() {
  const [evnet, setEvent] = useState('');

  useEffect(() => {
    const source = new EventSource('/events/');
    source.onmessage = function (e) {
      setEvent(e.data as any);
    };
  }, []);

  return <div>{evnet}</div>;
}
render(<App />, document.getElementById('app'));
