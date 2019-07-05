import React, { Component } from "react";
import Button from "@material-ui/core/Button";
import { styled } from "@material-ui/styles";
import Chart from "./Chart";

const CpuButton = styled(Button)({
  background: "linear-gradient(45deg, #FE6B8B 30%, #FF8E53 90%)",
  border: 0,
  borderRadius: 3,

  color: "white",
  height: 48,
  padding: "0 30px",
  margin: "10px 10px"
});

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
    this.wailsOn();
  }

  wailsOn() {
    window.wails.Events.On("cpu_stats", cpu_stats => {
      if (cpu_stats) {
        //this.series = [cpu_stats.avg];
        this.setState({
          chartData: {
            labels: [
              "Boston",
              "Worcester",
              "Springfield",
              "Lowell",
              "Cambridge",
              "New Bedford"
            ],
            datasets: [
              {
                label: "Population",
                data: [617594, 181045, 153060, 106519, 105162, 95072],
                backgroundColor: [
                  "rgba(255, 99, 132, 0.6)",
                  "rgba(54, 162, 235, 0.6)",
                  "rgba(255, 206, 86, 0.6)",
                  "rgba(75, 192, 192, 0.6)",
                  "rgba(153, 102, 255, 0.6)",
                  "rgba(255, 159, 64, 0.6)",
                  "rgba(255, 99, 132, 0.6)"
                ]
              }
            ]
          }
        });
      }
    });
  }

  render() {
    return (
      <Chart className="Cpu-Page"
        chartData={this.state.chartData}
        location="Massachusetts"
        legendPosition="bottom"
      />
    );
  }
}

export default Cpu;
