import React, { Component } from "react";
import Chart from "./CpuChart";

class Cpu extends Component {
  constructor() {
    super();
    this.state = {
      chartData: {},
      value: {},
      setValue: {}
    };
  }

  componentWillMount() {
    this.chartInit();
    this.wailsOn();
  }

  chartInit() {
    this.setState({
      chartData: {
        labels: [
          "message.avg",
          "User",
          "System",
          "Idle",
          "Nice",
          "New Bedford"
        ],
        datasets: [
          {
            label: "%",
            data: [873476, 181045, 153060, 106519, 105162, 95072],
            backgroundColor: [
              "rgba(255, 99, 132, 0.6)",
              "rgba(54, 162, 235, 0.6)",
              "rgba(255, 206, 86, 0.6)",
              "rgba(75, 192, 192, 0.6)",
              "rgba(153, 102, 255, 0.6)",
              "rgba(255, 159, 64, 0.6)",
              "rgba(255, 99, 132, 0.6)"
            ],
            borderColor: "rgba(255, 0, 0, 0.1)",
            pointBorderWidth: 50,
            pointHitRadius: 50,
            fill: false,
            borderWidth: 0,
            pointStyle: "circle"
          }
        ]
      }
    });
  }
  wailsOn() {
    window.wails.Events.On("error", message => {
      if (message) {
        this.setState({
          chartData: {
            datasets: [
              {
                data: [7777, message.avg, 153060, 106519, 105162, 95072],
              }
            ]
          }
        });
      }
    });
  }

  render() {
    return (
      <Chart
        chartData={this.state.chartData}
        location="Massachusetts"
        legendPosition="right"
      />
    );
  }
}

export default Cpu;
