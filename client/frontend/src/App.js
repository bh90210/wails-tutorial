import React from 'react';
import './App.css';
import FilesDrop from './components/FilesDrop';

class App extends React.Component {
  constructor(props, context) {
    super();
    this.state = {
      result2: " "
    };
  }
  
  componentDidMount(){
    window.wails.Events.On("cpu_usage", (cpu_usage) => {
      this.setState({
        result2: cpu_usage
      })
    });
  }

  render() {
    const { result2 } = this.state;
    return (
      <div id="app" className="App">
        <header className="App-header">
          <p>
            Welcome to your new <code>wails/react</code> project.
          </p>
          <h1>{result2}</h1>
          <FilesDrop />
        </header>
      </div>
    );
  }
}

export default App;
