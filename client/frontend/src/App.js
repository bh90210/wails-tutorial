import React from 'react';
import './App.css';
import InteractiveList from './components/InteractiveList';
import Drop from './components/Drop';

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
          <Drop />
        </header>
        <body className="App-body">
          <InteractiveList />
        </body>
      </div>
    );
  }
}

export default App;
