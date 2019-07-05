import React, { Component } from "react";
import Chart from "./LoadChart";

class Load extends Component {
  constructor() {
    super();
    this.state = {
      chartData: {},
      value: {},
      setValue: {}
    };
  }

  componentWillMount() {
    this.wailsOn();
  }

  wailsOn() {
    //window.wails.Events.On("cpu_stats", cpu_stats => {
      //if (cpu_stats) {
        //this.series = [cpu_stats.avg];
        this.setState({
          chartData: {
            labels: [
              "Percentage",
              "User",
              "System",
              "Idle",
              "Nice",
              "New Bedford"
            ],
            datasets: [
              {
                label: "%",
                data: [617594, 181045, 153060, 106519, 105162, 95072],
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
                pointStyle: 'circle'
              }
            ]
          }
        });
      //}
    //});
  }

  render() {
    return (
      <Chart
        chartData={this.state.chartData}
        location="Massachusetts"
      />
    );
  }
}

export default Load;
