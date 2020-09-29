import React from 'react';
import ReactDOM from 'react-dom';
import MapContainer from './components/MapContainer';

const App: React.FC<{ compiler: string; framework: string }> = (props) => {
  return (
    <div>
      <div>{props.compiler}</div>
      <div>{props.framework}</div>

      <MapContainer
        // @ts-ignore
        latitude={35.689846}
        longitude={139.700534}
        zoom={15}
      />
    </div>
  );
};

ReactDOM.render(
  <App compiler="TypeScript" framework="React" />,
  document.getElementById('root')
);
