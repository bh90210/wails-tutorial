import React, {
  Component
} from 'react';
import logo from './logo.png';
import './App.css';
import HelloWorld from './components/HelloWorld';
import Chart from './components/Chart';

class App extends Component {
  constructor() {
    super();
    this.state = {
      chartData: {}
    }
  }

  componentWillMount() {
    this.wailsOn();
  }

  wailsOn() {
    window.wails.events.on("cpu_stats", cpu_stats => {
      if (cpu_stats) {
        //this.series = [cpu_stats.avg];
        this.setState({
          chartData: {
            labels: ['Boston', 'Worcester', 'Springfield', 'Lowell', 'Cambridge', 'New Bedford'],
            datasets: [{
              label: 'Population',
              data: [
                617594,
                181045,
                153060,
                106519,
                105162,
                95072
              ],
              backgroundColor: [
                'rgba(255, 99, 132, 0.6)',
                'rgba(54, 162, 235, 0.6)',
                'rgba(255, 206, 86, 0.6)',
                'rgba(75, 192, 192, 0.6)',
                'rgba(153, 102, 255, 0.6)',
                'rgba(255, 159, 64, 0.6)',
                'rgba(255, 99, 132, 0.6)'
              ]
            }]
          }
        });
      }
    });
  }

  render() {
    return (
      <div id="app" className = "App">
        <Chart chartData = {this.state.chartData} location = "Massachusetts" legendPosition = "bottom" />
      </div>
    );
  }
}

export default App;
