// Firebase App (the core Firebase SDK) is always required and must be listed first
import * as firebase from "firebase/app";

// Add the Firebase products that you want to use
import "firebase/firestore";

import React from "react";
import ReactDOM from "react-dom";
import * as Sentry from "@sentry/react";
import { Integrations } from "@sentry/tracing";
import * as geofirex from "geofirex";

import MapContainer from "./components/MapContainer";

const firebaseConfig = {
  apiKey: process.env.REACT_APP_GOOGLE_BROWSER_API_KEY,
  authDomain: "primap.firebaseapp.com",
  databaseURL: "https://primap.firebaseio.com",
  projectId: "primap",
  storageBucket: "primap.appspot.com",
  messagingSenderId: "659376400894",
  appId: "1:659376400894:web:46a6da52d40c6983c238af",
  measurementId: "G-W2NTFNL7QE",
};
firebase.initializeApp(firebaseConfig);
const geo = geofirex.init(firebase);

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
        geo={geo}
      />
    </div>
  );
};

ReactDOM.render(
  <App compiler="TypeScript" framework="React" />,
  document.getElementById("root")
);
