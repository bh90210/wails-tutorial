import React, { Component } from "react";
import "./App.css";

import Menu from "./components/Tab";

class App extends Component {
  render() {
    return (
        <div id="app" className="App">
          <header className="App-header">
          </header>
          <Menu />
        </div>
    );
  }
}

export default App;
