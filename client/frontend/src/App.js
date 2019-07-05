import React, { Component } from "react";
import "./App.css";
import Cpu from "./components/Cpu";
import Disk from "./components/Disk";
import Load from "./components/Load";
import Mem from "./components/Mem";

import IconLabelTabs from "./components/Tab";

class App extends Component {
  render() {
    return (
        <div id="app" className="App">
          <header className="App-header">
          </header>
          <body className="App-body">
            <IconLabelTabs />

          </body>
        </div>
    );
  }
}

export default App;
