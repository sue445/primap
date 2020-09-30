import React from "react";
import ReactDOM from "react-dom";
import * as Sentry from "@sentry/react";
import { Integrations } from "@sentry/tracing";
import MapContainer from "./components/MapContainer";

Sentry.init({
  dsn: process.env.REACT_APP_SENTRY_DSN,
  environment: process.env.NODE_ENV,
  integrations: [new Integrations.BrowserTracing()],
  tracesSampleRate: 1.0,
});

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
  document.getElementById("root")
);
